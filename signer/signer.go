package signer

import (
	"crypto/ed25519"
	"encoding/hex"
	"github.com/block-vision/sui-go-sdk/common/keypair"
	"golang.org/x/crypto/sha3"
)

const (
	SigntureFlagEd25519     = 0x0
	SigntureFlagSecp256k1   = 0x1
	AddressLength           = 64
	DerivationPathEd25519   = `m/44'/784'/0'/0'/0'`
	DerivationPathSecp256k1 = `m/54'/784'/0'/0/0`
)

type Account struct {
	priKey  ed25519.PrivateKey
	pubKey  ed25519.PublicKey
	Address string
}

func NewSigner(seed []byte) *Account {
	priKey := ed25519.NewKeyFromSeed(seed[:])
	pubKey := priKey.Public().(ed25519.PublicKey)

	tmp := []byte{byte(keypair.Ed25519Flag)}
	tmp = append(tmp, pubKey...)
	addrBytes := sha3.Sum256(tmp)
	addr := "0x" + hex.EncodeToString(addrBytes[:])[:AddressLength]

	return &Account{
		priKey:  priKey,
		pubKey:  pubKey,
		Address: addr,
	}
}

func NewSignerWithPriKey(priKey ed25519.PrivateKey) *Account {
	pubKey := priKey.Public().(ed25519.PublicKey)

	tmp := []byte{byte(keypair.Ed25519Flag)}
	tmp = append(tmp, pubKey...)
	addrBytes := sha3.Sum256(tmp)
	addr := "0x" + hex.EncodeToString(addrBytes[:])[:AddressLength]

	return &Account{
		priKey:  priKey,
		pubKey:  pubKey,
		Address: addr,
	}
}

