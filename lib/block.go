package lib

import (
	"Blockchain-Go/utils"
	"encoding/json"
	"fmt"
	"strconv"
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
	buf, err := json.Marshal(b)

	if err != nil {
		fmt.Println(err)
		return "", err
	}

	return string(buf), nil
}

func (b Block) CalcBlockHash() string {
	b.blockHeader = strconv.Itoa(b.Index) + b.PrevHash + b.Data + b.Timestamp.String() + b.Bits
	b.BlockHash = utils.SHA256(b.blockHeader)
	return b.BlockHash
}
