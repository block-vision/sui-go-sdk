package scheme

type SignatureScheme string

const (
	ED25519   SignatureScheme = "ED25519"
	Secp256k1 SignatureScheme = "Secp256k1"
	Secp256r1 SignatureScheme = "Secp256r1"
	MultiSig  SignatureScheme = "MultiSig"
	ZkLogin   SignatureScheme = "ZkLogin"
)

var SignatureSchemeToSize = map[SignatureScheme]int{
	ED25519:   32,
	Secp256k1: 33,
	Secp256r1: 33,
}

var SignatureSchemeToFlag = map[SignatureScheme]byte{
	ED25519:   0x00,
	Secp256k1: 0x01,
	Secp256r1: 0x02,
	MultiSig:  0x03,
	ZkLogin:   0x05,
}

var SignatureFlagToScheme = map[byte]SignatureScheme{
	0x00: ED25519,
	0x01: Secp256k1,
	0x02: Secp256r1,
	0x03: MultiSig,
	0x05: ZkLogin,
}
