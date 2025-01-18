package lc92

import (
	"reflect"
	"testing"
)

func newListNode(val int, next *ListNode) *ListNode {
	return &ListNode{val, next}
}

func Test_reverseBetween(t *testing.T) {
	type args struct {
		head  *ListNode
		left  int
		right int
	}

	h1 := newListNode(1, newListNode(2, newListNode(3, newListNode(4, newListNode(5, nil)))))
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{"case1", args{h1, 2, 4}, h1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseBetween(tt.args.head, tt.args.left, tt.args.right); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseBetween() = %v, want %v", got, tt.want)
			}
		})
	}
}
