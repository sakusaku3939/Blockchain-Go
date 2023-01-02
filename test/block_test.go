package test

import (
	"Blockchain-Go/lib"
	"fmt"
	"testing"
	"time"
)

const InitialBits = 0x1e377777

func TestToJson(t *testing.T) {
	b := lib.NewBlock(0, "0000000000000000000000000000000000000000000000000000000000000000", "ジェネシスブロック", time.Now(), InitialBits)

	fmt.Println(b)

	json, err := b.ToJson()
	if err != nil {
		t.Error(err)
		return
	}

	fmt.Println(string(json))
}
