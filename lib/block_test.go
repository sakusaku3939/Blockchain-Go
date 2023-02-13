package lib

import (
	"Blockchain-Go/constant"
	"testing"
)

func TestBlock_ToJson(t *testing.T) {
	tests := []struct {
		name    string
		fields  constant.TestBlockFields
		want    string
		wantErr bool
	}{
		{
			name:   "Generate genesis block",
			fields: constant.TestBlock,
			want: "{" +
				"\"Index\":0," +
				"\"PrevHash\":\"0000000000000000000000000000000000000000000000000000000000000000\"," +
				"\"Data\":\"Test block\"," +
				"\"Timestamp\":\"2022-04-01T00:00:00+09:00\"," +
				"\"Bits\":\"1e777777\"," +
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
			if got := b.ToJson(); got != tt.want {
				t.Errorf("ToJson() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBlock_CalcBlockHash(t *testing.T) {
	tests := []struct {
		name   string
		fields constant.TestBlockFields
		want   string
	}{
		{
			name:   "Calculate block hash",
			fields: constant.TestBlock,
			want:   "a60657db75a7ed3e7a4fca925f020fedb2264ba1b28c7094fdfa99051e0d1325",
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
			if got := b.calcBlockHash(); got != tt.want {
				t.Errorf("calcBlockHash() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBlock_CalcTarget(t *testing.T) {
	tests := []struct {
		name    string
		fields  constant.TestBlockFields
		strWant string
	}{
		{
			name:    "Calculate target",
			fields:  constant.TestBlock,
			strWant: "777777000000000000000000000000000000000000000000000000000000",
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
			if got := b.calcTarget().Text(16); got != tt.strWant {
				t.Errorf("calcTarget() = %v, want %v", got, tt.strWant)
			}
		})
	}
}

func TestBlock_CheckValidHash(t *testing.T) {
	tests := []struct {
		name   string
		fields constant.TestBlockFields
		want   bool
	}{
		{
			name:   "Check valid hash",
			fields: constant.TestBlock,
			want:   false,
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
