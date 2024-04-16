package keypair

import (
	"encoding/base64"
	"encoding/hex"

	"golang.org/x/crypto/sha3"
)

func fromPublicKeyBytesToAddress(publicKey []byte, scheme byte) string {
	if scheme != byte(Ed25519Flag) && scheme != byte(Secp256k1Flag) {
		return ""
	}
	tmp := make([]byte, len(publicKey)+1)
	tmp[0] = scheme
	for i := range publicKey {
		tmp[i+1] = publicKey[i]
	}
	hexHash := sha3.Sum256(tmp)
	return "0x" + hexEncode(hexHash[:])[:AccountAddress32Length*2]
}

func hexEncode(b []byte) string {
	enc := make([]byte, len(b)*2)
	hex.Encode(enc, b)
	return string(enc)
}

func encodeBase64(value []byte) string {
	return base64.StdEncoding.EncodeToString(value)
}
