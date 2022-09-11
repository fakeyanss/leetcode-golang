package lc95

import (
	"reflect"
	"testing"

	"github.com/fakeYanss/leetcode-golang/helper"
)

func Test_generateTrees(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want []*TreeNode
	}{
		{
			"case1",
			args{3},
			[]*TreeNode{
				helper.NewTreeNode(1, nil, helper.NewTreeNode(3, helper.NewTreeNode(2, nil, nil), nil)),
				helper.NewTreeNode(1, nil, helper.NewTreeNode(2, nil, helper.NewTreeNode(3, nil, nil))),
				helper.NewTreeNode(2, helper.NewTreeNode(1, nil, nil), helper.NewTreeNode(3, nil, nil)),
				helper.NewTreeNode(3, helper.NewTreeNode(2, helper.NewTreeNode(1, nil, nil), nil), nil),
				helper.NewTreeNode(3, helper.NewTreeNode(1, nil, helper.NewTreeNode(2, nil, nil)), nil),
			},
		},
		{
			"case2",
			args{0},
			[]*TreeNode{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := generateTrees(tt.args.n)
			if len(got) != len(tt.want) {
				t.Errorf("generateTrees() = %v, want %v", got, tt.want)
				return
			}

			errList := make([]bool, len(got))
			for i, v := range got {
				err := true
				for _, w := range tt.want {
					if reflect.DeepEqual(v, w) {
						err = false
					}
				}
				errList[i] = err
			}

			for _, e := range errList {
				if e {
					t.Errorf("generateTrees() = %v, want %v", got, tt.want)
					return
				}
			}
		})
	}
}
