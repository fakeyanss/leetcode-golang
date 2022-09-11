package lc112

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
					Val: 5,
					Left: &TreeNode{
						Val: 4,
						Left: &TreeNode{
							Val:   11,
							Left:  &TreeNode{Val: 7, Left: nil, Right: nil},
							Right: &TreeNode{Val: 2, Left: nil, Right: nil},
						},
						Right: nil,
					},
					Right: &TreeNode{
						Val:  8,
						Left: &TreeNode{Val: 13, Left: nil, Right: nil},
						Right: &TreeNode{
							Val:   4,
							Left:  nil,
							Right: &TreeNode{Val: 1, Left: nil, Right: nil},
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
				&TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: nil, Right: nil}, Right: &TreeNode{Val: 3, Left: nil, Right: nil}},
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
