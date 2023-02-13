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
				initialBits: constant.TestGenesisBlock.Bits,
			},
			args: args{
				block: &Block{
					Index:       constant.TestGenesisBlock.Index,
					PrevHash:    constant.TestGenesisBlock.PrevHash,
					Data:        constant.TestGenesisBlock.Data,
					Timestamp:   constant.TestGenesisBlock.Timestamp,
					Bits:        constant.TestGenesisBlock.Bits,
					Nonce:       constant.TestGenesisBlock.Nonce,
					ElapsedTime: constant.TestGenesisBlock.ElapsedTime,
					BlockHash:   constant.TestGenesisBlock.BlockHash,
					BlockHeader: constant.TestGenesisBlock.BlockHeader,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := Blockchain{
				initialBits: tt.fields.initialBits,
			}
			b.Mining(tt.args.block)
		})
	}
}
