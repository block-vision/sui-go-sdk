package signer

import (
	"bytes"
	"encoding/base64"
	"fmt"

	"github.com/block-vision/sui-go-sdk/constant"
	"github.com/block-vision/sui-go-sdk/cryptography/scheme"
	"github.com/block-vision/sui-go-sdk/models"
	"github.com/block-vision/sui-go-sdk/mystenbcs"
	"golang.org/x/crypto/blake2b"
)

// Keypair signs Sui intent messages with one signature scheme. *Signer (Ed25519),
// *Secp256k1Signer and *Secp256r1Signer implement it.
type Keypair interface {
	// SignMessage signs intent || raw bytes of the base64 data. Transactions are signed
	// this way with constant.TransactionDataIntentScope.
	SignMessage(data string, scope constant.IntentScope) (*SignedMessageSerializedSig, error)
	// SignMessageV1 signs intent || BCS(bytes) of the base64 data. Personal messages are
	// signed this way with constant.PersonalMessageIntentScope.
	SignMessageV1(data string, scope constant.IntentScope) (*SignedMessageSerializedSig, error)
	// SignTransaction signs base64 TransactionData BCS bytes.
	SignTransaction(b64TxBytes string) (*models.SignedTransactionSerializedSig, error)
	// SignPersonalMessageV1 signs a personal message the way Sui wallets and verifiers expect.
	SignPersonalMessageV1(message string) (*SignedMessageSerializedSig, error)

	GetAddress() string
	PublicKeyBytes() []byte
	SignatureScheme() scheme.SignatureScheme
	// ExportSuiPrivateKey returns the private key as a Bech32 "suiprivkey…" string, the
	// format the Sui CLI and wallets import (like getSecretKey in the TypeScript SDK).
	ExportSuiPrivateKey() (string, error)
}

var (
	_ Keypair = (*Signer)(nil)
	_ Keypair = (*Secp256k1Signer)(nil)
	_ Keypair = (*Secp256r1Signer)(nil)
)

type SignedMessageSerializedSig struct {
	Message   string `json:"message"`
	Signature string `json:"signature"`
}

// ecdsaKeypair holds what the Secp256k1 and Secp256r1 signers share: both sign the
// blake2b intent digest with ECDSA over SHA-256 and serialize flag || r || s || pubkey.
type ecdsaKeypair struct {
	scheme  scheme.SignatureScheme
	pubKey  []byte // 33-byte compressed
	address string
	// signDigest returns the 64-byte r || s signature, s normalized to the lower half order.
	signDigest func(digest [32]byte) ([]byte, error)
}

func (k *ecdsaKeypair) GetAddress() string { return k.address }

func (k *ecdsaKeypair) PublicKeyBytes() []byte { return bytes.Clone(k.pubKey) }

func (k *ecdsaKeypair) SignatureScheme() scheme.SignatureScheme { return k.scheme }

func (k *ecdsaKeypair) SignMessage(data string, scope constant.IntentScope) (*SignedMessageSerializedSig, error) {
	payload, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return nil, fmt.Errorf("invalid base64 message: %w", err)
	}
	return k.sign(data, payload, scope)
}

func (k *ecdsaKeypair) SignMessageV1(data string, scope constant.IntentScope) (*SignedMessageSerializedSig, error) {
	raw, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return nil, fmt.Errorf("invalid base64 message: %w", err)
	}
	payload, err := mystenbcs.Marshal(raw)
	if err != nil {
		return nil, err
	}
	return k.sign(data, payload, scope)
}

func (k *ecdsaKeypair) SignTransaction(b64TxBytes string) (*models.SignedTransactionSerializedSig, error) {
	result, err := k.SignMessage(b64TxBytes, constant.TransactionDataIntentScope)
	if err != nil {
		return nil, err
	}
	return &models.SignedTransactionSerializedSig{TxBytes: result.Message, Signature: result.Signature}, nil
}

func (k *ecdsaKeypair) SignPersonalMessageV1(message string) (*SignedMessageSerializedSig, error) {
	return k.SignMessageV1(base64.StdEncoding.EncodeToString([]byte(message)), constant.PersonalMessageIntentScope)
}

func (k *ecdsaKeypair) sign(data string, payload []byte, scope constant.IntentScope) (*SignedMessageSerializedSig, error) {
	digest := blake2b.Sum256(models.NewMessageWithIntent(payload, scope))
	sig, err := k.signDigest(digest)
	if err != nil {
		return nil, err
	}
	flag := scheme.SignatureSchemeToFlag[k.scheme]
	return &SignedMessageSerializedSig{
		Message:   data,
		Signature: models.ToSerializedSignatureWithScheme(sig, k.pubKey, flag),
	}, nil
}

// suiAddress is blake2b256(flag || public key), hex encoded.
func suiAddress(flag byte, pubKey []byte) string {
	return models.PublicKeyToSuiAddress(pubKey, flag)
}
