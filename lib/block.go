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
	strIndex := strconv.Itoa(b.Index)
	strBits := strconv.FormatInt(int64(b.Bits), 16)
	strNonce := strconv.Itoa(b.Nonce)

	b.BlockHeader = strIndex + b.PrevHash + b.Data + b.Timestamp.String() + strBits + strNonce
	b.BlockHash = utils.SHA256(b.BlockHeader)
	return b.BlockHash
}

func (b *Block) CalcTarget() *big.Int {
	exponentBytes := (b.Bits >> 24) - 3
	exponentBits := exponentBytes * 8
	coefficient := big.NewInt(int64(b.Bits & 0xffffff))
	return coefficient.Lsh(coefficient, uint(exponentBits))
}

func (b *Block) CheckValidHash() bool {
	blockHash, err := new(big.Int).SetString(b.CalcBlockHash(), 16)
	if !err {
		fmt.Println("SetString error: blockHash: ", blockHash)
		return false
	}
	return blockHash.Cmp(b.CalcTarget()) != 1 // blockHash <= b.CalcTarget()
}
