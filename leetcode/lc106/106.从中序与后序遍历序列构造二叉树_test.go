package lc106

import (
	"reflect"
	"testing"
)

func newTreeNode(val int, l *TreeNode, r *TreeNode) *TreeNode {
	return &TreeNode{val, l, r}
}

func Test_buildTree106_(t *testing.T) {
	type args struct {
		inorder   []int
		postorder []int
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			"case1",
			args{
				[]int{9, 3, 15, 20, 7},
				[]int{9, 15, 7, 20, 3},
			},
			newTreeNode(
				3,
				newTreeNode(9, nil, nil),
				newTreeNode(20, newTreeNode(15, nil, nil), newTreeNode(7, nil, nil)),
			),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buildTree(tt.args.inorder, tt.args.postorder); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("buildTree106_() = %v, want %v", got, tt.want)
			}
		})
	}
}
