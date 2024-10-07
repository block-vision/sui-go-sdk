package cryptography

import (
	"fmt"
	"strings"

	"github.com/block-vision/sui-go-sdk/cryptography/scheme"
	"github.com/block-vision/sui-go-sdk/mystenbcs"
	"github.com/block-vision/sui-go-sdk/zklogin"
)

type SignaturePubkeyPair struct {
	SignatureScheme scheme.SignatureScheme
	Signature       []byte
	PubKey          []byte
}

func parseSerializedSignature(serializedSignature string) (*SignaturePubkeyPair, error) {
	if strings.EqualFold(serializedSignature, "") {
		return nil, fmt.Errorf("multiSig is not supported")
	}

	bytes, err := mystenbcs.FromBase64(serializedSignature)
	if err != nil {
		return nil, err
	}

	signatureScheme, ok := scheme.SignatureFlagToScheme[bytes[0]]
	if !ok {
		return nil, fmt.Errorf("signature flag is not supported")
	}

	switch signatureScheme {
	case "ZkLogin":
		parsedSerializedZkLoginSignature, err := zklogin.ParseSerializedZkLoginSignature(serializedSignature)
		if err != nil {
			return nil, err
		}

		return &SignaturePubkeyPair{
			SignatureScheme: parsedSerializedZkLoginSignature.SignatureScheme,
			Signature:       parsedSerializedZkLoginSignature.Signature,
			PubKey:          parsedSerializedZkLoginSignature.PubKey,
		}, nil
	case "ED25519":
		fallthrough
	case "Secp256k1":
		fallthrough
	case "Secp256r1":
		size, ok := scheme.SignatureSchemeToSize[signatureScheme]
		if !ok {
			return nil, fmt.Errorf("signature scheme is not supported")
		}

		signature := bytes[1 : len(bytes)-size]
		pubKeyBytes := bytes[1+len(signature):]

		keyPair := &SignaturePubkeyPair{
			SignatureScheme: signatureScheme,
			Signature:       signature,
			PubKey:          pubKeyBytes,
		}

		return keyPair, nil
	default:
		return nil, fmt.Errorf("signature scheme is not supported")
	}
}
