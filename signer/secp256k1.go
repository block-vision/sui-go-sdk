package signer

import (
	"crypto/sha256"
	"fmt"

	"github.com/block-vision/sui-go-sdk/cryptography/scheme"
	"github.com/decred/dcrd/dcrec/secp256k1/v4"
	"github.com/decred/dcrd/dcrec/secp256k1/v4/ecdsa"
)

// Secp256k1Signer signs with a Secp256k1 key (Sui signature flag 0x01).
type Secp256k1Signer struct {
	ecdsaKeypair
	// unexported so fmt and encoding/json never print the private key
	privateKey *secp256k1.PrivateKey
}

// NewSecp256k1Signer creates a signer from a 32-byte big-endian private scalar in [1, n-1].
func NewSecp256k1Signer(secretKey []byte) (*Secp256k1Signer, error) {
	if len(secretKey) != 32 {
		return nil, fmt.Errorf("invalid secp256k1 private key length: %d", len(secretKey))
	}
	var scalar secp256k1.ModNScalar
	if overflow := scalar.SetByteSlice(secretKey); overflow || scalar.IsZero() {
		return nil, fmt.Errorf("invalid secp256k1 private key: scalar out of range")
	}
	privKey := secp256k1.NewPrivateKey(&scalar)
	pubKey := privKey.PubKey().SerializeCompressed()

	s := &Secp256k1Signer{privateKey: privKey}
	s.ecdsaKeypair = ecdsaKeypair{
		scheme:  scheme.Secp256k1,
		pubKey:  pubKey,
		address: suiAddress(scheme.SignatureSchemeToFlag[scheme.Secp256k1], pubKey),
		signDigest: func(digest [32]byte) ([]byte, error) {
			hash := sha256.Sum256(digest[:])
			// RFC 6979 deterministic and BIP 62 canonical (low-S), as Sui requires.
			sig := ecdsa.Sign(privKey, hash[:])
			r, sv := sig.R(), sig.S()
			out := make([]byte, 64)
			r.PutBytesUnchecked(out[:32])
			sv.PutBytesUnchecked(out[32:])
			return out, nil
		},
	}
	return s, nil
}

// ExportSuiPrivateKey returns the key as a Bech32 "suiprivkey…" string.
func (s *Secp256k1Signer) ExportSuiPrivateKey() (string, error) {
	return EncodeSuiPrivateKey(scheme.Secp256k1, s.privateKey.Serialize())
}

// String identifies the signer without exposing key material.
func (s *Secp256k1Signer) String() string {
	return fmt.Sprintf("Secp256k1Signer(%s)", s.address)
}
