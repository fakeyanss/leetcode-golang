package lc

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
					1,
					&TreeNode{
						2,
						&TreeNode{3, nil, nil},
						&TreeNode{4, nil, nil},
					},
					&TreeNode{
						2,
						&TreeNode{4, nil, nil},
						&TreeNode{3, nil, nil},
					},
				},
			},
			true,
		},
		{
			"case2",
			args{
				&TreeNode{
					1,
					&TreeNode{
						2,
						nil,
						&TreeNode{3, nil, nil},
					},
					&TreeNode{
						2,
						nil,
						&TreeNode{3, nil, nil},
					},
				},
			},
			false,
		},
		{
			"case3",
			args{
				&TreeNode{
					1,
					&TreeNode{
						2,
						nil,
						&TreeNode{3, nil, nil},
					},
					&TreeNode{
						3,
						nil,
						&TreeNode{3, nil, nil},
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
