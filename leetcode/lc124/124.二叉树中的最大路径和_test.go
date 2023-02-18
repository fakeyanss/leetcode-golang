package lc124

import (
	"testing"

	"github.com/fakeyanss/leetcode-golang/helper"
)

func Test_maxPathSum(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"case1", args{helper.NewTreeNode(1, helper.NewTreeNode(2, nil, nil), helper.NewTreeNode(3, nil, nil))}, 6},
		{"case2", args{helper.NewTreeNode(-10, helper.NewTreeNode(9, nil, nil), helper.NewTreeNode(20, helper.NewTreeNode(15, nil, nil), helper.NewTreeNode(7, nil, nil)))}, 42},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxPathSum(tt.args.root); got != tt.want {
				t.Errorf("maxPathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMaxInt(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"case1", args{1, 2}, 2},
		{"case2", args{1, -2}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxInt(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("MaxInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
