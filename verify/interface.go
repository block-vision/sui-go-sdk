package verify

import "github.com/machinebox/graphql"

// Deprecated: use models.PublicKeyToSuiAddress to derive an address and
// models.VerifyPersonalMessage to verify a signature.
type IPublicKey interface {
	ToSuiAddress() string
	VerifyPersonalMessage(message []byte, signature []byte, client *graphql.Client) (bool, error)
}
