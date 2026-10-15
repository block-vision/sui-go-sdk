package transaction

import (
	"bytes"
	"context"
	"errors"
	"os"
	"testing"

	"github.com/mr-tron/base58"
	"github.com/samber/lo"
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/blake2b"

	"github.com/block-vision/sui-go-sdk/models"
	"github.com/block-vision/sui-go-sdk/mystenbcs"
	"github.com/block-vision/sui-go-sdk/signer"
)

// mainnetChain is the mainnet chain identifier (genesis checkpoint digest).
const mainnetChain = "4btiuiMPvEENsttpZC7CZ53DruC3MAgfznDbASZ7DR6S"

func txDigest(bcsBytes []byte) string {
	h, _ := blake2b.New256(nil)
	h.Write([]byte("TransactionData::"))
	h.Write(bcsBytes)
	return base58.Encode(h.Sum(nil))
}

// Mainnet transactions that pay gas from an address balance (empty payment,
// ValidDuring expiration) and withdraw funds with redeem_funds; testdata holds
// their TransactionData BCS as returned by gRPC GetTransaction.
func TestMainnetTransactionRoundTrip(t *testing.T) {
	for _, digest := range []string{
		"G8BBZ82tRFogPtXSBUh7h1Q4CyvG8PYJ1wXMNKSVQWQk", // balance::redeem_funds + send_funds
		"BQxBiH5cD8jAoSGRkDokCQHgbsa2PPRNWHdsqeffKGqf", // coin::redeem_funds, shared objects, swap
	} {
		t.Run(digest, func(t *testing.T) {
			raw, err := os.ReadFile("testdata/" + digest + ".bcs")
			require.NoError(t, err)

			var td TransactionData
			n, err := mystenbcs.Unmarshal(raw, &td)
			require.NoError(t, err)
			require.Equal(t, len(raw), n)

			require.Empty(t, *td.V1.GasData.Payment)
			require.NotNil(t, td.V1.Expiration.ValidDuring)
			require.Equal(t, mainnetChain, base58.Encode(td.V1.Expiration.ValidDuring.Chain))
			withdrawals := lo.Filter(td.V1.Kind.ProgrammableTransaction.Inputs, func(in *CallArg, _ int) bool {
				return in.FundsWithdrawal != nil
			})
			require.Len(t, withdrawals, 1)
			require.NotNil(t, withdrawals[0].FundsWithdrawal.WithdrawFrom.Sender)

			encoded, err := td.Marshal()
			require.NoError(t, err)
			require.Equal(t, raw, encoded)
			require.Equal(t, digest, txDigest(encoded))
		})
	}
}

// One mainnet transaction per layout the old types got wrong; each must decode,
// re-encode byte-identically and hash to its digest.
func TestMainnetTransactionVariantsRoundTrip(t *testing.T) {
	commands := func(td *TransactionData) []*Command { return td.V1.Kind.ProgrammableTransaction.Commands }
	inputs := func(td *TransactionData) []*CallArg { return td.V1.Kind.ProgrammableTransaction.Inputs }
	cases := []struct {
		digest string
		check  func(t *testing.T, td *TransactionData)
	}{
		{"2RAkH6NNXzGZacRxaHKXJX3kQhDC9pomavcwhC3YzMn1", func(t *testing.T, td *TransactionData) {
			require.NotNil(t, td.V1.Expiration.Epoch, "Epoch expiration")
		}},
		{"2FN28ZkiPbsvJELxsJGpfuuFEpByZ9vYeFPu88GWHx5i", func(t *testing.T, td *TransactionData) {
			require.True(t, lo.ContainsBy(commands(td), func(c *Command) bool {
				return c.MakeMoveVec != nil && c.MakeMoveVec.Type != nil && c.MakeMoveVec.Type.U64 != nil
			}), "typed MakeMoveVec<u64>")
		}},
		{"6zBQYzPfCaqVjqCFYYXwSXnooLRpJTaf92gXs1RmLPi5", func(t *testing.T, td *TransactionData) {
			require.True(t, lo.ContainsBy(inputs(td), func(in *CallArg) bool {
				return in.Object != nil && in.Object.Receiving != nil
			}), "Receiving input")
		}},
		{"85N9ByQ3GL3U5ypiRpHjTtZY6TYYd4qYdVjtUx6f73JV", func(t *testing.T, td *TransactionData) {
			require.NotNil(t, td.V1.Expiration.Validity, "Validity expiration")
			require.NotNil(t, td.V1.Expiration.Validity.AllowedProposers)
			require.NotEmpty(t, td.V1.Expiration.Validity.AllowedProposers.Proposers)
		}},
		{"DffW3F5BYQD7Fke67rnMrRcLBKkPurVc2e2zus28uupj", func(t *testing.T, td *TransactionData) {
			in, ok := lo.Find(inputs(td), func(in *CallArg) bool { return in.FundsWithdrawal != nil })
			require.True(t, ok, "FundsWithdrawal input")
			sa := in.FundsWithdrawal.WithdrawFrom.SenderAllowance
			require.NotNil(t, sa, "SenderAllowance source")
			require.Equal(t, models.SuiAddress("0xec1dd42c6d20c43c8d00fdbbf18e737c779e45c887e203dc9a2d4d0485b00d40"), ConvertSuiAddressBytesToString(sa.Funder))
			require.Equal(t, models.SuiAddress("0xdd2cb366f6a0a36c524d7071ceee3d97002b5fb04e06fc58a3eb97d26e23962a"), ConvertSuiAddressBytesToString(sa.Allowance))
			require.Equal(t, uint64(949239), *in.FundsWithdrawal.Reservation.MaxAmountU64)
		}},
		{"7tD8myHtKqZaY9ZRsorCubvRTbn5ZaBEC6BTcLi5Ztv", func(t *testing.T, td *TransactionData) {
			pst := td.V1.Kind.ProgrammableSystemTransaction
			require.NotNil(t, pst, "ProgrammableSystemTransaction")
			require.True(t, lo.ContainsBy(pst.Inputs, func(in *CallArg) bool {
				return in.Object != nil && in.Object.SharedObject != nil && in.Object.SharedObject.NonExclusiveWrite
			}), "NonExclusiveWrite shared input")
		}},
	}
	for _, c := range cases {
		t.Run(c.digest, func(t *testing.T) {
			raw, err := os.ReadFile("testdata/" + c.digest + ".bcs")
			require.NoError(t, err)

			var td TransactionData
			n, err := mystenbcs.Unmarshal(raw, &td)
			require.NoError(t, err)
			require.Equal(t, len(raw), n)
			c.check(t, &td)

			encoded, err := td.Marshal()
			require.NoError(t, err)
			require.Equal(t, raw, encoded)
			require.Equal(t, c.digest, txDigest(encoded))
		})
	}
}

// ConsensusCommitPrologueV4 is one of the system kinds the SDK does not model.
func TestMainnetUnsupportedSystemTransaction(t *testing.T) {
	raw, err := os.ReadFile("testdata/CcL9155Q47BeQJUrRjg6o4fa2vxTVWdQQLSuvCKQpjcV.bcs")
	require.NoError(t, err)
	var td TransactionData
	_, err = mystenbcs.Unmarshal(raw, &td)
	require.ErrorIs(t, err, ErrUnsupportedTransactionKind)
}

func TestSharedObjectMutabilityEncoding(t *testing.T) {
	id := models.SuiAddressBytes{31: 6}
	for _, c := range []struct {
		ref  SharedObjectRef
		last byte
	}{
		{SharedObjectRef{ObjectId: id, InitialSharedVersion: 1}, 0},
		{SharedObjectRef{ObjectId: id, InitialSharedVersion: 1, Mutable: true}, 1},
		{SharedObjectRef{ObjectId: id, InitialSharedVersion: 1, Mutable: true, NonExclusiveWrite: true}, 2},
	} {
		bz, err := mystenbcs.Marshal(&ObjectArg{SharedObject: &c.ref})
		require.NoError(t, err)
		require.Equal(t, append(append([]byte{0x01}, id[:]...), 1, 0, 0, 0, 0, 0, 0, 0, c.last), bz)

		var back ObjectArg
		_, err = mystenbcs.Unmarshal(bz, &back)
		require.NoError(t, err)
		require.Equal(t, c.ref, *back.SharedObject)
	}

	bad := append(append([]byte{0x01}, id[:]...), 1, 0, 0, 0, 0, 0, 0, 0, 3)
	var back ObjectArg
	_, err := mystenbcs.Unmarshal(bad, &back)
	require.Error(t, err)
}

// A corrupt length prefix must fail on EOF instead of allocating the claimed size.
func TestDecodeHugeLengthFails(t *testing.T) {
	var move ProgrammableMoveCall
	data := append(make([]byte, 32), 0xff, 0xff, 0xff, 0xff, 0xff, 0x7f) // module name of ~4 TiB
	_, err := mystenbcs.Unmarshal(data, &move)
	require.Error(t, err)
}

// The builder must produce the exact input and expiration bytes found on chain
// in G8BBZ82t… .
func TestFundsWithdrawalAndValidDuringEncoding(t *testing.T) {
	raw, err := os.ReadFile("testdata/G8BBZ82tRFogPtXSBUh7h1Q4CyvG8PYJ1wXMNKSVQWQk.bcs")
	require.NoError(t, err)

	tx := NewTransaction()
	tx.FundsWithdrawal(94999900000, "0x2::sui::SUI", WithdrawFromSender())
	input, err := mystenbcs.Marshal(tx.Data.V1.Kind.ProgrammableTransaction.Inputs[0])
	require.NoError(t, err)
	require.True(t, bytes.Contains(raw, input), "FundsWithdrawal input %x not found in the on-chain bytes", input)

	chain, err := base58.Decode(mainnetChain)
	require.NoError(t, err)
	expiration, err := mystenbcs.Marshal(&TransactionExpiration{ValidDuring: &ValidDuring{
		MinEpoch: lo.ToPtr(uint64(1281)),
		MaxEpoch: lo.ToPtr(uint64(1282)),
		Chain:    chain,
		Nonce:    2556449602,
	}})
	require.NoError(t, err)
	require.True(t, bytes.HasSuffix(raw, expiration), "ValidDuring %x is not the on-chain tail", expiration)
}

func TestTransactionExpirationIsNotAnOption(t *testing.T) {
	tx := NewTransaction()
	sender := models.SuiAddress("0x1")
	tx.SetSender(sender).SetGasOwner(sender).SetGasPrice(1000).SetGasBudget(1)
	tx.SetGasPayment([]SuiObjectRef{})

	bz, err := tx.Data.Marshal()
	require.NoError(t, err)
	require.Equal(t, byte(0x00), bz[len(bz)-1], "a nil expiration encodes as TransactionExpiration::None")

	tx.SetExpiration(TransactionExpiration{Epoch: lo.ToPtr(uint64(5))})
	bz, err = tx.Data.Marshal()
	require.NoError(t, err)
	require.Equal(t, []byte{0x01, 0x05, 0, 0, 0, 0, 0, 0, 0}, bz[len(bz)-9:])
}

func TestEmptyGasPaymentNeedsValidDuring(t *testing.T) {
	tx := NewTransaction()
	sender := models.SuiAddress("0x1")
	tx.SetSender(sender).SetGasOwner(sender).SetGasPrice(1000).SetGasBudget(1)
	tx.SetGasPayment([]SuiObjectRef{})

	_, err := tx.build(false)
	require.ErrorIs(t, err, ErrAddressBalanceGasNeedsValidDuring)

	tx.SetExpiration(TransactionExpiration{ValidDuring: &ValidDuring{Chain: make([]byte, 32), Nonce: 1}})
	_, err = tx.build(false)
	require.NoError(t, err)
}

func TestUnresolvedObjectFailsToBuild(t *testing.T) {
	tx := NewTransaction()
	tx.TransferObjects([]Argument{tx.Object("0x5")}, tx.Pure("0x9"))

	_, err := tx.Data.V1.Kind.Marshal()
	require.ErrorIs(t, err, ErrUnresolvedInput)
}

func TestDecodeRejectsUnknownAndUnsupportedVariants(t *testing.T) {
	var arg CallArg
	_, err := mystenbcs.Unmarshal([]byte{0x03}, &arg) // index of the SDK-only UnresolvedPure
	require.Error(t, err)
	_, err = mystenbcs.Unmarshal([]byte{0x09}, &arg)
	require.Error(t, err)

	var td TransactionData
	_, err = mystenbcs.Unmarshal([]byte{0x00, 0x01}, &td) // V1, ChangeEpoch
	require.True(t, errors.Is(err, ErrUnsupportedTransactionKind), "got %v", err)
}

func TestTypeTagEncoding(t *testing.T) {
	cases := []struct {
		in   string
		want []byte
	}{
		// variant indices from move-core-types TypeTag:
		// Bool, U8, U64, U128, Address, Signer, Vector, Struct, U16, U32, U256
		{"bool", []byte{0x00}},
		{"u8", []byte{0x01}},
		{"u64", []byte{0x02}},
		{"u128", []byte{0x03}},
		{"address", []byte{0x04}},
		{"signer", []byte{0x05}},
		{"u16", []byte{0x08}},
		{"u32", []byte{0x09}},
		{"u256", []byte{0x0a}},
		{"vector<u8>", []byte{0x06, 0x01}},
		{"vector<vector<address>>", []byte{0x06, 0x06, 0x04}},
		{
			"0x2::coin::Coin<0x2::sui::SUI>",
			append(append(append([]byte{0x07}, addr2()...), "\x04coin\x04Coin\x01\x07"...),
				append(addr2(), "\x03sui\x03SUI\x00"...)...),
		},
	}
	for _, c := range cases {
		t.Run(c.in, func(t *testing.T) {
			tag, err := ParseTypeTag(c.in)
			require.NoError(t, err)
			bz, err := mystenbcs.Marshal(tag)
			require.NoError(t, err)
			require.Equal(t, c.want, bz)

			var back TypeTag
			_, err = mystenbcs.Unmarshal(bz, &back)
			require.NoError(t, err)
			again, err := mystenbcs.Marshal(&back)
			require.NoError(t, err)
			require.Equal(t, bz, again)
		})
	}

	for _, bad := range []string{"", "u7", "vector<u8", "0x2::sui", "0x2::sui::SUI<", "sui::SUI", "0x" + string(bytes.Repeat([]byte{'1'}, 65)) + "::a::B"} {
		_, err := ParseTypeTag(bad)
		require.ErrorIs(t, err, ErrInvalidTypeTag, bad)
	}
}

func addr2() []byte {
	b := make([]byte, 32)
	b[31] = 2
	return b
}

func TestKeypairSignersSignTransactions(t *testing.T) {
	one := append(make([]byte, 31), 1)
	two := append(make([]byte, 31), 2)
	sender, err := signer.NewSecp256k1Signer(one)
	require.NoError(t, err)
	sponsor, err := signer.NewSecp256r1Signer(two)
	require.NoError(t, err)

	tx := NewTransaction().SetKeypairSigner(sender).SetKeypairSponsoredSigner(sponsor)
	require.Nil(t, tx.Signer)
	require.Nil(t, tx.SponsoredSigner)
	tx.SetGasOwner(models.SuiAddress(sponsor.GetAddress())).SetGasPrice(1000).SetGasBudget(10_000_000)
	ref, err := NewSuiObjectRef("0x5", "7", "1thX6LZfHDZZGkq4tt1q2yRAPVfCTpX99XN4RHFsxM")
	require.NoError(t, err)
	tx.SetGasPayment([]SuiObjectRef{*ref})
	tx.TransferObjects([]Argument{tx.Gas()}, tx.Pure("0x9"))

	req, err := tx.ToSuiExecuteTransactionBlockRequest(context.Background(), models.SuiTransactionBlockOptions{}, "")
	require.NoError(t, err)
	require.Len(t, req.Signature, 2)

	var td TransactionData
	raw, err := mystenbcs.FromBase64(req.TxBytes)
	require.NoError(t, err)
	_, err = mystenbcs.Unmarshal(raw, &td)
	require.NoError(t, err)
	require.Equal(t, models.SuiAddress(sender.GetAddress()), ConvertSuiAddressBytesToString(*td.V1.Sender))

	for i, want := range []string{sponsor.GetAddress(), sender.GetAddress()} {
		who, pass, err := models.VerifyTransaction(req.TxBytes, req.Signature[i])
		require.NoError(t, err)
		require.True(t, pass)
		require.Equal(t, want, who)
	}
}

func TestSignerSettersReplaceEachOther(t *testing.T) {
	k1, err := signer.NewSecp256k1Signer(append(make([]byte, 31), 1))
	require.NoError(t, err)
	ed := signer.NewSigner(make([]byte, 32))

	tx := NewTransaction().SetKeypairSigner(k1)
	require.Equal(t, signer.Keypair(k1), tx.activeSigner())

	tx.SetSigner(ed)
	require.Nil(t, tx.KeypairSigner)
	require.Equal(t, signer.Keypair(ed), tx.activeSigner())

	// an Ed25519 *Signer passed as a Keypair lands in the legacy field
	tx.SetKeypairSigner(ed)
	require.Equal(t, ed, tx.Signer)
	require.Nil(t, tx.KeypairSigner)

	require.Nil(t, NewTransaction().activeSigner())
	require.Nil(t, NewTransaction().activeSponsoredSigner())
}

// A nil signer pointer (e.g. from a constructor whose error was ignored) must read as
// "not set" and fail with ErrSignerNotSet rather than panic.
func TestNilKeypairPointerIsNotSet(t *testing.T) {
	var k1 *signer.Secp256k1Signer
	var r1 *signer.Secp256r1Signer
	var ed *signer.Signer

	for _, k := range []signer.Keypair{k1, r1, ed, nil} {
		tx := NewTransaction().SetKeypairSigner(k).SetKeypairSponsoredSigner(k)
		require.Nil(t, tx.activeSigner())
		require.Nil(t, tx.activeSponsoredSigner())
		_, err := tx.ToSuiExecuteTransactionBlockRequest(context.Background(), models.SuiTransactionBlockOptions{}, "")
		require.ErrorIs(t, err, ErrSignerNotSet)
		_, err = tx.BuildBCSBytes(context.Background())
		require.ErrorIs(t, err, ErrSignerNotSet)
	}

	// so must a nil pointer assigned to the field directly
	direct := NewTransaction()
	direct.KeypairSigner, direct.KeypairSponsoredSigner = k1, r1
	_, err := direct.ToSuiExecuteTransactionBlockRequest(context.Background(), models.SuiTransactionBlockOptions{}, "")
	require.ErrorIs(t, err, ErrSignerNotSet)

	// setting a nil keypair also clears a previously set one
	real, err := signer.NewSecp256k1Signer(append(make([]byte, 31), 1))
	require.NoError(t, err)
	tx := NewTransaction().SetKeypairSigner(real).SetKeypairSigner(k1)
	require.Nil(t, tx.activeSigner())
}
