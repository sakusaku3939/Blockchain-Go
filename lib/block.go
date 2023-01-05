package lib

import (
	"encoding/json"
	"fmt"
	"time"
)

type Block struct {
	Index     int
	PrevHash  string
	Data      string
	Timestamp time.Time
	Bits      int
}

func (b Block) ToJson() (string, error) {
	jsonData, err := json.Marshal(b)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return string(jsonData), nil
}
