package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func IndentJSON(s string) (string, error) {
	var buf bytes.Buffer
	err := json.Indent(&buf, []byte(s), "", "  ")

	if err != nil {
		fmt.Println(err)
		return "", err
	}

	return buf.String(), nil
}

func CompactJSON(s string) (string, error) {
	var buf bytes.Buffer
	err := json.Compact(&buf, []byte(s))

	if err != nil {
		fmt.Println(err)
		return "", err
	}

	return buf.String(), nil
}
