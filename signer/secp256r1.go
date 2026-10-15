package signer

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"math/big"

	"github.com/block-vision/sui-go-sdk/cryptography/scheme"
)

// Secp256r1Signer signs with a Secp256r1 (NIST P-256) key (Sui signature flag 0x02).
type Secp256r1Signer struct {
	ecdsaKeypair
	// unexported so fmt and encoding/json never print the private key
	privateKey *ecdsa.PrivateKey
}

// NewSecp256r1Signer creates a signer from a 32-byte big-endian private scalar in [1, n-1].
func NewSecp256r1Signer(secretKey []byte) (*Secp256r1Signer, error) {
	privKey, err := ecdsa.ParseRawPrivateKey(elliptic.P256(), secretKey)
	if err != nil {
		return nil, fmt.Errorf("invalid secp256r1 private key: %w", err)
	}
	uncompressed, err := privKey.PublicKey.Bytes()
	if err != nil {
		return nil, err
	}
	pubKey := compressP256(uncompressed)
	halfOrder := new(big.Int).Rsh(elliptic.P256().Params().N, 1)
	order := elliptic.P256().Params().N

	s := &Secp256r1Signer{privateKey: privKey}
	s.ecdsaKeypair = ecdsaKeypair{
		scheme:  scheme.Secp256r1,
		pubKey:  pubKey,
		address: suiAddress(scheme.SignatureSchemeToFlag[scheme.Secp256r1], pubKey),
		signDigest: func(digest [32]byte) ([]byte, error) {
			hash := sha256.Sum256(digest[:])
			r, sv, err := ecdsa.Sign(rand.Reader, privKey, hash[:])
			if err != nil {
				return nil, err
			}
			// Sui only accepts the low-S form.
			if sv.Cmp(halfOrder) > 0 {
				sv = new(big.Int).Sub(order, sv)
			}
			out := make([]byte, 64)
			r.FillBytes(out[:32])
			sv.FillBytes(out[32:])
			return out, nil
		},
	}
	return s, nil
}

// ExportSuiPrivateKey returns the key as a Bech32 "suiprivkey…" string.
func (s *Secp256r1Signer) ExportSuiPrivateKey() (string, error) {
	raw, err := s.privateKey.Bytes()
	if err != nil {
		return "", err
	}
	return EncodeSuiPrivateKey(scheme.Secp256r1, raw)
}

// compressP256 turns a 65-byte SEC 1 uncompressed point into its 33-byte compressed form.
func compressP256(uncompressed []byte) []byte {
	out := make([]byte, 33)
	out[0] = 0x02 | uncompressed[64]&1
	copy(out[1:], uncompressed[1:33])
	return out
}

// String identifies the signer without exposing key material.
func (s *Secp256r1Signer) String() string {
	return fmt.Sprintf("Secp256r1Signer(%s)", s.address)
}
