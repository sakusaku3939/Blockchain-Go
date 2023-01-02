package lib

import (
	"encoding/json"
	"fmt"
	"time"
)

type Block struct {
	index     int
	prevHash  string
	data      string
	timestamp time.Time
	bits      int
}

func NewBlock(index int, prevHash string, data string, timestamp time.Time, bits int) *Block {
	b := new(Block)
	b.index = index
	b.prevHash = prevHash
	b.data = data
	b.timestamp = timestamp
	b.bits = bits
	return b
}

func (b Block) ToJson() ([]byte, error) {
	jsonData, err := json.Marshal(b)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return jsonData, nil
}
