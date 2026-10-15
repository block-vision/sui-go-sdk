package transaction

import (
	"bytes"
	"errors"
	"os"
	"testing"

	"github.com/mr-tron/base58"
	"github.com/samber/lo"
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/blake2b"

	"github.com/block-vision/sui-go-sdk/models"
	"github.com/block-vision/sui-go-sdk/mystenbcs"
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
		{"u8", []byte{0x01}},
		{"u64", []byte{0x0a}},
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
