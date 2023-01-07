package lc1803

import "testing"

func Test_countPairs(t *testing.T) {
	type args struct {
		nums []int
		low  int
		high int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{
				nums: []int{1, 4, 2, 7},
				low:  2,
				high: 6,
			},
			want: 6,
		},
		{
			name: "case2",
			args: args{
				nums: []int{9, 8, 4, 2, 1},
				low:  5,
				high: 14,
			},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countPairs(tt.args.nums, tt.args.low, tt.args.high); got != tt.want {
				t.Errorf("countPairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
