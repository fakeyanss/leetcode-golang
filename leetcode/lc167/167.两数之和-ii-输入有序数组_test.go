package lc167

import (
	"reflect"
	"testing"
)

func Test_twoSumII(t *testing.T) {
	type args struct {
		numbers []int
		target  int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"case1", args{[]int{2, 7, 11, 15}, 9}, []int{1, 2}},
		{"case2", args{[]int{2, 3, 4}, 6}, []int{1, 3}},
		{"case3", args{[]int{-1, 0}, -1}, []int{1, 2}},
		{"case4", args{[]int{2, 7, 11, 15}, 22}, []int{2, 4}},
		{"case5", args{[]int{-1, 0}, -2}, []int{-1, -1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := twoSum(tt.args.numbers, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoSumII() = %v, want %v", got, tt.want)
			}
		})
	}
}
