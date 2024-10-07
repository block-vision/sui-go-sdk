package mystenbcs

import "encoding/base64"

func FromBase64(base64String string) ([]byte, error) {
	bytes, err := base64.StdEncoding.DecodeString(base64String)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}

func ToBase64(bytes []byte) string {
	return base64.StdEncoding.EncodeToString(bytes)
}
