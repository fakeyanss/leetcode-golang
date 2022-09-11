package lc237

import (
	"reflect"
	"testing"

	"github.com/fakeYanss/leetcode-golang/helper"
)

func Test_deleteNode(t *testing.T) {
	type args struct {
		node *ListNode
	}

	c1 := helper.NewListNode(5, nil)
	r1 := helper.NewListNode(4, c1)
	c1.Next = helper.NewListNode(1, helper.NewListNode(9, nil))

	c2 := helper.NewListNode(1, nil)
	r2 := helper.NewListNode(4, helper.NewListNode(5, c2))
	c2.Next = helper.NewListNode(9, nil)

	tests := []struct {
		name string
		args args
		got  *ListNode
		want *ListNode
	}{
		{"case1", args{c1}, r1, helper.NewListNode(4, helper.NewListNode(1, helper.NewListNode(9, nil)))},
		{"case2", args{c2}, r2, helper.NewListNode(4, helper.NewListNode(5, helper.NewListNode(9, nil)))},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if deleteNode(tt.args.node); !reflect.DeepEqual(tt.got, tt.want) {
				t.Errorf("deleteNode(), got = %v, want %v", tt.got, tt.want)
			}
		})
	}
}
