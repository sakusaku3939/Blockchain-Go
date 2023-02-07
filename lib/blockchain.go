package lib

import (
	"fmt"
	"math"
	"time"
)

type Blockchain struct {
	initialBits int
}

func (chain Blockchain) Mining(b *Block) {
	startTime := time.Now()
	fmt.Println("start mining: ", startTime)

	for i := 1; i < math.MaxInt32; i++ {
		b.Nonce = i
		if b.CheckValidHash() {
			diff := time.Now().Sub(startTime)
			fmt.Println("finish mining: ", diff)

			json, _ := b.ToJson()
			fmt.Println(json)
			return
		}
	}
}
