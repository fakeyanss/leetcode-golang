package lc1312

import (
	"testing"
)

func Test_minInsertions(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"case1", args{s: "zzazz"}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minInsertions(tt.args.s); got != tt.want {
				t.Errorf("minInsertions() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dpMemo(t *testing.T) {
	type args struct {
		memo [][]int
		s    string
		i    int
		j    int
	}
	memo1 := make([][]int, 5)
	for i := range memo1 {
		memo1[i] = make([]int, 5)
	}
	memo2 := make([][]int, 8)
	for i := range memo2 {
		memo2[i] = make([]int, 8)
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"case1", args{memo: memo1, s: "mbadm", i: 0, j: 4}, 2},
		{"case2", args{memo: memo2, s: "leetcode", i: 0, j: 7}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := dpMemo(tt.args.memo, tt.args.s, tt.args.i, tt.args.j); got != tt.want {
				t.Errorf("dpMemo() = %v, want %v", got, tt.want)
			}
		})
	}
}
