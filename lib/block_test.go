package lib

import (
	"testing"
	"time"
)

type fields struct {
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

var genesisBlock = fields{
	0,
	"0000000000000000",
	"ジェネシスブロック",
	time.Date(2022, 4, 1, 0, 0, 0, 0, time.Local),
	0x9777777,
	0,
	"",
	"",
	"",
}

func TestBlock_ToJson(t *testing.T) {

	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		{
			name:   "Generate genesis block",
			fields: genesisBlock,
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
			b := Block{
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
		fields fields
		want   string
	}{
		{
			name:   "Calculate block hash",
			fields: genesisBlock,
			want:   "a0a26ae7d8df4de5dfe368b1be842780a4790308f4c2762d0ce2f2b116b1e711",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := Block{
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
		fields fields
		want   int
	}{
		{
			name:   "Calculate target",
			fields: genesisBlock,
			want:   0x7777770000000000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := Block{
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
