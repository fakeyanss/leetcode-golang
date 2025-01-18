package lc0086

import (
	"reflect"
	"testing"
)

func Test_partition(t *testing.T) {
	type args struct {
		head *ListNode
		x    int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{"case1", args{newListNode(1, newListNode(4, newListNode(3, newListNode(2, newListNode(5, newListNode(2, nil)))))), 3}, newListNode(1, newListNode(2, newListNode(2, newListNode(4, newListNode(3, newListNode(5, nil))))))},
		{"case2", args{newListNode(2, newListNode(1, nil)), 2}, newListNode(1, newListNode(2, nil))},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := partition(tt.args.head, tt.args.x); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("partition() = %v, want %v", got, tt.want)
			}
		})
	}
}

func newListNode(val int, next *ListNode) *ListNode {
	return &ListNode{val, next}
}
