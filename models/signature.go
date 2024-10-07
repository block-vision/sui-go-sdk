package models

import (
	"crypto"
	"crypto/ed25519"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"log"
	"strings"

	"golang.org/x/crypto/blake2b"

	"github.com/block-vision/sui-go-sdk/constant"
)

type InputObjectKind map[string]interface{}
type ObjectId = HexData
type Digest = Base64Data

type ObjectRef struct {
	Digest   string   `json:"digest"`
	ObjectId ObjectId `json:"objectId"`
	Version  uint64   `json:"version"`
}

type SigScheme string

const (
	SigEd25519   SigScheme = "ED25519"
	SigSecp256k1 SigScheme = "Secp256k1"
)

type SigFlag byte

const (
	SigFlagEd25519   SigFlag = 0x00
	SigFlagSecp256k1 SigFlag = 0x01
)

type HexData struct {
	data []byte
}

type SignaturePubkeyPair struct {
	SignatureScheme string
	Signature       []byte
	PubKey          []byte
}

func NewHexData(str string) (*HexData, error) {
	if strings.HasPrefix(str, "0x") || strings.HasPrefix(str, "0X") {
		str = str[2:]
	}
	data, err := hex.DecodeString(str)
	if err != nil {
		return nil, err
	}
	return &HexData{data}, nil
}

func (a HexData) Data() []byte {
	return a.data
}

type Bytes []byte

func (b Bytes) GetHexData() HexData {
	return HexData{b}
}
func (b Bytes) GetBase64Data() Base64Data {
	return Base64Data{b}
}

type Base64Data struct {
	data []byte
}

func NewBase64Data(str string) (*Base64Data, error) {
	data, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return nil, err
	}
	return &Base64Data{data}, nil
}

func (h Base64Data) Data() []byte {
	return h.data
}

type SignedTransaction struct {
	// transaction data bytes
	TxBytes string `json:"tx_bytes"`

	// Flag of the signature scheme that is used.
	SigScheme SigScheme `json:"sig_scheme"`

	// transaction signature
	Signature *Base64Data `json:"signature"`

	// signer's public key
	PublicKey *Base64Data `json:"pub_key"`
}

type SignedTransactionSerializedSig struct {
	// transaction data bytes
	TxBytes string `json:"tx_bytes"`

	// transaction signature
	Signature string `json:"signature"`
}

var IntentBytes = []byte{0, 0, 0}

func (txn *TxnMetaData) SignSerializedSigWith(privateKey ed25519.PrivateKey) *SignedTransactionSerializedSig {
	txBytes, _ := base64.StdEncoding.DecodeString(txn.TxBytes)
	message := messageWithIntent(txBytes)
	digest := blake2b.Sum256(message)
	var noHash crypto.Hash
	sigBytes, err := privateKey.Sign(nil, digest[:], noHash)
	if err != nil {
		log.Fatal(err)
	}
	return &SignedTransactionSerializedSig{
		TxBytes:   txn.TxBytes,
		Signature: ToSerializedSignature(sigBytes, privateKey.Public().(ed25519.PublicKey)),
	}
}

func messageWithIntent(message []byte) []byte {
	intent := IntentBytes
	intentMessage := make([]byte, len(intent)+len(message))
	copy(intentMessage, intent)
	copy(intentMessage[len(intent):], message)
	return intentMessage
}

func ToSerializedSignature(signature, pubKey []byte) string {
	signatureLen := len(signature)
	pubKeyLen := len(pubKey)
	serializedSignature := make([]byte, 1+signatureLen+pubKeyLen)
	serializedSignature[0] = byte(SigFlagEd25519)
	copy(serializedSignature[1:], signature)
	copy(serializedSignature[1+signatureLen:], pubKey)
	return base64.StdEncoding.EncodeToString(serializedSignature)
}

func FromSerializedSignature(serializedSignature string) (*SignaturePubkeyPair, error) {
	_bytes, err := base64.StdEncoding.DecodeString(serializedSignature)
	if err != nil {
		return nil, err
	}
	signatureScheme := parseSignatureScheme(_bytes[0])
	if strings.EqualFold(serializedSignature, "") {
		return nil, fmt.Errorf("multiSig is not supported")
	}

	signature := _bytes[1 : len(_bytes)-32]
	pubKeyBytes := _bytes[1+len(signature):]

	keyPair := &SignaturePubkeyPair{
		SignatureScheme: signatureScheme,
		Signature:       signature,
		PubKey:          pubKeyBytes,
	}
	return keyPair, nil
}

func parseSignatureScheme(scheme byte) string {
	switch scheme {
	case 0:
		return "ED25519"
	case 1:
		return "Secp256k1"
	case 2:
		return "Secp256r1"
	case 3:
		return "MultiSig"
	case 5:
		return "ZkLogin"
	default:
		return "ED25519"
	}
}

func VerifyPersonalMessage(message string, signature string) (signer string, pass bool, err error) {
	b64Message := base64.StdEncoding.EncodeToString([]byte(message))
	return VerifyMessage(b64Message, signature, constant.PersonalMessageIntentScope)
}

func VerifyTransaction(b64Message string, signature string) (signer string, pass bool, err error) {
	return VerifyMessage(b64Message, signature, constant.TransactionDataIntentScope)
}

func VerifyMessage(message, signature string, scope constant.IntentScope) (signer string, pass bool, err error) {
	b64Bytes, _ := base64.StdEncoding.DecodeString(message)
	messageBytes := NewMessageWithIntent(b64Bytes, scope)

	serializedSignature, err := FromSerializedSignature(signature)
	if err != nil {
		return "", false, err
	}
	digest := blake2b.Sum256(messageBytes)

	pass = ed25519.Verify(serializedSignature.PubKey[:], digest[:], serializedSignature.Signature)

	signer = Ed25519PublicKeyToSuiAddress(serializedSignature.PubKey)
	if err != nil {
		return "", false, fmt.Errorf("invalid signer %v", err)
	}

	return
}

func Ed25519PublicKeyToSuiAddress(pubKey []byte) string {
	newPubkey := []byte{byte(SigFlagEd25519)}
	newPubkey = append(newPubkey, pubKey...)

	addrBytes := blake2b.Sum256(newPubkey)
	return fmt.Sprintf("0x%s", hex.EncodeToString(addrBytes[:])[:64])
}
