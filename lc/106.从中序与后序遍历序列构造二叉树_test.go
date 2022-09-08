package lc

import (
	"reflect"
	"testing"
)

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
			NewTreeNode(
				3,
				NewTreeNode(9, nil, nil),
				NewTreeNode(20, NewTreeNode(15, nil, nil), NewTreeNode(7, nil, nil)),
			),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buildTree106_(tt.args.inorder, tt.args.postorder); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("buildTree106_() = %v, want %v", got, tt.want)
			}
		})
	}
}
