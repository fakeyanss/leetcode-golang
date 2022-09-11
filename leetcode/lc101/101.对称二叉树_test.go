package lc101

import "testing"

func Test_isSymmetric(t *testing.T) {
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
				&TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val:   2,
						Left:  &TreeNode{Val: 3, Left: nil, Right: nil},
						Right: &TreeNode{Val: 4, Left: nil, Right: nil},
					},
					Right: &TreeNode{
						Val:   2,
						Left:  &TreeNode{Val: 4, Left: nil, Right: nil},
						Right: &TreeNode{Val: 3, Left: nil, Right: nil},
					},
				},
			},
			true,
		},
		{
			"case2",
			args{
				&TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val:   2,
						Left:  nil,
						Right: &TreeNode{Val: 3, Left: nil, Right: nil},
					},
					Right: &TreeNode{
						Val:   2,
						Left:  nil,
						Right: &TreeNode{Val: 3, Left: nil, Right: nil},
					},
				},
			},
			false,
		},
		{
			"case3",
			args{
				&TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val:   2,
						Left:  nil,
						Right: &TreeNode{Val: 3, Left: nil, Right: nil},
					},
					Right: &TreeNode{
						Val:   3,
						Left:  nil,
						Right: &TreeNode{Val: 3, Left: nil, Right: nil},
					},
				},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSymmetric(tt.args.root); got != tt.want {
				t.Errorf("isSymmetric() = %v, want %v", got, tt.want)
			}
		})
	}
}
