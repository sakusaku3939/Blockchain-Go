package lib

import (
	"testing"
	"time"
)

func TestBlock_ToJson(t *testing.T) {
	type fields struct {
		Index     int
		PrevHash  string
		Data      string
		Timestamp time.Time
		Bits      string
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		{
			name: "Generate genesis block",
			fields: fields{
				0,
				"0000000000000000000000000000000000000000000000000000000000000000",
				"ジェネシスブロック",
				time.Date(2022, 4, 1, 0, 0, 0, 0, time.Local),
				"1e377777",
			},
			want: "{" +
				"\"Index\":0," +
				"\"PrevHash\":\"0000000000000000000000000000000000000000000000000000000000000000\"," +
				"\"Data\":\"ジェネシスブロック\"," +
				"\"Timestamp\":\"2022-04-01T00:00:00+09:00\"," +
				"\"Bits\":\"1e377777\"," +
				"\"Nonce\":0," +
				"\"ElapsedTime\":\"\"," +
				"\"BlockHash\":\"\"" +
				"}",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := Block{
				Index:     tt.fields.Index,
				PrevHash:  tt.fields.PrevHash,
				Data:      tt.fields.Data,
				Timestamp: tt.fields.Timestamp,
				Bits:      tt.fields.Bits,
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
