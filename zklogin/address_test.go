package zklogin

import (
	"bufio"
	"encoding/base64"
	"math/big"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

// testdata holds zkLogin signatures of mainnet transactions as "digest sender signature".
func TestParseMainnetZkLoginSignatures(t *testing.T) {
	f, err := os.Open("testdata/mainnet_zklogin_signatures.txt")
	require.NoError(t, err)
	defer f.Close()
	sc := bufio.NewScanner(f)
	sc.Buffer(make([]byte, 1<<20), 1<<20)

	wantIss := map[string]string{
		"22DVTKoKreTkHj2oqUaeHhzrSUCQMU5NzKFDrjGKbG77": "https://accounts.google.com",
		"4BiaKm1fASH96gENw1W93VS8ioUGWZdsHp28juAUiMkD": "https://login.onepassport.onefc.com/de3ee5c1-5644-4113-922d-e8336569a462/v2.0/",
	}
	n := 0
	for sc.Scan() {
		parts := strings.Fields(sc.Text())
		digest, sender, sig := parts[0], parts[1], parts[2]
		t.Run(digest, func(t *testing.T) {
			p, err := ParseZkLoginSignature(sig)
			require.NoError(t, err)
			require.Equal(t, wantIss[digest], p.Iss)
			require.NotZero(t, p.MaxEpoch)
			require.NotEmpty(t, p.UserSignature)

			addr, err := p.Address()
			require.NoError(t, err)
			require.Equal(t, sender, addr, "zkLogin address must be the transaction sender")
		})
		n++
	}
	require.Equal(t, 2, n)
}

func TestComputeZkLoginAddressUnpadded(t *testing.T) {
	iss := "https://accounts.google.com"
	small := big.NewInt(42) // top bytes zero: the two forms differ
	padded, err := ComputeZkLoginAddress(iss, small)
	require.NoError(t, err)
	unpadded, err := ComputeZkLoginAddressUnpadded(iss, small)
	require.NoError(t, err)
	require.NotEqual(t, padded, unpadded)

	full := new(big.Int).Lsh(big.NewInt(1), 255) // top byte non-zero: identical
	padded, err = ComputeZkLoginAddress(iss, full)
	require.NoError(t, err)
	unpadded, err = ComputeZkLoginAddressUnpadded(iss, full)
	require.NoError(t, err)
	require.Equal(t, padded, unpadded)

	_, err = ComputeZkLoginAddress(iss, new(big.Int).Lsh(big.NewInt(1), 256))
	require.Error(t, err)
	_, err = ComputeZkLoginAddress(iss, big.NewInt(-1))
	require.Error(t, err)
	_, err = ComputeZkLoginAddress("", small)
	require.Error(t, err)
}

func TestParseZkLoginSignatureRejectsOtherInput(t *testing.T) {
	ed25519Sig := base64.StdEncoding.EncodeToString(append([]byte{0x00}, make([]byte, 96)...))
	_, err := ParseZkLoginSignature(ed25519Sig)
	require.Error(t, err)
	_, err = ParseZkLoginSignature("not base64!")
	require.Error(t, err)
	_, err = ParseZkLoginSignature(base64.StdEncoding.EncodeToString([]byte{0x05, 0x01}))
	require.Error(t, err)
}
