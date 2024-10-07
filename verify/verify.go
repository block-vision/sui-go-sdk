package verify

import (
	"errors"
	"fmt"

	"github.com/block-vision/sui-go-sdk/cryptography"
	"github.com/block-vision/sui-go-sdk/cryptography/scheme"
	"github.com/block-vision/sui-go-sdk/keypairs/ed25519"
	"github.com/block-vision/sui-go-sdk/zklogin"
)

func VerifyPersonalMessageSignature(message []byte, signature []byte, options *zklogin.ZkLoginPublicIdentifierOptions) (signer string, pass bool, err error) {
	parsedSignature := parseSignature(signature, options)

	publicKey, err := publicKeyFromRawBytes(parsedSignature.SignatureScheme, parsedSignature.PubKey, options)
	if err != nil {
		return "", false, err
	}

	pass, err = publicKey.VerifyPersonalMessage(message, parsedSignature.Signature, options.Client)
	if err != nil {
		return "", false, err
	}

	address := publicKey.ToSuiAddress()

	return address, true, nil
}

func parseSignature(signature []byte, options *zklogin.ZkLoginPublicIdentifierOptions) *cryptography.SignaturePubkeyPair {
	return nil
}

// publicKeyFromRawBytes function in Go
func publicKeyFromRawBytes(signatureScheme scheme.SignatureScheme, bytes []byte, options *zklogin.ZkLoginPublicIdentifierOptions) (IPublicKey, error) {
	switch signatureScheme {
	case scheme.ED25519:
		return ed25519.NewEd25519PublicKey(bytes), nil
	case scheme.ZkLogin:
		return zklogin.NewZkLoginPublicIdentifier(bytes, options), nil
	default:
		return nil, errors.New(fmt.Sprintf("Unsupported signature scheme %s", signatureScheme))
	}
}
