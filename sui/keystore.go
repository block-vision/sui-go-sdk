package sui

import (
	"crypto/sha256"
	"strings"

	"github.com/block-vision/sui-go-sdk/common/sui_error"
	"github.com/ledgerwatch/secp256k1"

	"github.com/block-vision/sui-go-sdk/common/keypair"
	"github.com/block-vision/sui-go-sdk/config"
	"github.com/block-vision/sui-go-sdk/models"
	"golang.org/x/crypto/ed25519"
)

type IAccountKeyStore interface {
	Sign(address string, msg []byte) (signature []byte, err error)
	GetKey(address string) (models.SuiKeyPair, error)
	AddKey(keypair string) (err error)
	Keys() (publicKeys []string)
}

type AccountKeyStoreImpl struct {
	defaultAddress string
	keystore       map[string]models.SuiKeyPair
}

var accountStore *AccountKeyStoreImpl

func SetAccountKeyStore(configDir string) (IAccountKeyStore, error) {
	kpStrs, err := config.GetKeyStore(configDir)
	if err != nil {
		return accountStore, err
	}
	accountStore = &AccountKeyStoreImpl{
		keystore: make(map[string]models.SuiKeyPair),
	}
	for i := range kpStrs.Keys {
		if i == 0 {
			accountStore.defaultAddress = kpStrs.Keys[i]
		}
		kp, subErr := keypair.FetchKeyPair(kpStrs.Keys[i])
		if subErr != nil {
			return accountStore, subErr
		}
		accountStore.keystore[strings.ToLower(kp.Address)] = kp
	}
	return accountStore, nil
}

func (a *AccountKeyStoreImpl) Sign(address string, msg []byte) (signature []byte, err error) {
	if a.keystore == nil {
		return []byte{}, sui_error.ErrNoKeyStoreInfo
	}

	kp, ok := a.keystore[strings.ToLower(address)]
	if !ok {
		return []byte{}, sui_error.ErrAddressNotInKeyStore
	}
	if kp.Flag == keypair.Ed25519Flag {
		return ed25519.Sign(ed25519.NewKeyFromSeed(kp.PrivateKey), msg), nil
	}
	data := sha256.Sum256(msg)
	return secp256k1.Sign(data[:], kp.PrivateKey)
}

func (a *AccountKeyStoreImpl) GetKey(address string) (models.SuiKeyPair, error) {
	if a.keystore == nil {
		return models.SuiKeyPair{}, sui_error.ErrNoKeyStoreInfo
	}
	kp, ok := a.keystore[strings.ToLower(address)]
	if !ok {
		return models.SuiKeyPair{}, sui_error.ErrInvalidAddress
	}
	return kp, nil
}

func (a *AccountKeyStoreImpl) AddKey(key string) (err error) {
	if a.keystore == nil {
		return sui_error.ErrNoKeyStoreInfo
	}
	kp, err := keypair.FetchKeyPair(key)
	if err != nil {
		return err
	}
	a.keystore[strings.ToLower(kp.Address)] = kp
	return nil
}

func (a *AccountKeyStoreImpl) Keys() (publicKeys []string) {
	if a.keystore == nil {
		return []string{}
	}
	for _, kp := range a.keystore {
		publicKeys = append(publicKeys, kp.PublicKeyBase64)

	}
	return publicKeys
}
