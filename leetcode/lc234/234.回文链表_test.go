package lc234

import (
	"testing"
)

func newListNode(val int, next *ListNode) *ListNode {
	return &ListNode{val, next}
}

func Test_isPalindrome234(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"case1", args{newListNode(1, newListNode(2, newListNode(2, newListNode(1, nil))))}, true},
		{"case2", args{newListNode(1, newListNode(2, newListNode(3, nil)))}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.args.head); got != tt.want {
				t.Errorf("isPalindrome234() = %v, want %v", got, tt.want)
			}
		})
	}
}
