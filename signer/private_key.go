package signer

import (
	"fmt"

	"github.com/block-vision/sui-go-sdk/cryptography/scheme"
	"github.com/btcsuite/btcutil/bech32"
)

const suiPrivateKeyPrefix = "suiprivkey"

// DecodeSuiPrivateKey decodes a Bech32 "suiprivkey…" string (as exported by the Sui CLI
// and wallets) into its signature scheme and the 32-byte private key.
func DecodeSuiPrivateKey(value string) (scheme.SignatureScheme, []byte, error) {
	hrp, data, err := bech32.Decode(value)
	if err != nil {
		return "", nil, err
	}
	if hrp != suiPrivateKeyPrefix {
		return "", nil, fmt.Errorf("invalid bech32 prefix: %s", hrp)
	}
	decoded, err := bech32.ConvertBits(data, 5, 8, false)
	if err != nil {
		return "", nil, err
	}
	if len(decoded) != 33 {
		return "", nil, fmt.Errorf("invalid sui private key length: %d", len(decoded))
	}
	keyScheme, ok := scheme.SignatureFlagToScheme[decoded[0]]
	if !ok || (keyScheme != scheme.ED25519 && keyScheme != scheme.Secp256k1 && keyScheme != scheme.Secp256r1) {
		return "", nil, fmt.Errorf("unsupported private key signature flag: 0x%02x", decoded[0])
	}
	return keyScheme, decoded[1:], nil
}

// EncodeSuiPrivateKey encodes a 32-byte private key as a Bech32 "suiprivkey…" string.
func EncodeSuiPrivateKey(keyScheme scheme.SignatureScheme, secretKey []byte) (string, error) {
	flag, ok := scheme.SignatureSchemeToFlag[keyScheme]
	if !ok || (keyScheme != scheme.ED25519 && keyScheme != scheme.Secp256k1 && keyScheme != scheme.Secp256r1) {
		return "", fmt.Errorf("unsupported private key scheme: %s", keyScheme)
	}
	if len(secretKey) != 32 {
		return "", fmt.Errorf("invalid private key length: %d", len(secretKey))
	}
	data, err := bech32.ConvertBits(append([]byte{flag}, secretKey...), 8, 5, true)
	if err != nil {
		return "", err
	}
	return bech32.Encode(suiPrivateKeyPrefix, data)
}

// NewKeypairFromSuiPrivateKey creates an Ed25519, Secp256k1 or Secp256r1 signer from a
// Bech32 "suiprivkey…" string, according to the scheme flag it carries.
func NewKeypairFromSuiPrivateKey(value string) (Keypair, error) {
	keyScheme, secretKey, err := DecodeSuiPrivateKey(value)
	if err != nil {
		return nil, err
	}
	switch keyScheme {
	case scheme.ED25519:
		return NewSigner(secretKey), nil
	case scheme.Secp256k1:
		return NewSecp256k1Signer(secretKey)
	case scheme.Secp256r1:
		return NewSecp256r1Signer(secretKey)
	}
	return nil, fmt.Errorf("unsupported private key scheme: %s", keyScheme)
}
