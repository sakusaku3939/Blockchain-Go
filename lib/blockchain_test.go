package lib

import (
	"Blockchain-Go/constant"
	"testing"
)

func TestBlock_Mining(t *testing.T) {
	type fields struct {
		initialBits int
	}
	type args struct {
		block *Block
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "Mining",
			fields: fields{
				initialBits: constant.TestBlock.Bits,
			},
			args: args{
				block: &Block{
					Index:       constant.TestBlock.Index,
					PrevHash:    constant.TestBlock.PrevHash,
					Data:        constant.TestBlock.Data,
					Timestamp:   constant.TestBlock.Timestamp,
					Bits:        constant.TestBlock.Bits,
					Nonce:       constant.TestBlock.Nonce,
					ElapsedTime: constant.TestBlock.ElapsedTime,
					BlockHash:   constant.TestBlock.BlockHash,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := Blockchain{
				InitialBits: tt.fields.initialBits,
			}
			b.mining(tt.args.block)
		})
	}
}
