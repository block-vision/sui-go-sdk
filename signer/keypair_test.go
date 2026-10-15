package signer

import (
	"crypto/elliptic"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/big"
	"os"
	"strings"
	"testing"

	"github.com/block-vision/sui-go-sdk/constant"
	"github.com/block-vision/sui-go-sdk/cryptography/scheme"
	"github.com/block-vision/sui-go-sdk/models"
	"github.com/decred/dcrd/dcrec/secp256k1/v4"
	"github.com/stretchr/testify/require"
)

// The expected values below were checked against a Sui mainnet fullnode
// (SignatureVerificationService.VerifySignature with the address).
var scalarOne = append(make([]byte, 31), 1)

func scalar(n *big.Int) []byte {
	b := make([]byte, 32)
	n.FillBytes(b)
	return b
}

func testKeypairs(t *testing.T) []Keypair {
	t.Helper()
	k1, err := NewSecp256k1Signer(scalarOne)
	require.NoError(t, err)
	r1, err := NewSecp256r1Signer(scalarOne)
	require.NoError(t, err)
	return []Keypair{NewSigner(scalarOne), k1, r1}
}

func TestKeypairPublicKeyAndAddress(t *testing.T) {
	want := map[scheme.SignatureScheme][2]string{
		scheme.ED25519:   {"4cb5abf6ad79fbf5abbccafcc269d85cd2651ed4b885b5869f241aedf0a5ba29", "0xd0c2c91eda34bbfbaec6cfb9c7bb913e57dab3cbec4018a4b3f5e55531cd63af"},
		scheme.Secp256k1: {"0279be667ef9dcbbac55a06295ce870b07029bfcdb2dce28d959f2815b16f81798", "0xd4c3524e6642b2e54945c02378024f822ac3f80b0870a5f95f06e68a61890a6c"},
		scheme.Secp256r1: {"036b17d1f2e12c4247f8bce6e563a440f277037d812deb33a0f4a13945d898c296", "0x173e0d2ec575814f055dee0c3c0ce1357c9f3d58a8019b04dd369ca263f6db55"},
	}
	for _, k := range testKeypairs(t) {
		t.Run(string(k.SignatureScheme()), func(t *testing.T) {
			require.Equal(t, want[k.SignatureScheme()][0], hex.EncodeToString(k.PublicKeyBytes()))
			require.Equal(t, want[k.SignatureScheme()][1], k.GetAddress())
		})
	}
}

// Secp256k1 signing is RFC 6979 deterministic, so the node-accepted signature is fixed.
func TestSecp256k1SignatureIsDeterministic(t *testing.T) {
	k1, err := NewSecp256k1Signer(scalarOne)
	require.NoError(t, err)
	sig, err := k1.SignPersonalMessageV1("hello sui")
	require.NoError(t, err)
	require.Equal(t, "AVo6+fjs9ENspMqTE2c5paTHAfsmRx38Pp+fwCeOEMleEjTDoUjAMVP6Pn9yTzWX+X6/7pq2emQLStEz1rPrQAwCeb5mfvncu6xVoGKVzocLBwKb/NstzijZWfKBWxb4F5g=", sig.Signature)
}

func TestSignAndVerifyAllSchemes(t *testing.T) {
	txBytes, err := os.ReadFile("testdata/3sxUyUzNE2HJjxmM6K9abExqS1Vj1ca1ETTSzQg7dxFK.bcs")
	require.NoError(t, err)
	b64Tx := base64.StdEncoding.EncodeToString(txBytes)

	for _, k := range testKeypairs(t) {
		t.Run(string(k.SignatureScheme()), func(t *testing.T) {
			tx, err := k.SignTransaction(b64Tx)
			require.NoError(t, err)
			signer, pass, err := models.VerifyTransaction(b64Tx, tx.Signature)
			require.NoError(t, err)
			require.True(t, pass)
			require.Equal(t, k.GetAddress(), signer)

			viaScope, err := k.SignMessage(b64Tx, constant.TransactionDataIntentScope)
			require.NoError(t, err)
			_, pass, err = models.VerifyTransaction(b64Tx, viaScope.Signature)
			require.NoError(t, err)
			require.True(t, pass)

			pm, err := k.SignPersonalMessageV1("hello sui")
			require.NoError(t, err)
			signer, pass, err = models.VerifyPersonalMessage("hello sui", pm.Signature)
			require.NoError(t, err)
			require.True(t, pass)
			require.Equal(t, k.GetAddress(), signer)

			_, pass, err = models.VerifyPersonalMessage("hello sui!", pm.Signature)
			require.NoError(t, err)
			require.False(t, pass, "tampered message")
			_, pass, err = models.VerifyTransaction(b64Tx, pm.Signature)
			require.NoError(t, err)
			require.False(t, pass, "a personal message signature is not a transaction signature")
		})
	}
}

// A real mainnet transaction signed with Secp256k1.
func TestVerifyMainnetSecp256k1TransactionSignature(t *testing.T) {
	txBytes, err := os.ReadFile("testdata/3sxUyUzNE2HJjxmM6K9abExqS1Vj1ca1ETTSzQg7dxFK.bcs")
	require.NoError(t, err)
	signer, pass, err := models.VerifyTransaction(base64.StdEncoding.EncodeToString(txBytes),
		"Ad1bHgkqNWzPXfV19dkpoz56rcE68EKgAFzum+1buHeZTwnRu3/qwCHj8BnhyfBzMG6QB2tjAkYjEJfHfTX+xpwCleiEASmtiWhyRRjSX0Xf3Cnvepjv2AeZwJjfY7n3if4=")
	require.NoError(t, err)
	require.True(t, pass)
	require.Equal(t, "0xe98aaadcdf0dfdbaf182b91269566adc413401c423229bfc4b690e884e25e3ee", signer)
}

// Node-accepted personal message signatures for the scalar-one keys.
func TestVerifyNodeAcceptedPersonalMessageSignatures(t *testing.T) {
	for sig, address := range map[string]string{
		"AGeJyOvOAF+a1EQMsd882esZN+PBKYDwMmiHgZopQex6JyAgHta3NduT6cZmxn5xWFX+jYFwI2uZP5MxuVOi8wRMtav2rXn79au8yvzCadhc0mUe1LiFtYafJBrt8KW6KQ==": "0xd0c2c91eda34bbfbaec6cfb9c7bb913e57dab3cbec4018a4b3f5e55531cd63af",
		"Av8m9x6BqsswolGzqa+OGKCJ9BZYvtM3C+2AFbXua+3vGYdnuRoaqnHSz6pqeiO/X6I0eq10OvfHHBrL+5VD8bUDaxfR8uEsQkf4vOblY6RA8ncDfYEt6zOg9KE5RdiYwpY=": "0x173e0d2ec575814f055dee0c3c0ce1357c9f3d58a8019b04dd369ca263f6db55",
	} {
		signer, pass, err := models.VerifyPersonalMessage("hello sui", sig)
		require.NoError(t, err)
		require.True(t, pass)
		require.Equal(t, address, signer)
	}
}

// Validators (fastcrypto) reject the high-S twin of both Secp256k1 and Secp256r1
// signatures; only the gRPC VerifySignature endpoint accepts it for Secp256r1.
func TestHighSSignatures(t *testing.T) {
	flipS := func(sig string, order *big.Int) string {
		b, err := base64.StdEncoding.DecodeString(sig)
		require.NoError(t, err)
		s := new(big.Int).SetBytes(b[33:65])
		new(big.Int).Sub(order, s).FillBytes(b[33:65])
		return base64.StdEncoding.EncodeToString(b)
	}
	k1, err := NewSecp256k1Signer(scalarOne)
	require.NoError(t, err)
	r1, err := NewSecp256r1Signer(scalarOne)
	require.NoError(t, err)

	k1Sig, err := k1.SignPersonalMessageV1("hello sui")
	require.NoError(t, err)
	_, pass, err := models.VerifyPersonalMessage("hello sui", flipS(k1Sig.Signature, secp256k1.Params().N))
	require.NoError(t, err)
	require.False(t, pass)

	r1Sig, err := r1.SignPersonalMessageV1("hello sui")
	require.NoError(t, err)
	_, pass, err = models.VerifyPersonalMessage("hello sui", flipS(r1Sig.Signature, elliptic.P256().Params().N))
	require.NoError(t, err)
	require.False(t, pass)
}

// The signers themselves always produce low-S signatures.
func TestSignaturesAreLowS(t *testing.T) {
	r1, err := NewSecp256r1Signer(scalarOne)
	require.NoError(t, err)
	half := new(big.Int).Rsh(elliptic.P256().Params().N, 1)
	for i := 0; i < 64; i++ {
		sig, err := r1.SignPersonalMessageV1(string(rune('a' + i)))
		require.NoError(t, err)
		b, _ := base64.StdEncoding.DecodeString(sig.Signature)
		require.LessOrEqual(t, new(big.Int).SetBytes(b[33:65]).Cmp(half), 0)
	}
}

func TestSuiPrivateKeyRoundTrip(t *testing.T) {
	for _, k := range testKeypairs(t) {
		t.Run(string(k.SignatureScheme()), func(t *testing.T) {
			encoded, err := k.ExportSuiPrivateKey()
			require.NoError(t, err)

			keyScheme, secret, err := DecodeSuiPrivateKey(encoded)
			require.NoError(t, err)
			require.Equal(t, k.SignatureScheme(), keyScheme)
			require.Equal(t, scalarOne, secret)

			back, err := NewKeypairFromSuiPrivateKey(encoded)
			require.NoError(t, err)
			require.Equal(t, k.SignatureScheme(), back.SignatureScheme())
			require.Equal(t, k.GetAddress(), back.GetAddress())

			_, err = NewSignerWithSecretKey(encoded)
			if k.SignatureScheme() == scheme.ED25519 {
				require.NoError(t, err)
			} else {
				require.Error(t, err, "a non-Ed25519 key must not be read as an Ed25519 seed")
			}
		})
	}
}

func TestInvalidPrivateKeys(t *testing.T) {
	for name, key := range map[string][]byte{
		"zero":     make([]byte, 32),
		"k1 order": scalar(secp256k1.Params().N),
		"31 bytes": make([]byte, 31),
		"33 bytes": make([]byte, 33),
	} {
		_, err := NewSecp256k1Signer(key)
		require.Error(t, err, "secp256k1 "+name)
	}
	for name, key := range map[string][]byte{
		"zero":     make([]byte, 32),
		"r1 order": scalar(elliptic.P256().Params().N),
		"31 bytes": make([]byte, 31),
	} {
		_, err := NewSecp256r1Signer(key)
		require.Error(t, err, "secp256r1 "+name)
	}

	_, err := EncodeSuiPrivateKey(scheme.MultiSig, scalarOne)
	require.Error(t, err)
	_, _, err = DecodeSuiPrivateKey("notsui1qyqszqgpqyqszqgpqyqszqgpqyqszqgpqyqszqgpqyqszqgpqyqsrrzg6c")
	require.Error(t, err)
}

// Logging or serializing a Secp signer must never expose the private key.
func TestSecpSignersDoNotExposePrivateKey(t *testing.T) {
	secret, err := hex.DecodeString("1c7e2b3a4d5f60718293a4b5c6d7e8f90a1b2c3d4e5f60718293a4b5c6d7e8f9")
	require.NoError(t, err)
	k1, err := NewSecp256k1Signer(secret)
	require.NoError(t, err)
	r1, err := NewSecp256r1Signer(secret)
	require.NoError(t, err)
	for _, k := range []Keypair{k1, r1} {
		for _, f := range []string{"%v", "%+v", "%#v", "%s"} {
			out := fmt.Sprintf(f, k)
			require.NotContains(t, strings.ToLower(out), "1c7e2b3a", f)
			require.NotContains(t, out, "[28 126 43 58", f)
			require.NotContains(t, out, "0x1c, 0x7e", f)
		}
		j, err := json.Marshal(k)
		require.NoError(t, err)
		require.Equal(t, "{}", string(j))
	}
	require.NotContains(t, fmt.Sprintf("%+v", r1), r1.privateKey.D.String())
}
