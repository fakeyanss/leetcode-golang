package lc124

import (
	"testing"
)

func newTreeNode(val int, l *TreeNode, r *TreeNode) *TreeNode {
	return &TreeNode{val, l, r}
}

func Test_maxPathSum(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"case1", args{newTreeNode(1, newTreeNode(2, nil, nil), newTreeNode(3, nil, nil))}, 6},
		{"case2", args{newTreeNode(-10, newTreeNode(9, nil, nil), newTreeNode(20, newTreeNode(15, nil, nil), newTreeNode(7, nil, nil)))}, 42},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxPathSum(tt.args.root); got != tt.want {
				t.Errorf("maxPathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
