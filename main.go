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

	for i := 2; i <= 10; i++ {
		fmt.Printf("%d番目のブロックを作成中...\n", i)
		bc.AddNewBlock()
	}
}
