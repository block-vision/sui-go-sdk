package models

import (
	"bytes"
	"crypto"
	"crypto/ecdsa"
	"crypto/ed25519"
	"crypto/elliptic"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"log"
	"math/big"
	"strings"

	"golang.org/x/crypto/blake2b"

	"github.com/block-vision/sui-go-sdk/constant"
	"github.com/block-vision/sui-go-sdk/mystenbcs"
	"github.com/decred/dcrd/dcrec/secp256k1/v4"
	secp256k1ecdsa "github.com/decred/dcrd/dcrec/secp256k1/v4/ecdsa"
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
	SigFlagSecp256r1 SigFlag = 0x02
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

// ToSerializedSignature serializes an Ed25519 signature as flag || signature || public key.
func ToSerializedSignature(signature, pubKey []byte) string {
	return ToSerializedSignatureWithScheme(signature, pubKey, byte(SigFlagEd25519))
}

// ToSerializedSignatureWithScheme serializes a signature as flag || signature || public key
// for the given Sui signature scheme flag.
func ToSerializedSignatureWithScheme(signature, pubKey []byte, sigFlag byte) string {
	serializedSignature := make([]byte, 0, 1+len(signature)+len(pubKey))
	serializedSignature = append(serializedSignature, sigFlag)
	serializedSignature = append(serializedSignature, signature...)
	serializedSignature = append(serializedSignature, pubKey...)
	return base64.StdEncoding.EncodeToString(serializedSignature)
}

// simpleSignatureSchemes maps the single-key signature flags to their scheme name and
// public key length; every one of them has a 64-byte signature.
var simpleSignatureSchemes = map[byte]struct {
	name      string
	pubKeyLen int
}{
	byte(SigFlagEd25519):   {"ED25519", ed25519.PublicKeySize},
	byte(SigFlagSecp256k1): {"Secp256k1", 33},
	byte(SigFlagSecp256r1): {"Secp256r1", 33},
}

// FromSerializedSignature parses a flag || signature || public key signature. Only the
// single-key schemes (Ed25519, Secp256k1, Secp256r1) are supported.
func FromSerializedSignature(serializedSignature string) (*SignaturePubkeyPair, error) {
	_bytes, err := base64.StdEncoding.DecodeString(serializedSignature)
	if err != nil {
		return nil, err
	}
	if len(_bytes) == 0 {
		return nil, fmt.Errorf("serialized signature is empty")
	}
	sigScheme, ok := simpleSignatureSchemes[_bytes[0]]
	if !ok {
		return nil, fmt.Errorf("signature flag 0x%02x (%s) is not supported", _bytes[0], parseSignatureScheme(_bytes[0]))
	}
	if want := 1 + 64 + sigScheme.pubKeyLen; len(_bytes) != want {
		return nil, fmt.Errorf("invalid %s serialized signature length: %d, want %d", sigScheme.name, len(_bytes), want)
	}
	return &SignaturePubkeyPair{
		SignatureScheme: sigScheme.name,
		Signature:       _bytes[1:65],
		PubKey:          _bytes[65:],
	}, nil
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
	case 6:
		return "Passkey"
	default:
		return "unknown"
	}
}

func VerifyPersonalMessage(message string, signature string) (signer string, pass bool, err error) {
	b64Message := base64.StdEncoding.EncodeToString([]byte(message))
	return VerifyMessage(b64Message, signature, constant.PersonalMessageIntentScope)
}

func VerifyTransaction(b64Message string, signature string) (signer string, pass bool, err error) {
	return VerifyMessage(b64Message, signature, constant.TransactionDataIntentScope)
}

// VerifyMessage verifies an Ed25519, Secp256k1 or Secp256r1 signature over the base64
// message and returns the signer address. Personal messages are signed over their BCS
// encoding (vector<u8>), every other intent over the raw bytes, as on Sui.
func VerifyMessage(message, signature string, scope constant.IntentScope) (signer string, pass bool, err error) {
	payload, err := base64.StdEncoding.DecodeString(message)
	if err != nil {
		return "", false, err
	}
	if scope == constant.PersonalMessageIntentScope {
		bcsEncodedMsg := bytes.Buffer{}
		if err := mystenbcs.NewEncoder(&bcsEncodedMsg).Encode(payload); err != nil {
			return "", false, err
		}
		payload = bcsEncodedMsg.Bytes()
	}

	serializedSignature, err := FromSerializedSignature(signature)
	if err != nil {
		return "", false, err
	}
	digest := blake2b.Sum256(NewMessageWithIntent(payload, scope))

	pass, err = verifyDigest(serializedSignature, digest)
	if err != nil {
		return "", false, err
	}
	flag := byte(SigFlagEd25519)
	for f, sigScheme := range simpleSignatureSchemes {
		if sigScheme.name == serializedSignature.SignatureScheme {
			flag = f
		}
	}
	return PublicKeyToSuiAddress(serializedSignature.PubKey, flag), pass, nil
}

// verifyDigest checks the signature over the blake2b intent digest. ECDSA schemes sign
// SHA-256(digest) and, as validators require when executing transactions, only the
// low-S form is accepted.
func verifyDigest(sig *SignaturePubkeyPair, digest [32]byte) (bool, error) {
	switch sig.SignatureScheme {
	case "ED25519":
		return ed25519.Verify(sig.PubKey, digest[:], sig.Signature), nil
	case "Secp256k1":
		pubKey, err := secp256k1.ParsePubKey(sig.PubKey)
		if err != nil {
			return false, fmt.Errorf("invalid Secp256k1 public key: %w", err)
		}
		var r, s secp256k1.ModNScalar
		if r.SetByteSlice(sig.Signature[:32]) || s.SetByteSlice(sig.Signature[32:]) || r.IsZero() || s.IsZero() {
			return false, nil
		}
		if s.IsOverHalfOrder() {
			return false, nil
		}
		hash := sha256.Sum256(digest[:])
		return secp256k1ecdsa.NewSignature(&r, &s).Verify(hash[:], pubKey), nil
	case "Secp256r1":
		x, y := elliptic.UnmarshalCompressed(elliptic.P256(), sig.PubKey)
		if x == nil {
			return false, fmt.Errorf("invalid Secp256r1 public key")
		}
		r, s := new(big.Int).SetBytes(sig.Signature[:32]), new(big.Int).SetBytes(sig.Signature[32:])
		// Validators (fastcrypto) reject the high-S form. The gRPC VerifySignature endpoint
		// (sui-crypto) accepts it, but such a transaction would not execute.
		if s.Cmp(new(big.Int).Rsh(elliptic.P256().Params().N, 1)) > 0 {
			return false, nil
		}
		hash := sha256.Sum256(digest[:])
		return ecdsa.Verify(&ecdsa.PublicKey{Curve: elliptic.P256(), X: x, Y: y}, hash[:], r, s), nil
	}
	return false, fmt.Errorf("signature scheme %s is not supported", sig.SignatureScheme)
}

func Ed25519PublicKeyToSuiAddress(pubKey []byte) string {
	return PublicKeyToSuiAddress(pubKey, byte(SigFlagEd25519))
}

// PublicKeyToSuiAddress derives the Sui address blake2b256(flag || public key).
func PublicKeyToSuiAddress(pubKey []byte, sigFlag byte) string {
	newPubkey := []byte{sigFlag}
	newPubkey = append(newPubkey, pubKey...)

	addrBytes := blake2b.Sum256(newPubkey)
	return fmt.Sprintf("0x%s", hex.EncodeToString(addrBytes[:])[:64])
}
