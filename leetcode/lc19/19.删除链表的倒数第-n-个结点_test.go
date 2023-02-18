package lc0019

import (
	"reflect"
	"testing"

	"github.com/fakeyanss/leetcode-golang/helper"
)

func Test_removeNthFromEnd(t *testing.T) {
	type args struct {
		head *ListNode
		n    int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{"case1", args{helper.NewListNode(1, helper.NewListNode(2, helper.NewListNode(3, helper.NewListNode(4, helper.NewListNode(5, nil))))), 2}, helper.NewListNode(1, helper.NewListNode(2, helper.NewListNode(3, helper.NewListNode(5, nil))))},
		{"case2", args{helper.NewListNode(1, nil), 1}, nil},
		{"case3", args{helper.NewListNode(1, helper.NewListNode(2, nil)), 1}, helper.NewListNode(1, nil)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeNthFromEnd(tt.args.head, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeNthFromEnd() = %v, want %v", got, tt.want)
			}
		})
	}
}
