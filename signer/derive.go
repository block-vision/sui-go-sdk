package signer

import (
	"bytes"
	"crypto/ed25519"
	"crypto/hmac"
	"crypto/sha512"
	"encoding/binary"
	"errors"
	"regexp"
	"strconv"
	"strings"
)

const (
	FirstHardenedIndex = uint32(0x80000000)
	seedModifier       = "ed25519 seed"
)

var (
	ErrInvalidPath        = errors.New("invalid derivation path")
	ErrNoPublicDerivation = errors.New("no public derivation for ed25519")

	pathRegex = regexp.MustCompile(`^m(\/[0-9]+')+$`)
)

type Key struct {
	Key       []byte
	ChainCode []byte
}

// DeriveForPath derives key for a path in BIP-44 format and a seed.
// Ed25119 derivation operated on hardened keys only.
func DeriveForPath(path string, seed []byte) (*Key, error) {
	if !isValidPath(path) {
		return nil, ErrInvalidPath
	}

	key, err := NewMasterKey(seed)
	if err != nil {
		return nil, err
	}

	segments := strings.Split(path, "/")
	for _, segment := range segments[1:] {
		i64, err := strconv.ParseUint(strings.TrimRight(segment, "'"), 10, 32)
		if err != nil {
			return nil, err
		}

		i := uint32(i64) + FirstHardenedIndex
		key, err = key.Derive(i)
		if err != nil {
			return nil, err
		}
	}

	return key, nil
}

// NewMasterKey generates a new master key from seed.
func NewMasterKey(seed []byte) (*Key, error) {
	hash := hmac.New(sha512.New, []byte(seedModifier))
	_, err := hash.Write(seed)
	if err != nil {
		return nil, err
	}
	sum := hash.Sum(nil)
	key := &Key{
		Key:       sum[:32],
		ChainCode: sum[32:],
	}
	return key, nil
}

func (k *Key) Derive(i uint32) (*Key, error) {
	// no public derivation for ed25519
	if i < FirstHardenedIndex {
		return nil, ErrNoPublicDerivation
	}

	iBytes := make([]byte, 4)
	binary.BigEndian.PutUint32(iBytes, i)
	key := append([]byte{0x0}, k.Key...)
	data := append(key, iBytes...)

	hash := hmac.New(sha512.New, k.ChainCode)
	_, err := hash.Write(data)
	if err != nil {
		return nil, err
	}
	sum := hash.Sum(nil)
	newKey := &Key{
		Key:       sum[:32],
		ChainCode: sum[32:],
	}
	return newKey, nil
}

// PublicKey returns public key for a derived private key.
func (k *Key) PublicKey() ([]byte, error) {
	reader := bytes.NewReader(k.Key)
	pub, _, err := ed25519.GenerateKey(reader)
	if err != nil {
		return nil, err
	}
	return pub[:], nil
}

// RawSeed returns raw seed bytes
func (k *Key) RawSeed() [32]byte {
	var rawSeed [32]byte
	copy(rawSeed[:], k.Key[:])
	return rawSeed
}

func isValidPath(path string) bool {
	if !pathRegex.MatchString(path) {
		return false
	}

	// Check for overflows
	segments := strings.Split(path, "/")
	for _, segment := range segments[1:] {
		_, err := strconv.ParseUint(strings.TrimRight(segment, "'"), 10, 32)
		if err != nil {
			return false
		}
	}

	return true
}
