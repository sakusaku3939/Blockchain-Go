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

func (b *Block) ToJson() (string, error) {
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

func (b *Block) CalcBlockHash() string {
	b.BlockHeader = strconv.Itoa(b.Index) + b.PrevHash + b.Data + b.Timestamp.String() + strconv.FormatInt(int64(b.Bits), 16)
	b.BlockHash = utils.SHA256(b.BlockHeader)
	fmt.Println("hash: ", b.BlockHash)

	return b.BlockHash
}

func (b *Block) CalcTarget() int64 {
	exponentBytes := (b.Bits >> 24) - 4
	exponentBits := exponentBytes * 8
	coefficient := b.Bits & 0xffffff
	return int64(coefficient << exponentBits)
}

func (b *Block) CheckValidHash() bool {
	blockHash, _ := strconv.ParseInt(b.CalcBlockHash(), 16, 64)
	return blockHash <= b.CalcTarget()
}
