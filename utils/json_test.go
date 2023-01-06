package utils

import "testing"

func TestFormatJSON(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Indent & Compact JSON",
			args: args{"{\"Index\":0,\"Data\":\"ジェネシスブロック\"}"},
			want: "{\"Index\":0,\"Data\":\"ジェネシスブロック\"}",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := IndentJSON(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("IndentJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			got, err = CompactJSON(got)
			if (err != nil) != tt.wantErr {
				t.Errorf("CompactJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("IndentJSON() got = %v, want %v", got, tt.want)
			}
		})
	}
}
