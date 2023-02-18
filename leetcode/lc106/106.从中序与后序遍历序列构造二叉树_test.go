package lc106

import (
	"reflect"
	"testing"

	"github.com/fakeyanss/leetcode-golang/helper"
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
			helper.NewTreeNode(
				3,
				helper.NewTreeNode(9, nil, nil),
				helper.NewTreeNode(20, helper.NewTreeNode(15, nil, nil), helper.NewTreeNode(7, nil, nil)),
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
