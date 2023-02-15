package lib

import (
	"Blockchain-Go/utils"
	"encoding/json"
	"fmt"
	"math/big"
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

func (b *Block) ToJson() string {
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
	}

	return string(buf)
}

func (b *Block) CheckValidHash() bool {
	blockHash, _ := new(big.Int).SetString(b.calcBlockHash(), 16)
	return blockHash.Cmp(b.calcTarget()) != 1 // blockHash <= b.calcTarget()
}

func (b *Block) calcBlockHash() string {
	strIndex := strconv.Itoa(b.Index)
	strHash := b.PrevHash
	strData := b.Data
	strTimestamp := b.Timestamp.String()
	strBits := strconv.FormatInt(int64(b.Bits), 16)
	strNonce := strconv.Itoa(b.Nonce)

	b.BlockHeader = strIndex + strHash + strData + strTimestamp + strBits + strNonce
	b.BlockHash = utils.SHA256(b.BlockHeader)
	return b.BlockHash
}

func (b *Block) calcTarget() *big.Int {
	exponentBytes := (b.Bits >> 24) - 3
	exponentBits := exponentBytes * 8
	coefficient := big.NewInt(int64(b.Bits & 0xffffff))
	return coefficient.Lsh(coefficient, uint(exponentBits))
}
