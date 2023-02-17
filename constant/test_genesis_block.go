package constant

import "time"

type TestBlockFields struct {
	Index       int
	PrevHash    string
	Data        string
	Timestamp   time.Time
	Bits        int
	Nonce       int
	ElapsedTime string
	BlockHash   string
}

var TestBlock = TestBlockFields{
	0,
	"0000000000000000000000000000000000000000000000000000000000000000",
	"Test block",
	time.Date(2022, 4, 1, 0, 0, 0, 0, time.Local),
	0x1e777777,
	0,
	"",
	"",
}
