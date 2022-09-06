package models

type SuiKeyPair struct {
	Flag    byte
	Address string

	PublicKey       []byte
	PublicKeyBase64 string

	PrivateKey []byte
}
