package zklogin

import (
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"strings"

	"github.com/block-vision/sui-go-sdk/cryptography/scheme"
	"github.com/block-vision/sui-go-sdk/mystenbcs"
	"golang.org/x/crypto/blake2b"
)

// zkLoginSignatureBCS is the on-chain BCS layout of a zkLogin signature (without the 0x05 flag).
type zkLoginSignatureBCS struct {
	Inputs struct {
		ProofPoints struct {
			A []string
			B [][]string
			C []string
		}
		IssBase64Details struct {
			Value     string
			IndexMod4 uint8
		}
		HeaderBase64 string
		AddressSeed  string // decimal Bn254 field element
	}
	MaxEpoch      uint64
	UserSignature []byte // serialized signature of the ephemeral key
}

// ParsedZkLoginSignature is a decoded zkLogin signature.
type ParsedZkLoginSignature struct {
	Iss          string   // OpenID provider, decoded from the JWT's iss claim
	AddressSeed  *big.Int // hash of the user's salt and JWT identity claims
	MaxEpoch     uint64
	HeaderBase64 string
	// UserSignature is the ephemeral key's serialized signature (flag || signature || public key).
	UserSignature []byte
}

// ParseZkLoginSignature decodes a base64 serialized zkLogin signature (flag 0x05). It only
// parses; verifying it needs on-chain state (JWKs, current epoch), use the v2 gRPC client's
// VerifyZkLoginSignature.
func ParseZkLoginSignature(serialized string) (*ParsedZkLoginSignature, error) {
	raw, err := mystenbcs.FromBase64(serialized)
	if err != nil {
		return nil, fmt.Errorf("invalid base64 signature: %w", err)
	}
	if len(raw) == 0 || raw[0] != scheme.SignatureSchemeToFlag[scheme.ZkLogin] {
		return nil, errors.New("not a zkLogin signature")
	}
	var sig zkLoginSignatureBCS
	n, err := mystenbcs.Unmarshal(raw[1:], &sig)
	if err != nil {
		return nil, fmt.Errorf("invalid zkLogin signature: %w", err)
	}
	if n != len(raw)-1 {
		return nil, fmt.Errorf("invalid zkLogin signature: %d trailing bytes", len(raw)-1-n)
	}
	iss, err := decodeStringClaim(sig.Inputs.IssBase64Details.Value, int(sig.Inputs.IssBase64Details.IndexMod4), "iss")
	if err != nil {
		return nil, fmt.Errorf("invalid zkLogin iss claim: %w", err)
	}
	seed, ok := new(big.Int).SetString(sig.Inputs.AddressSeed, 10)
	if !ok {
		return nil, fmt.Errorf("invalid zkLogin address seed: %q", sig.Inputs.AddressSeed)
	}
	return &ParsedZkLoginSignature{
		Iss:           iss,
		AddressSeed:   seed,
		MaxEpoch:      sig.MaxEpoch,
		HeaderBase64:  sig.Inputs.HeaderBase64,
		UserSignature: sig.UserSignature,
	}, nil
}

// Address returns the zkLogin address in the padded form; see ComputeZkLoginAddress.
func (p *ParsedZkLoginSignature) Address() (string, error) {
	return ComputeZkLoginAddress(p.Iss, p.AddressSeed)
}

// UnpaddedAddress returns the zkLogin address in the unpadded form; see ComputeZkLoginAddressUnpadded.
func (p *ParsedZkLoginSignature) UnpaddedAddress() (string, error) {
	return ComputeZkLoginAddressUnpadded(p.Iss, p.AddressSeed)
}

// ComputeZkLoginAddress derives the zkLogin address
// blake2b256(0x05 || len(iss) || iss || address seed as 32 big-endian bytes).
func ComputeZkLoginAddress(iss string, addressSeed *big.Int) (string, error) {
	seed, err := addressSeedBytes(addressSeed)
	if err != nil {
		return "", err
	}
	return zkLoginAddress(iss, seed)
}

// ComputeZkLoginAddressUnpadded derives the zkLogin address with the address seed's
// leading zero bytes removed. Sui accepts a zkLogin signature for either form; they
// only differ when the seed's top byte is zero.
func ComputeZkLoginAddressUnpadded(iss string, addressSeed *big.Int) (string, error) {
	seed, err := addressSeedBytes(addressSeed)
	if err != nil {
		return "", err
	}
	i := 0
	for i < len(seed)-1 && seed[i] == 0 {
		i++
	}
	return zkLoginAddress(iss, seed[i:])
}

func addressSeedBytes(addressSeed *big.Int) ([]byte, error) {
	if addressSeed == nil || addressSeed.Sign() < 0 || addressSeed.BitLen() > 256 {
		return nil, errors.New("zkLogin address seed must be a non-negative 256-bit integer")
	}
	return addressSeed.FillBytes(make([]byte, 32)), nil
}

func zkLoginAddress(iss string, seed []byte) (string, error) {
	if len(iss) == 0 || len(iss) > 255 {
		return "", fmt.Errorf("invalid zkLogin iss length: %d", len(iss))
	}
	h, _ := blake2b.New256(nil)
	h.Write([]byte{scheme.SignatureSchemeToFlag[scheme.ZkLogin], byte(len(iss))})
	h.Write([]byte(iss))
	h.Write(seed)
	return "0x" + hex.EncodeToString(h.Sum(nil)), nil
}

const base64URLAlphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-_"

// decodeStringClaim decodes a zkLogin claim: value is a base64url slice of the JWT payload
// starting at a position ≡ indexMod4 (mod 4), which decodes to `"name":"value",` (or a
// trailing `}`). The partial characters at both ends are trimmed bit-wise, as in Sui.
func decodeStringClaim(value string, indexMod4 int, name string) (string, error) {
	if len(value) < 2 {
		return "", errors.New("claim too short")
	}
	bits := make([]byte, 0, len(value)*6)
	for _, c := range value {
		idx := strings.IndexRune(base64URLAlphabet, c)
		if idx < 0 {
			return "", fmt.Errorf("invalid base64url character %q in claim", c)
		}
		for b := 5; b >= 0; b-- {
			bits = append(bits, byte(idx>>b)&1)
		}
	}
	switch indexMod4 % 4 {
	case 1:
		bits = bits[2:]
	case 2:
		bits = bits[4:]
	case 3:
		return "", errors.New("claim is not tightly packed (index mod 4 = 3)")
	}
	switch (indexMod4 + len(value) - 1) % 4 {
	case 1:
		bits = bits[:len(bits)-4]
	case 2:
		bits = bits[:len(bits)-2]
	case 0:
		return "", errors.New("claim is not tightly packed at its end")
	}
	if len(bits)%8 != 0 {
		return "", errors.New("claim bit length is not a multiple of 8")
	}
	out := make([]byte, len(bits)/8)
	for i, b := range bits {
		out[i/8] |= b << (7 - i%8)
	}
	claim := string(out)
	if !strings.HasSuffix(claim, ",") && !strings.HasSuffix(claim, "}") {
		return "", fmt.Errorf("invalid claim %q", claim)
	}
	var kv map[string]json.RawMessage
	if err := json.Unmarshal([]byte("{"+claim[:len(claim)-1]+"}"), &kv); err != nil {
		return "", fmt.Errorf("invalid claim %q: %w", claim, err)
	}
	raw, ok := kv[name]
	if !ok || len(kv) != 1 {
		return "", fmt.Errorf("claim %q is not %q", claim, name)
	}
	var v string
	if err := json.Unmarshal(raw, &v); err != nil {
		return "", fmt.Errorf("claim %q is not a string: %w", name, err)
	}
	return v, nil
}
