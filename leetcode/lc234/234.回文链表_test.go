package lc234

import (
	"testing"

	"github.com/fakeYanss/leetcode-golang/helper"
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
		{"case1", args{helper.NewListNode(1, helper.NewListNode(2, helper.NewListNode(2, helper.NewListNode(1, nil))))}, true},
		{"case2", args{helper.NewListNode(1, helper.NewListNode(2, helper.NewListNode(3, nil)))}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome234(tt.args.head); got != tt.want {
				t.Errorf("isPalindrome234() = %v, want %v", got, tt.want)
			}
		})
	}
}
