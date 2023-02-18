package lc104

import (
	"testing"

	"github.com/fakeyanss/leetcode-golang/helper"
)

func Test_maxDepth(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"case1", args{helper.NewTreeNode(3, helper.NewTreeNode(9, nil, nil), helper.NewTreeNode(20, helper.NewTreeNode(15, nil, nil), helper.NewTreeNode(7, nil, nil)))}, 3},
		{"case2", args{helper.NewTreeNode(3, helper.NewTreeNode(9, helper.NewTreeNode(20, helper.NewTreeNode(15, nil, nil), nil), nil), helper.NewTreeNode(7, nil, nil))}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxDepth(tt.args.root); got != tt.want {
				t.Errorf("maxDepth() = %v, want %v", got, tt.want)
			}
		})
	}
}
