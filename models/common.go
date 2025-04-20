package models

type SuiAddress string
type SuiAddressBytes [32]byte
type TransactionDigest string
type ObjectDigest string
type ObjectDigestBytes [32]byte

func (s *SuiAddressBytes) IsZero() bool {
	for _, b := range s {
		if b != 0 {
			return false
		}
	}
	return true
}

func (s SuiAddressBytes) IsEqual(other SuiAddressBytes) bool {
	for i, b := range s {
		if b != other[i] {
			return false
		}
	}
	return true
}

func (o *ObjectDigestBytes) IsZero() bool {
	for _, b := range o {
		if b != 0 {
			return false
		}
	}
	return true
}

func (o ObjectDigestBytes) IsEqual(other ObjectDigestBytes) bool {
	for i, b := range o {
		if b != other[i] {
			return false
		}
	}
	return true
}
