package lc114

import (
	"reflect"
	"testing"

	"github.com/fakeyanss/leetcode-golang/helper"
)

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
				helper.NewTreeNode(1,
					helper.NewTreeNode(2,
						helper.NewTreeNode(3, nil, nil),
						helper.NewTreeNode(4, nil, nil),
					),
					helper.NewTreeNode(5,
						nil,
						helper.NewTreeNode(6, nil, nil),
					),
				),
			},
			helper.NewTreeNode(1, nil, helper.NewTreeNode(2, nil, helper.NewTreeNode(3, nil, helper.NewTreeNode(4, nil, helper.NewTreeNode(5, nil, helper.NewTreeNode(6, nil, nil)))))),
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
