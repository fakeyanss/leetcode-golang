package lc235

import (
	"reflect"
	"testing"

	"github.com/fakeyanss/leetcode-golang/helper"
)

func Test_lowestCommonAncestor(t *testing.T) {
	type args struct {
		root *TreeNode
		p    *TreeNode
		q    *TreeNode
	}

	root := helper.NewTreeNode(6, nil, nil)
	p1 := helper.NewTreeNode(2, nil, nil)
	q1 := helper.NewTreeNode(8, nil, nil)
	w1 := root
	p2 := p1
	q2 := helper.NewTreeNode(4, nil, nil)
	w2 := p2

	root.Left = p1
	root.Right = q1
	p1.Left = helper.NewTreeNode(0, nil, nil)
	p1.Right = q2
	q2.Left = helper.NewTreeNode(3, nil, nil)
	q2.Right = helper.NewTreeNode(5, nil, nil)
	q1.Left = helper.NewTreeNode(7, nil, nil)
	q1.Right = helper.NewTreeNode(9, nil, nil)

	p3 := root.Right
	q3 := root.Right.Right

	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{"case1", args{root, p1, q1}, w1},
		{"case2", args{root, p2, q2}, w2},
		{"case3", args{root, p3, q3}, p3},
		{"case4", args{nil, p3, q3}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lowestCommonAncestor(tt.args.root, tt.args.p, tt.args.q); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("lowestCommonAncestor() = %v, want %v", got, tt.want)
			}
		})
	}
}
