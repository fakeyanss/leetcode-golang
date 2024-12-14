package lc23

import (
	"reflect"
	"testing"
)

func newListNode(val int, next *ListNode) *ListNode {
	return &ListNode{val, next}
}

func Test_mergeKLists(t *testing.T) {
	type args struct {
		lists []*ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			"case1",
			args{
				[]*ListNode{
					newListNode(1, newListNode(4, newListNode(5, nil))),
					newListNode(1, newListNode(3, newListNode(4, nil))),
					newListNode(2, newListNode(6, nil)),
				},
			},
			newListNode(1, newListNode(1, newListNode(2, newListNode(3, newListNode(4, newListNode(4, newListNode(5, newListNode(6, nil)))))))),
		},
		{"case2", args{[]*ListNode{newListNode(1, nil)}}, newListNode(1, nil)},
		{"case3", args{[]*ListNode{}}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeKLists(tt.args.lists); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeKLists() = %v, want %v", got, tt.want)
			}
		})
	}
}
