package lib

import (
	"encoding/json"
	"fmt"
	"time"
)

type Block struct {
	Index       int
	PrevHash    string
	Data        string
	Timestamp   time.Time
	Bits        string
	Nonce       int
	ElapsedTime string
	BlockHash   string
	blockHeader string
}

func (b Block) ToJson() (string, error) {
	jsonData, err := json.Marshal(b)

	if err != nil {
		fmt.Printf("Failed to convert to JSON: %s", err)
		return "", err
	}

	return string(jsonData), nil
}
