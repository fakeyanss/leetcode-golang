package lc1616

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_checkPalindromeFormation(t *testing.T) {
	type args struct {
		a string
		b string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case1: test single char",
			args: args{
				a: "x",
				b: "y",
			},
			want: true,
		},
		{
			name: "case2: test multiple char",
			args: args{
				a: "ulacfdx",
				b: "jizaluy",
			},
			want: true,
		},
		{
			name: "case3: test none palindrome",
			args: args{
				a: "abc",
				b: "def",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := checkPalindromeFormation(tt.args.a, tt.args.b)
			require.Equal(t, tt.want, got)
		})
	}
}
