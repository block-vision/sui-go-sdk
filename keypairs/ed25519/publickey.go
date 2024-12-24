package ed25519

import (
	"bytes"
	"crypto/ed25519"
	"encoding/base64"
	"encoding/hex"
	"fmt"

	"github.com/machinebox/graphql"
	"golang.org/x/crypto/blake2b"

	"github.com/block-vision/sui-go-sdk/constant"
	"github.com/block-vision/sui-go-sdk/models"
	"github.com/block-vision/sui-go-sdk/mystenbcs"
)

type Ed25519PublicKey struct {
	signature []byte
}

func NewEd25519PublicKey(signature []byte) *Ed25519PublicKey {
	return &Ed25519PublicKey{
		signature: signature,
	}
}

func (e *Ed25519PublicKey) ToSuiAddress() string {
	panic("not implemented")
}

func (e *Ed25519PublicKey) VerifyPersonalMessage(message []byte, signature []byte, client *graphql.Client) (bool, error) {
	b64Message := base64.StdEncoding.EncodeToString([]byte(message))
	b64Signature := base64.StdEncoding.EncodeToString([]byte(signature))
	_, pass, err := VerifyMessage(b64Message, b64Signature, constant.PersonalMessageIntentScope)
	return pass, err
}

func VerifyMessage(message, signature string, scope constant.IntentScope) (signer string, pass bool, err error) {
	b64Bytes, _ := base64.StdEncoding.DecodeString(message)

	bcsEncodedMsg := bytes.Buffer{}
	bcsEncoder := mystenbcs.NewEncoder(&bcsEncodedMsg)
	bcsEncoder.Encode(b64Bytes)

	messageBytes := models.NewMessageWithIntent(bcsEncodedMsg.Bytes(), scope)

	serializedSignature, err := models.FromSerializedSignature(signature)
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
	newPubkey := []byte{byte(models.SigFlagEd25519)}
	newPubkey = append(newPubkey, pubKey...)

	addrBytes := blake2b.Sum256(newPubkey)
	return fmt.Sprintf("0x%s", hex.EncodeToString(addrBytes[:])[:64])
}
