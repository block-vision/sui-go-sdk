package models

import (
	"reflect"

	"github.com/block-vision/sui-go-sdk/mystenbcs"
)

type SuiAddress string
type SuiAddressBytes [32]byte
type TransactionDigest string
type ObjectDigest string
type ObjectDigestBytes []byte

func init() {
	var suiAddressBytes SuiAddressBytes
	if reflect.ValueOf(suiAddressBytes).Type().Name() != mystenbcs.SuiAddressBytesName {
		panic("SuiAddressBytes type name not match")
	}
}

func (s SuiAddressBytes) IsEqual(other SuiAddressBytes) bool {
	for i, b := range s {
		if b != other[i] {
			return false
		}
	}
	return true
}

func (o ObjectDigestBytes) IsEqual(other ObjectDigestBytes) bool {
	if len(o) != len(other) {
		return false
	}

	for i, b := range o {
		if b != other[i] {
			return false
		}
	}
	return true
}
