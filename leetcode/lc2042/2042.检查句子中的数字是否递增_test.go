package lc2042

import "testing"

func Test_areNumbersAscending(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case1",
			args: args{
				s: "1 box has 3 blue 4 red 6 green and 12 yellow marbles",
			},
			want: true,
		},
		{
			name: "case2",
			args: args{
				s: "hello world 5 x 5",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := areNumbersAscending(tt.args.s); got != tt.want {
				t.Errorf("areNumbersAscending() = %v, want %v", got, tt.want)
			}
		})
	}
}
