package lc

import (
	"testing"
)

func Test_isPalindrome234(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"case1", args{NewListNode(1, NewListNode(2, NewListNode(2, NewListNode(1, nil))))}, true},
		{"case2", args{NewListNode(1, NewListNode(2, NewListNode(3, nil)))}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome234(tt.args.head); got != tt.want {
				t.Errorf("isPalindrome234() = %v, want %v", got, tt.want)
			}
		})
	}
}
