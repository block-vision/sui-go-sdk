package keypair

import (
	"encoding/base64"
	"math"

	"github.com/block-vision/sui-go-sdk/common/sui_error"

	"github.com/block-vision/sui-go-sdk/models"
)

type KeyPair byte

const (
	Ed25519Flag   KeyPair = 0
	Secp256k1Flag KeyPair = 1
	ErrorFlag     byte    = math.MaxUint8
)

const (
	ed25519PublicKeyLength   = 32
	secp256k1PublicKeyLength = 33
)

const (
	DefaultAccountAddressLength = 16
	AccountAddress20Length      = 20
	AccountAddress32Length      = 32
)

func FetchKeyPair(value string) (models.SuiKeyPair, error) {
	result, err := base64.StdEncoding.DecodeString(value)
	if err != nil {
		return models.SuiKeyPair{}, err
	}
	if len(result) == 0 {
		return models.SuiKeyPair{}, err
	}
	switch result[0] {
	case byte(Ed25519Flag):
		pb := result[1 : ed25519PublicKeyLength+1]
		sk := result[1+ed25519PublicKeyLength:]
		pbInBase64 := encodeBase64(pb)
		return models.SuiKeyPair{
			Flag:            byte(Ed25519Flag),
			PrivateKey:      sk,
			PublicKeyBase64: pbInBase64,
			PublicKey:       pb,
			Address:         fromPublicKeyBytesToAddress(pb, byte(Ed25519Flag)),
		}, nil
	case byte(Secp256k1Flag):
		pb := result[1 : secp256k1PublicKeyLength+1]
		sk := result[1+secp256k1PublicKeyLength:]
		pbInBase64 := encodeBase64(pb)
		return models.SuiKeyPair{
			Flag:            byte(Secp256k1Flag),
			PrivateKey:      sk,
			PublicKey:       pb,
			PublicKeyBase64: pbInBase64,
			Address:         fromPublicKeyBytesToAddress(pb, byte(Secp256k1Flag)),
		}, nil
	default:
		return models.SuiKeyPair{}, sui_error.ErrInvalidEncryptFlag
	}
}
