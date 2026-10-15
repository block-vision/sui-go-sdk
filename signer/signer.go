package signer

import (
	"bytes"
	"crypto"
	"crypto/ed25519"
	"encoding/base64"
	"encoding/hex"
	"fmt"

	"github.com/block-vision/sui-go-sdk/common/keypair"
	"github.com/block-vision/sui-go-sdk/constant"
	"github.com/block-vision/sui-go-sdk/cryptography/scheme"
	"github.com/block-vision/sui-go-sdk/models"
	"github.com/block-vision/sui-go-sdk/mystenbcs"
	"github.com/cosmos/go-bip39"
	"golang.org/x/crypto/blake2b"
)

const (
	SigntureFlagEd25519     = 0x0
	SigntureFlagSecp256k1   = 0x1
	AddressLength           = 64
	DerivationPathEd25519   = `m/44'/784'/0'/0'/0'`
	DerivationPathSecp256k1 = `m/54'/784'/0'/0/0`
)

type Signer struct {
	PriKey  ed25519.PrivateKey
	PubKey  ed25519.PublicKey
	Address string
}

func NewSigner(seed []byte) *Signer {
	priKey := ed25519.NewKeyFromSeed(seed[:])
	pubKey := priKey.Public().(ed25519.PublicKey)

	tmp := []byte{byte(keypair.Ed25519Flag)}
	tmp = append(tmp, pubKey...)
	addrBytes := blake2b.Sum256(tmp)
	addr := "0x" + hex.EncodeToString(addrBytes[:])[:AddressLength]

	return &Signer{
		PriKey:  priKey,
		PubKey:  pubKey,
		Address: addr,
	}
}

// NewSignerWithSecretKey creates an Ed25519 signer from a Bech32 "suiprivkey…" string.
// It rejects Secp256k1/Secp256r1 keys instead of misreading them as Ed25519 seeds; use
// NewKeypairFromSuiPrivateKey to accept any supported scheme.
func NewSignerWithSecretKey(secret string) (*Signer, error) {
	keyScheme, privKey, err := DecodeSuiPrivateKey(secret)
	if err != nil {
		return nil, err
	}
	if keyScheme != scheme.ED25519 {
		return nil, fmt.Errorf("%s private key is not an Ed25519 key, use NewKeypairFromSuiPrivateKey", keyScheme)
	}
	return NewSigner(privKey), nil
}

func NewSignertWithMnemonic(mnemonic string) (*Signer, error) {
	seed, err := bip39.NewSeedWithErrorChecking(mnemonic, "")
	if err != nil {
		return nil, err
	}
	key, err := DeriveForPath("m/44'/784'/0'/0'/0'", seed)
	if err != nil {
		return nil, err
	}
	return NewSigner(key.Key), nil
}

func (s *Signer) GetAddress() string { return s.Address }

func (s *Signer) PublicKeyBytes() []byte { return bytes.Clone(s.PubKey) }

func (s *Signer) SignatureScheme() scheme.SignatureScheme { return scheme.ED25519 }

// ExportSuiPrivateKey returns the key as a Bech32 "suiprivkey…" string.
func (s *Signer) ExportSuiPrivateKey() (string, error) {
	return EncodeSuiPrivateKey(scheme.ED25519, s.PriKey.Seed())
}

func (s *Signer) SignMessage(data string, scope constant.IntentScope) (*SignedMessageSerializedSig, error) {
	txBytes, _ := base64.StdEncoding.DecodeString(data)
	message := models.NewMessageWithIntent(txBytes, scope)
	digest := blake2b.Sum256(message)
	var noHash crypto.Hash
	sigBytes, err := s.PriKey.Sign(nil, digest[:], noHash)
	if err != nil {
		return nil, err
	}

	ret := &SignedMessageSerializedSig{
		Message:   data,
		Signature: models.ToSerializedSignature(sigBytes, s.PriKey.Public().(ed25519.PublicKey)),
	}
	return ret, nil
}

// SignTransaction signs base64 TransactionData BCS bytes with the TransactionData intent.
func (s *Signer) SignTransaction(b64TxBytes string) (*models.SignedTransactionSerializedSig, error) {
	result, err := s.SignMessage(b64TxBytes, constant.TransactionDataIntentScope)
	if err != nil {
		return nil, err
	}

	return &models.SignedTransactionSerializedSig{
		TxBytes:   result.Message,
		Signature: result.Signature,
	}, nil
}

// Deprecated: the signature is over the raw message rather than its BCS encoding, so Sui
// nodes and wallets do not accept it as a personal message signature. Use SignPersonalMessageV1.
func (s *Signer) SignPersonalMessage(message string) (*SignedMessageSerializedSig, error) {
	b64Message := base64.StdEncoding.EncodeToString([]byte(message))
	return s.SignMessage(b64Message, constant.PersonalMessageIntentScope)
}

// SignPersonalMessageV1 is the same as SignPersonalMessage, but it uses the new message format for personal messages.
func (s *Signer) SignPersonalMessageV1(message string) (*SignedMessageSerializedSig, error) {
	b64Message := base64.StdEncoding.EncodeToString([]byte(message))
	return s.SignMessageV1(b64Message, constant.PersonalMessageIntentScope)
}

// SignMessageV1 is the same as SignMessage, but it uses the new message format for personal messages.
func (s *Signer) SignMessageV1(data string, scope constant.IntentScope) (*SignedMessageSerializedSig, error) {
	b64Bytes, _ := base64.StdEncoding.DecodeString(data)

	bcsEncodedMsg := bytes.Buffer{}
	bcsEncoder := mystenbcs.NewEncoder(&bcsEncodedMsg)
	bcsEncoder.Encode(b64Bytes)
	message := models.NewMessageWithIntent(bcsEncodedMsg.Bytes(), scope)
	digest := blake2b.Sum256(message)
	var noHash crypto.Hash
	sigBytes, err := s.PriKey.Sign(nil, digest[:], noHash)
	if err != nil {
		return nil, err
	}

	ret := &SignedMessageSerializedSig{
		Message:   data,
		Signature: models.ToSerializedSignature(sigBytes, s.PriKey.Public().(ed25519.PublicKey)),
	}
	return ret, nil
}
