package utils

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/mr-tron/base58"
	"golang.org/x/crypto/blake2b"
)

func PrettyPrint(v interface{}) {
	b, err := json.Marshal(v)
	if err != nil {
		fmt.Println(v)
		return
	}

	var out bytes.Buffer
	err = json.Indent(&out, b, "", "  ")
	if err != nil {
		fmt.Println(v)
		return
	}

	fmt.Println(out.String())
}

func IsFieldNonEmpty(v interface{}, fieldName string) bool {
	rv := reflect.ValueOf(v)

	field := rv.FieldByName(fieldName)

	if !field.IsValid() {
		return false
	}

	return !field.IsZero()
}

// GetTxDigest get transaction digest from tx bytes in base64
// go version of https://github.com/MystenLabs/sui/blob/main/sdk/typescript/src/transactions/TransactionData.ts
func GetTxDigest(txBytesB64 string) (string, error) {

	txBytes, err := base64.StdEncoding.DecodeString(txBytesB64)
	if err != nil {
		return "", err
	}

	return GetTxDigestFromBytes(txBytes)
}

// GetTxDigestFromBytes get transaction digest from tx bytes
func GetTxDigestFromBytes(txBytes []byte) (string, error) {
	typedData, err := hashTypedData("TransactionData", txBytes)
	if err != nil {
		return "", err
	}

	return base58.Encode(typedData), nil
}

func hashTypedData(typeTag string, data []byte) ([]byte, error) {
	// Convert typeTag to bytes and append "::"
	typeTagBytes := []byte(typeTag + "::")

	// Create a new byte array to hold typeTagBytes and data
	dataWithTag := append(typeTagBytes, data...)

	// Perform BLAKE2b hashing with a digest size of 32 bytes
	hash, err := blake2b.New(32, nil)
	if err != nil {
		return nil, err
	}

	_, err = hash.Write(dataWithTag)
	if err != nil {
		return nil, err
	}

	return hash.Sum(nil), nil
}
