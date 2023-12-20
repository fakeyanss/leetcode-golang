package lc2828

import "testing"

func Test_isAcronym(t *testing.T) {
	type args struct {
		words []string
		s     string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case 1",
			args: args{
				words: []string{"dog", "cat", "apple"},
				s:     "dca",
			},
			want: true,
		},
		{
			name: "case 2",
			args: args{
				words: []string{"an", "apple"},
				s:     "a",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isAcronym(tt.args.words, tt.args.s); got != tt.want {
				t.Errorf("isAcronym() = %v, want %v", got, tt.want)
			}
		})
	}
}
