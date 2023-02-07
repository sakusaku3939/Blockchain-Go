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
				initialBits: constant.GenesisBlock.Bits,
			},
			args: args{
				block: &Block{
					Index:       constant.GenesisBlock.Index,
					PrevHash:    constant.GenesisBlock.PrevHash,
					Data:        constant.GenesisBlock.Data,
					Timestamp:   constant.GenesisBlock.Timestamp,
					Bits:        constant.GenesisBlock.Bits,
					Nonce:       constant.GenesisBlock.Nonce,
					ElapsedTime: constant.GenesisBlock.ElapsedTime,
					BlockHash:   constant.GenesisBlock.BlockHash,
					BlockHeader: constant.GenesisBlock.BlockHeader,
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
