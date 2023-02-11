package lib

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math"
	"time"
)

type Blockchain struct {
	initialBits int
	chain       []Block
}

func (c Blockchain) AddBlock(b *Block) {
	c.chain = append(c.chain, *b)
}

func (c Blockchain) GetBlockInfo(i int) string {
	if len(c.chain) <= 0 {
		return ""
	}
	var buf bytes.Buffer
	err := json.Indent(&buf, []byte(c.chain[i].ToJson()), "", "  ")
	if err != nil {
		fmt.Println(err)
	}
	return buf.String()
}

func (c Blockchain) Mining(b *Block) {
	startTime := time.Now()
	fmt.Println("start mining")

	for i := 1; i < math.MaxInt32; i++ {
		b.Nonce = i
		if b.CheckValidHash() {
			diff := time.Now().Sub(startTime)
			fmt.Println("finish mining: ", diff)

			fmt.Println(b.ToJson())
			return
		}
	}
}
