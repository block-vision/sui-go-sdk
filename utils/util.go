package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"
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
