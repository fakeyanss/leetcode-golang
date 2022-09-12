package lc300

import (
	"testing"
)

func Test_lengthOfLIS(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"case1", args{[]int{10, 9, 2, 5, 3, 7, 101, 18}}, 4},
		{"case2", args{[]int{}}, 0},
		{"case3", args{[]int{10, 9, 2, 5, 3, 7, 101, 6}}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLIS(tt.args.nums); got != tt.want {
				t.Errorf("lengthOfLIS() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dp(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"case1", args{[]int{10, 9, 2, 5, 3, 7, 101, 18}}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := dp(tt.args.nums); got != tt.want {
				t.Errorf("dp() = %v, want %v", got, tt.want)
			}
		})
	}
}
