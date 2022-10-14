package lc940

import "testing"

func Test_distinctSubseqII(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"case1", args{s: "abc"}, 7},
		{"case1", args{s: "aba"}, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := distinctSubseqII(tt.args.s); got != tt.want {
				t.Errorf("distinctSubseqII() = %v, want %v", got, tt.want)
			}
		})
	}
}
