package lib

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math"
	"strconv"
	"time"
)

type Blockchain struct {
	InitialBits int
	chain       []Block
}

func (bc *Blockchain) GetBlockInfo(i int) {
	if len(bc.chain) <= 0 {
		return
	}
	var buf bytes.Buffer
	err := json.Indent(&buf, []byte(bc.chain[i].ToJson()), "", "  ")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(buf.String())
}

func (bc *Blockchain) CreateGenesis() {
	genesisBlock := &Block{
		0,
		"0000000000000000000000000000000000000000000000000000000000000000",
		"Genesis block",
		time.Now(),
		bc.InitialBits,
		0,
		"",
		"",
	}
	bc.mining(genesisBlock)
}

func (bc *Blockchain) AddNewBlock() {
	lastBlock := bc.chain[len(bc.chain)-1]
	b := &Block{
		lastBlock.Index + 1,
		lastBlock.BlockHash,
		"Block " + strconv.Itoa(lastBlock.Index+1),
		time.Now(),
		lastBlock.Bits,
		0,
		"",
		"",
	}
	bc.mining(b)
}

func (bc *Blockchain) mining(b *Block) {
	startTime := time.Now()
	fmt.Println("start mining")

	for i := 1; i <= math.MaxInt64; i++ {
		b.Nonce = i
		if b.CheckValidHash() {
			diff := time.Now().Sub(startTime)
			b.ElapsedTime = diff.String()
			fmt.Println("finish mining: ", diff)
			bc.addBlock(b)
			bc.GetBlockInfo(len(bc.chain) - 1)
			return
		}
	}
}

func (bc *Blockchain) addBlock(b *Block) {
	bc.chain = append(bc.chain, *b)
}
