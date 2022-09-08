package lc

import (
	"reflect"
	"testing"
)

func Test_buildTree(t *testing.T) {
	type args struct {
		preorder []int
		inorder  []int
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			"case1",
			args{
				[]int{3, 9, 20, 15, 7},
				[]int{9, 3, 15, 20, 7},
			},
			NewTreeNode(
				3,
				NewTreeNode(
					9, nil, nil),
				NewTreeNode(
					20,
					NewTreeNode(15, nil, nil),
					NewTreeNode(7, nil, nil),
				),
			),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buildTree(tt.args.preorder, tt.args.inorder); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("buildTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
