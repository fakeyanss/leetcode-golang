package lc

import (
	"reflect"
	"testing"
)

func Test_inorderTraversal(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"case1",
			args{
				&TreeNode{
					1,
					nil,
					&TreeNode{
						2,
						&TreeNode{
							3,
							nil,
							nil,
						},
						nil,
					},
				},
			},
			[]int{1, 3, 2},
		},
		{
			"case2",
			args{
				nil,
			},
			[]int{},
		},
		{
			"case3",
			args{
				&TreeNode{
					1,
					nil,
					nil,
				},
			},
			[]int{1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := inorderTraversal(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("inorderTraversal() = %v, want %v", got, tt.want)
			}
		})
	}
}
