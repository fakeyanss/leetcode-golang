package lc114

import (
	"reflect"
	"testing"
)

func newTreeNode(val int, l *TreeNode, r *TreeNode) *TreeNode {
	return &TreeNode{val, l, r}
}

func Test_flatten(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			"case1",
			args{
				newTreeNode(1,
					newTreeNode(2,
						newTreeNode(3, nil, nil),
						newTreeNode(4, nil, nil),
					),
					newTreeNode(5,
						nil,
						newTreeNode(6, nil, nil),
					),
				),
			},
			newTreeNode(1, nil, newTreeNode(2, nil, newTreeNode(3, nil, newTreeNode(4, nil, newTreeNode(5, nil, newTreeNode(6, nil, nil)))))),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if flatten(tt.args.root); !reflect.DeepEqual(tt.args.root, tt.want) {
				t.Errorf("flatten() = %v, want %v", tt.args.root, tt.want)
			}
		})
	}
}
