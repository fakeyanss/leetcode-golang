package lc494

import (
	"testing"
)

func Test_findTargetSumWays(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"case1", args{nums: []int{1, 1, 1, 1, 1}, target: 3}, 5},
		{"case2", args{nums: []int{1}, target: 1}, 1},
		{"case3", args{nums: []int{1}, target: 2}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findTargetSumWays(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("findTargetSumWays() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_backtrack(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"case1", args{nums: []int{1, 1, 1, 1, 1}, target: 3}, 5},
		{"case2", args{nums: []int{1}, target: 1}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := backtrack(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("backtrack() = %v, want %v", got, tt.want)
			}
		})
	}
}
