package lib

import (
	"Blockchain-Go/utils"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Block struct {
	Index       int
	PrevHash    string
	Data        string
	Timestamp   time.Time
	Bits        int
	Nonce       int
	ElapsedTime string
	BlockHash   string
	BlockHeader string
}

type jsonFields struct {
	Index       int
	PrevHash    string
	Data        string
	Timestamp   time.Time
	Bits        string
	Nonce       int
	ElapsedTime string
	BlockHash   string
}

func (b Block) ToJson() (string, error) {
	j := jsonFields{
		Index:       b.Index,
		PrevHash:    b.PrevHash,
		Data:        b.Data,
		Timestamp:   b.Timestamp,
		Bits:        strconv.FormatInt(int64(b.Bits), 16),
		Nonce:       b.Nonce,
		ElapsedTime: b.ElapsedTime,
		BlockHash:   b.BlockHash,
	}
	buf, err := json.Marshal(j)

	if err != nil {
		fmt.Println(err)
		return "", err
	}

	return string(buf), nil
}

func (b Block) CalcBlockHash() string {
	b.BlockHeader = strconv.Itoa(b.Index) + b.PrevHash + b.Data + b.Timestamp.String() + strconv.FormatInt(int64(b.Bits), 16)
	b.BlockHash = utils.SHA256(b.BlockHeader)
	return b.BlockHash
}

func (b Block) CalcTarget() string {
	exponentBytes := (b.Bits >> 24) - 3
	exponentBits := exponentBytes * 8
	coefficient := b.Bits & 0xffffff
	return b.LeftShift(int64(coefficient), exponentBits)
}

func (b Block) LeftShift(target int64, count int) string {
	hexTarget := strconv.FormatInt(target, 16)
	result := hexTarget + strings.Repeat("0", count/4)
	return result
}
