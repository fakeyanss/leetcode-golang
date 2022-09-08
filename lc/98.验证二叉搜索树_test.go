package lc

import "testing"

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
				NewTreeNode(2, NewTreeNode(1, nil, nil), NewTreeNode(3, nil, nil)),
			},
			true,
		},
		{
			"case2",
			args{
				NewTreeNode(
					5,
					NewTreeNode(
						1,
						nil,
						nil,
					),
					NewTreeNode(
						4,
						NewTreeNode(
							3,
							nil,
							nil,
						),
						NewTreeNode(
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
				NewTreeNode(
					5,
					NewTreeNode(
						6,
						nil,
						nil,
					),
					NewTreeNode(
						4,
						NewTreeNode(
							3,
							nil,
							nil,
						),
						NewTreeNode(
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
