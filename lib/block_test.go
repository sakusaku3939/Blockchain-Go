package lib

import (
	"Blockchain-Go/constant"
	"testing"
)

func TestBlock_ToJson(t *testing.T) {
	tests := []struct {
		name    string
		fields  constant.BlockFields
		want    string
		wantErr bool
	}{
		{
			name:   "Generate genesis block",
			fields: constant.GenesisBlock,
			want: "{" +
				"\"Index\":0," +
				"\"PrevHash\":\"0000000000000000\"," +
				"\"Data\":\"ジェネシスブロック\"," +
				"\"Timestamp\":\"2022-04-01T00:00:00+09:00\"," +
				"\"Bits\":\"9777777\"," +
				"\"Nonce\":0," +
				"\"ElapsedTime\":\"\"," +
				"\"BlockHash\":\"\"" +
				"}",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Block{
				Index:       tt.fields.Index,
				PrevHash:    tt.fields.PrevHash,
				Data:        tt.fields.Data,
				Timestamp:   tt.fields.Timestamp,
				Bits:        tt.fields.Bits,
				Nonce:       tt.fields.Nonce,
				ElapsedTime: tt.fields.ElapsedTime,
				BlockHash:   tt.fields.BlockHash,
				BlockHeader: tt.fields.BlockHeader,
			}
			got, err := b.ToJson()
			if (err != nil) != tt.wantErr {
				t.Errorf("ToJson() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ToJson() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBlock_CalcBlockHash(t *testing.T) {
	tests := []struct {
		name   string
		fields constant.BlockFields
		want   string
	}{
		{
			name:   "Calculate block hash",
			fields: constant.GenesisBlock,
			want:   "a0a26ae7d8df4de5dfe368b1be842780a4790308f4c2762d0ce2f2b116b1e711",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Block{
				Index:       tt.fields.Index,
				PrevHash:    tt.fields.PrevHash,
				Data:        tt.fields.Data,
				Timestamp:   tt.fields.Timestamp,
				Bits:        tt.fields.Bits,
				Nonce:       tt.fields.Nonce,
				ElapsedTime: tt.fields.ElapsedTime,
				BlockHash:   tt.fields.BlockHash,
				BlockHeader: tt.fields.BlockHeader,
			}
			if got := b.CalcBlockHash(); got != tt.want {
				t.Errorf("calcBlockHash() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBlock_CalcTarget(t *testing.T) {
	tests := []struct {
		name   string
		fields constant.BlockFields
		want   int64
	}{
		{
			name:   "Calculate target",
			fields: constant.GenesisBlock,
			want:   0x7777770000000000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Block{
				Index:       tt.fields.Index,
				PrevHash:    tt.fields.PrevHash,
				Data:        tt.fields.Data,
				Timestamp:   tt.fields.Timestamp,
				Bits:        tt.fields.Bits,
				Nonce:       tt.fields.Nonce,
				ElapsedTime: tt.fields.ElapsedTime,
				BlockHash:   tt.fields.BlockHash,
				BlockHeader: tt.fields.BlockHeader,
			}
			if got := b.CalcTarget(); got != tt.want {
				t.Errorf("CalcTarget() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBlock_CheckValidHash(t *testing.T) {
	tests := []struct {
		name   string
		fields constant.BlockFields
		want   bool
	}{
		{
			name:   "Check valid hash",
			fields: constant.GenesisBlock,
			want:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Block{
				Index:       tt.fields.Index,
				PrevHash:    tt.fields.PrevHash,
				Data:        tt.fields.Data,
				Timestamp:   tt.fields.Timestamp,
				Bits:        tt.fields.Bits,
				Nonce:       tt.fields.Nonce,
				ElapsedTime: tt.fields.ElapsedTime,
				BlockHash:   tt.fields.BlockHash,
				BlockHeader: tt.fields.BlockHeader,
			}
			if got := b.CheckValidHash(); got != tt.want {
				t.Errorf("CheckValidHash() = %v, want %v", got, tt.want)
			}
		})
	}
}
