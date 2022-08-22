package lc

import (
	"reflect"
	"testing"
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
				NewTreeNode(1,
					NewTreeNode(2,
						NewTreeNode(3, nil, nil),
						NewTreeNode(4, nil, nil),
					),
					NewTreeNode(5,
						nil,
						NewTreeNode(6, nil, nil),
					),
				),
			},
			NewTreeNode(1, nil, NewTreeNode(2, nil, NewTreeNode(3, nil, NewTreeNode(4, nil, NewTreeNode(5, nil, NewTreeNode(6, nil, nil)))))),
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
