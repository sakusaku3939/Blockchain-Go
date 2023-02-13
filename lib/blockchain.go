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
	Chain       []Block
}

func (bc *Blockchain) AddBlock(b *Block) {
	bc.Chain = append(bc.Chain, *b)
}

func (bc *Blockchain) GetBlockInfo(i int) {
	if len(bc.Chain) <= 0 {
		return
	}
	var buf bytes.Buffer
	err := json.Indent(&buf, []byte(bc.Chain[i].ToJson()), "", "  ")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(buf.String())
}

func (bc *Blockchain) Mining(b *Block) {
	startTime := time.Now()
	fmt.Println("start mining")

	for i := 1; i <= math.MaxInt64; i++ {
		b.Nonce = i
		if b.CheckValidHash() {
			diff := time.Now().Sub(startTime)
			fmt.Println("finish mining: ", diff)
			b.ElapsedTime = diff.String()
			bc.AddBlock(b)
			bc.GetBlockInfo(len(bc.Chain) - 1)
			return
		}
	}
}

func (bc *Blockchain) CreateGenesis() {
	genesisBlock := &Block{
		0,
		"0000000000000000000000000000000000000000000000000000000000000000",
		"Genesis block",
		time.Now(),
		0x1e777777,
		0,
		"",
		"",
		"",
	}
	bc.Mining(genesisBlock)
}

func (bc *Blockchain) AddNewBlock() {
	lastBlock := bc.Chain[len(bc.Chain)-1]
	b := &Block{
		lastBlock.Index + 1,
		lastBlock.BlockHash,
		"Block " + strconv.Itoa(lastBlock.Index+1),
		time.Now(),
		lastBlock.Bits,
		0,
		"",
		"",
		"",
	}
	bc.Mining(b)
}
