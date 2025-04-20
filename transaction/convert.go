package transaction

import (
	"encoding/hex"

	"github.com/block-vision/sui-go-sdk/models"
	"github.com/block-vision/sui-go-sdk/utils"
	"github.com/mr-tron/base58"
)

func ConvertSuiAddressStringToBytes(address models.SuiAddress) (*models.SuiAddressBytes, error) {
	normalized := utils.NormalizeSuiAddress(string(address))
	decoded, err := hex.DecodeString(string(normalized[2:]))
	if err != nil {
		return nil, err
	}
	if len(decoded) != 32 {
		return nil, ErrInvalidSuiAddress
	}

	var fixedBytes [32]byte
	copy(fixedBytes[:], decoded)

	return (*models.SuiAddressBytes)(&fixedBytes), nil
}

func ConvertSuiAddressBytesToString(addr models.SuiAddressBytes) models.SuiAddress {
	return models.SuiAddress("0x" + hex.EncodeToString(addr[:]))
}

func ConvertObjectDigestStringToBytes(digest models.ObjectDigest) (*models.ObjectDigestBytes, error) {
	decoded, err := base58.Decode(string(digest))
	if err != nil {
		return nil, err
	}
	if len(decoded) != 32 {
		return nil, ErrInvalidObjectId
	}

	var fixedBytes [32]byte
	copy(fixedBytes[:], decoded)

	return (*models.ObjectDigestBytes)(&fixedBytes), nil
}

func ConvertObjectDigestBytesToString(digest models.ObjectDigestBytes) models.ObjectDigest {
	return models.ObjectDigest(base58.Encode(digest[:]))
}
