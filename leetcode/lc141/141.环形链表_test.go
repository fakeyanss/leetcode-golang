package lc141

import (
	"testing"
)

func newListNode(val int, next *ListNode) *ListNode {
	return &ListNode{val, next}
}

func Test_hasCycle(t *testing.T) {
	type args struct {
		head *ListNode
	}

	var h1, h2, h3 *ListNode
	h1 = newListNode(3, newListNode(2, newListNode(0, newListNode(-4, nil))))
	h1.Next.Next.Next.Next = h1
	h2 = newListNode(1, newListNode(2, nil))
	h2.Next.Next = h2
	h3 = newListNode(1, nil)

	tests := []struct {
		name string
		args args
		want bool
	}{
		{"case1", args{h1}, true},
		{"case2", args{h2}, true},
		{"case3", args{h3}, false},
		{"case4", args{nil}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasCycle(tt.args.head); got != tt.want {
				t.Errorf("hasCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}
