package zklogin

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

// Claim represents a claim with value and an index.
type Claim struct {
	Value     string
	IndexMod4 int
}

// Function to convert a single base64 URL character to a 6-bit array
func base64UrlCharTo6Bits(base64UrlChar rune) ([]int, error) {
	// Define the base64URL character set
	base64UrlCharacterSet := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-_"

	// Find the index of the input character in the base64URL character set
	index := strings.IndexRune(base64UrlCharacterSet, base64UrlChar)
	if index == -1 {
		return nil, fmt.Errorf("invalid base64Url character: %c", base64UrlChar)
	}

	// Convert the index to a 6-bit binary array
	bits := make([]int, 6)
	for i := 5; i >= 0; i-- {
		bits[i] = index % 2
		index /= 2
	}

	return bits, nil
}

// Function to convert a base64 URL string into a bit vector
func base64UrlStringToBitVector(base64UrlString string) ([]int, error) {
	bitVector := []int{}

	// Iterate through each character of the base64 URL string
	for _, base64UrlChar := range base64UrlString {
		// Convert each character to a 6-bit binary array
		bits, err := base64UrlCharTo6Bits(base64UrlChar)
		if err != nil {
			return nil, err
		}

		// Append the bits to the bit vector
		bitVector = append(bitVector, bits...)
	}

	return bitVector, nil
}

// decodeBase64URL function in Go
func decodeBase64URL(s string, i int) (string, error) {
	if len(s) < 2 {
		return "", fmt.Errorf("Input (s = %s) is not tightly packed because s.length < 2", s)
	}

	bits, err := base64UrlStringToBitVector(s)
	if err != nil {
		return "", err
	}

	firstCharOffset := i % 4
	if firstCharOffset == 1 {
		bits = bits[2:]
	} else if firstCharOffset == 2 {
		bits = bits[4:]
	} else if firstCharOffset == 3 {
		return "", fmt.Errorf("Input (s = %s) is not tightly packed because i%%4 = 3 (i = %d)", s, i)
	}

	lastCharOffset := (i + len(s) - 1) % 4
	if lastCharOffset == 1 {
		bits = bits[:len(bits)-4]
	} else if lastCharOffset == 2 {
		bits = bits[:len(bits)-2]
	} else if lastCharOffset == 0 {
		return "", fmt.Errorf("Input (s = %s) is not tightly packed because (i + s.length - 1)%%4 = 0 (i = %d)", s, i)
	}

	if len(bits)%8 != 0 {
		return "", errors.New("bit length is not a multiple of 8, invalid format")
	}

	// Convert the bit string into bytes
	byteArray := make([]byte, len(bits)/8)
	for i := 0; i < len(bits); i += 8 {
		bitChunk := bits[i : i+8]

		// Convert bitChunk to a byte
		var byteValue byte
		for j, bit := range bitChunk {
			if bit == '1' {
				byteValue |= 1 << (7 - j)
			}
		}
		byteArray[i/8] = byteValue
	}

	return string(byteArray), nil
}

// verifyExtendedClaim in Go
func verifyExtendedClaim(claim string) (string, interface{}, error) {
	if !(claim[len(claim)-1] == '}' || claim[len(claim)-1] == ',') {
		return "", nil, errors.New("Invalid claim")
	}

	// Parse claim as JSON, excluding the last character
	var jsonMap map[string]interface{}
	claimToParse := "{" + claim[:len(claim)-1] + "}"

	err := json.Unmarshal([]byte(claimToParse), &jsonMap)
	if err != nil {
		return "", nil, fmt.Errorf("failed to parse claim: %v", err)
	}

	if len(jsonMap) != 1 {
		return "", nil, errors.New("Invalid claim: more than one key found")
	}

	for key, value := range jsonMap {
		return key, value, nil
	}

	return "", nil, errors.New("Claim parsing error")
}

// ExtractClaimValue decodes and verifies the claim.
func extractClaimValue(claim Claim, claimName string) (string, error) {
	extendedClaim, err := decodeBase64URL(claim.Value, claim.IndexMod4)
	if err != nil {
		return "", err
	}

	name, value, err := verifyExtendedClaim(extendedClaim)
	if err != nil {
		return "", err
	}

	if name != claimName {
		return "", errors.New("invalid field name: found " + name + " expected " + claimName)
	}

	// Assuming the value is actually a string
	if strValue, ok := value.(string); ok {
		return strValue, nil
	} else {
		// Handle the case where value was not a string
		// For example, you might want to return an empty string or an error
		return "", errors.New("value is not a string")
	}
}
