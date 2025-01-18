package lc104

import (
	"testing"
)

func newTreeNode(val int, l *TreeNode, r *TreeNode) *TreeNode {
	return &TreeNode{val, l, r}
}

func Test_maxDepth(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"case1", args{newTreeNode(3, newTreeNode(9, nil, nil), newTreeNode(20, newTreeNode(15, nil, nil), newTreeNode(7, nil, nil)))}, 3},
		{"case2", args{newTreeNode(3, newTreeNode(9, newTreeNode(20, newTreeNode(15, nil, nil), nil), nil), newTreeNode(7, nil, nil))}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxDepth(tt.args.root); got != tt.want {
				t.Errorf("maxDepth() = %v, want %v", got, tt.want)
			}
		})
	}
}
