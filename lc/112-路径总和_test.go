package lc

import "testing"

func Test_hasPathSum(t *testing.T) {
	type args struct {
		root      *TreeNode
		targetSum int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"case1",
			args{
				&TreeNode{
					5,
					&TreeNode{
						4,
						&TreeNode{
							11,
							&TreeNode{7, nil, nil},
							&TreeNode{2, nil, nil},
						},
						nil,
					},
					&TreeNode{
						8,
						&TreeNode{13, nil, nil},
						&TreeNode{
							4,
							nil,
							&TreeNode{1, nil, nil},
						},
					},
				},
				22,
			},
			true,
		},
		{
			"case2",
			args{
				&TreeNode{1, &TreeNode{2, nil, nil}, &TreeNode{3, nil, nil}},
				5,
			},
			false,
		},
		{
			"case3",
			args{nil, 0},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasPathSum(tt.args.root, tt.args.targetSum); got != tt.want {
				t.Errorf("hasPathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
