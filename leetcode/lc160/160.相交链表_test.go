package lc160

import (
	"testing"

	"github.com/fakeyanss/leetcode-golang/helper"
)

func Test_getIntersectionNode(t *testing.T) {
	type args struct {
		headA *ListNode
		headB *ListNode
	}

	var h11, h12, i1 *ListNode
	i1 = helper.NewListNode(8, helper.NewListNode(4, helper.NewListNode(5, nil)))
	h11 = helper.NewListNode(4, helper.NewListNode(1, helper.NewListNode(1, i1)))
	h12 = helper.NewListNode(5, helper.NewListNode(6, helper.NewListNode(1, i1)))

	var h21, h22, i2 *ListNode
	i2 = helper.NewListNode(2, helper.NewListNode(4, nil))
	h21 = helper.NewListNode(1, helper.NewListNode(9, helper.NewListNode(1, i2)))
	h22 = helper.NewListNode(3, i2)

	var h31, h32 *ListNode
	h31 = helper.NewListNode(2, helper.NewListNode(6, helper.NewListNode(4, nil)))
	h32 = helper.NewListNode(1, helper.NewListNode(5, nil))

	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{"case1", args{h11, h12}, i1},
		{"case2", args{h21, h22}, i2},
		{"case3", args{h31, h32}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getIntersectionNode(tt.args.headA, tt.args.headB); got != tt.want {
				t.Errorf("getIntersectionNode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getLenOfListNode(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"case1", args{helper.NewListNode(1, helper.NewListNode(2, helper.NewListNode(3, nil)))}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getLenOfListNode(tt.args.head); got != tt.want {
				t.Errorf("getLenOfListNode() = %v, want %v", got, tt.want)
			}
		})
	}
}
