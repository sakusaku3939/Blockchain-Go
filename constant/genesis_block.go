package constant

import "time"

type BlockFields struct {
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

var GenesisBlock = BlockFields{
	0,
	"0000000000000000000000000000000000000000000000000000000000000000",
	"ジェネシスブロック",
	time.Date(2022, 4, 1, 0, 0, 0, 0, time.Local),
	0x1e777777,
	0,
	"",
	"",
	"",
}
