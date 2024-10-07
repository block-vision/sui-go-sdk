package zklogin

import (
	"encoding/hex"
	"math/big"
)

// ToPaddedBigEndianBytes converts a big.Int to a big-endian byte array of a specified width.
func ToPaddedBigEndianBytes(num *big.Int, width int) []byte {
	hexStr := num.Text(16)                        // Convert big.Int to a hexadecimal string
	paddedHexStr := padLeft(hexStr, width*2, '0') // Pad to desired width

	decodedBytes, _ := hex.DecodeString(paddedHexStr)
	return decodedBytes[len(decodedBytes)-width:]
}

// padLeft pads the input string on the left with the specified padding character to the desired length.
func padLeft(str string, length int, padChar rune) string {
	for len(str) < length {
		str = string(padChar) + str
	}
	return str
}
