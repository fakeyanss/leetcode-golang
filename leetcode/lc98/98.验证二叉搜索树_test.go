package lc98

import (
	"testing"
)

func newTreeNode(val int, l *TreeNode, r *TreeNode) *TreeNode {
	return &TreeNode{val, l, r}
}

func Test_isValidBST(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"case1",
			args{
				newTreeNode(2, newTreeNode(1, nil, nil), newTreeNode(3, nil, nil)),
			},
			true,
		},
		{
			"case2",
			args{
				newTreeNode(
					5,
					newTreeNode(
						1,
						nil,
						nil,
					),
					newTreeNode(
						4,
						newTreeNode(
							3,
							nil,
							nil,
						),
						newTreeNode(
							5,
							nil,
							nil,
						),
					),
				),
			},
			false,
		},
		{
			"case3",
			args{
				newTreeNode(
					5,
					newTreeNode(
						6,
						nil,
						nil,
					),
					newTreeNode(
						4,
						newTreeNode(
							3,
							nil,
							nil,
						),
						newTreeNode(
							5,
							nil,
							nil,
						),
					),
				),
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidBST(tt.args.root); got != tt.want {
				t.Errorf("isValidBST() = %v, want %v", got, tt.want)
			}
		})
	}
}
