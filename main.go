package main

import (
	"Blockchain-Go/lib"
	"fmt"
)

const InitialBits = 0x1e777777

func main() {
	bc := &lib.Blockchain{InitialBits: InitialBits}
	fmt.Println("ジェネシスブロックを作成中...")
	bc.CreateGenesis()
}
