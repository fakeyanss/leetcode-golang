package lc1625

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_findLexSmallestString(t *testing.T) {
	type args struct {
		s string
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case1",
			args: args{s: "5525", a: 9, b: 2},
			want: "2050",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findLexSmallestString(tt.args.s, tt.args.a, tt.args.b)
			require.Equal(t, tt.want, got)
		})
	}
}
