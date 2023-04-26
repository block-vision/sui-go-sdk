package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
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
