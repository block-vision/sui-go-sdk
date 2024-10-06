package zklogin

import (
	"context"
	"errors"
	"fmt"
	"math/big"

	"github.com/machinebox/graphql"

	"github.com/block-vision/sui-go-sdk/cryptography/scheme"
	"github.com/block-vision/sui-go-sdk/mystenbcs"
)

type SignaturePubkeyPair struct {
	SerializedSignature string
	SignatureScheme     scheme.SignatureScheme
	ZkLogin             *ZkLoginSignature
	Signature           []byte
	PubKey              []byte
}

type ZkLoginPublicIdentifierOptions struct {
	Client *graphql.Client
}

type ZkLoginPublicIdentifier struct {
	data    []byte
	options *ZkLoginPublicIdentifierOptions
}

func NewZkLoginPublicIdentifier(data []byte, options *ZkLoginPublicIdentifierOptions) *ZkLoginPublicIdentifier {
	return &ZkLoginPublicIdentifier{
		data:    data,
		options: options,
	}
}

/**
 * Return the byte array representation of the zkLogin public identifier
 */
func (p *ZkLoginPublicIdentifier) toRawBytes() []byte {
	return p.data
}

func (p *ZkLoginPublicIdentifier) ToSuiAddress() string {

	// Each hex char represents half a byte, hence hex address doubles the length
	// return normalizeSuiAddress(
	// 	bytesToHex(blake2b(this.toSuiBytes(), { dkLen: 32 })).slice(0, SUI_ADDRESS_LENGTH * 2),
	// );

	// Convert the public identifier to a Sui address
	return "0x" + mystenbcs.ToHex(mystenbcs.Blake2b(p.toSuiBytes(), 32))[:40]
}

func (pk *ZkLoginPublicIdentifier) VerifyPersonalMessage(message []byte, signature []byte, client *graphql.Client) (bool, error) {
	// Parse the serialized zkLogin signature
	parsedSignature, err := ParseSerializedZkLoginSignature(signature)
	if err != nil {
		return false, fmt.Errorf("failed to parse serialized zkLogin signature: %w", err)
	}

	// convert the public key to a Sui address
	address := pk.ToSuiAddress()

	// Convert the message to Base64
	bytesEncoded := mystenbcs.ToBase64(message)

	// Call the GraphQL verification function
	return GraphqlVerifyZkLoginSignature(address, bytesEncoded, string(parsedSignature.SerializedSignature), "PERSONAL_MESSAGE", client)
}

func toZkLoginPublicIdentifier(addressSeed *big.Int, iss string, options *ZkLoginPublicIdentifierOptions) *ZkLoginPublicIdentifier {
	addressSeedBytesBigEndian := ToPaddedBigEndianBytes(addressSeed, 32)

	issBytes := []byte(iss)
	tmp := make([]byte, 1+len(issBytes)+32)

	tmp[0] = byte(len(issBytes))
	copy(tmp[1:], issBytes)
	copy(tmp[1+len(issBytes):], addressSeedBytesBigEndian)

	return NewZkLoginPublicIdentifier(tmp, options)
}

type VerifyZkloginSignatureResponse struct {
	VerifyZkloginSignature struct {
		Success bool     `json:"success"`
		Errors  []string `json:"errors"`
	} `json:"verifyZkloginSignature"`
}

func GraphqlVerifyZkLoginSignature(address string, bytes string, signature string, intentScope string, client *graphql.Client) (bool, error) {
	// Define the GraphQL request
	req := graphql.NewRequest(`
		query Zklogin($bytes: Base64!, $signature: Base64!, $intentScope: ZkLoginIntentScope!, $author: SuiAddress!) {
			verifyZkloginSignature(
				bytes: $bytes,
				signature: $signature,
				intentScope: $intentScope,
				author: $author
			) {
				success
				errors
			}
		}
	`)

	// Set the request variables
	req.Var("bytes", bytes)
	req.Var("signature", signature)
	req.Var("intentScope", intentScope)
	req.Var("author", address)

	// Execute the request
	ctx := context.Background()
	var respData VerifyZkloginSignatureResponse
	if err := client.Run(ctx, req, &respData); err != nil {
		return false, fmt.Errorf("failed to execute graphql query: %w", err)
	}

	// Evaluate the response
	success := respData.VerifyZkloginSignature.Success
	errors := respData.VerifyZkloginSignature.Errors
	return success && len(errors) == 0, nil
}

// Function to parse the serialized ZkLoginSignature
func ParseSerializedZkLoginSignature(signature interface{}) (*SignaturePubkeyPair, error) {
	var bytes []byte
	var err error

	// Check if the input is a base64 string or byte array
	switch sig := signature.(type) {
	case string:
		bytes, err = mystenbcs.FromBase64(sig)
		if err != nil {
			return nil, fmt.Errorf("failed to decode base64: %v", err)
		}
	case []byte:
		bytes = sig
	default:
		return nil, errors.New("unsupported input type")
	}

	// Check if the signature scheme is correct
	if bytes[0] != scheme.SignatureSchemeToFlag[scheme.ZkLogin] {
		return nil, errors.New("invalid signature scheme")
	}

	// Parse the signature bytes
	signatureBytes := bytes[1:]
	zkSig, err := parseZkLoginSignature(signatureBytes) // Assume parseZkLoginSignature is defined
	if err != nil {
		return nil, fmt.Errorf("failed to parse ZkLoginSignature: %v", err)
	}

	// Extract necessary fields from the parsed signature
	inputs := zkSig.Inputs
	issBase64Details := inputs.IssBase64Details
	addressSeed := inputs.AddressSeed

	zkSig.AddressSeed = addressSeed

	// Extract the claim value
	iss, err := extractClaimValue(Claim{
		Value:     issBase64Details.Value,
		IndexMod4: int(issBase64Details.IndexMod4),
	}, "iss")
	if err != nil {
		return nil, fmt.Errorf("failed to extract claim value: %v", err)
	}

	zkSig.Iss = iss

	// Calculate the public identifier (you need to implement toZkLoginPublicIdentifier)
	addressSeedBigInt, _ := new(big.Int).SetString(addressSeed, 10)
	publicIdentifier := toZkLoginPublicIdentifier(addressSeedBigInt, iss, nil)

	// Return the parsed signature data
	return &SignaturePubkeyPair{
		SerializedSignature: mystenbcs.ToBase64(bytes),
		SignatureScheme:     scheme.ZkLogin,
		ZkLogin:             zkSig,
		Signature:           signatureBytes,
		PubKey:              publicIdentifier.toRawBytes(),
	}, nil
}
