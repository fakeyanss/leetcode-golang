package lc98

import (
	"testing"

	"github.com/fakeyanss/leetcode-golang/helper"
)

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
				helper.NewTreeNode(2, helper.NewTreeNode(1, nil, nil), helper.NewTreeNode(3, nil, nil)),
			},
			true,
		},
		{
			"case2",
			args{
				helper.NewTreeNode(
					5,
					helper.NewTreeNode(
						1,
						nil,
						nil,
					),
					helper.NewTreeNode(
						4,
						helper.NewTreeNode(
							3,
							nil,
							nil,
						),
						helper.NewTreeNode(
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
				helper.NewTreeNode(
					5,
					helper.NewTreeNode(
						6,
						nil,
						nil,
					),
					helper.NewTreeNode(
						4,
						helper.NewTreeNode(
							3,
							nil,
							nil,
						),
						helper.NewTreeNode(
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
