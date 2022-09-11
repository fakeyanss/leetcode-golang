package lc876

import (
	"reflect"
	"testing"

	"github.com/fakeYanss/leetcode-golang/helper"
)

func Test_middleNode(t *testing.T) {
	type args struct {
		head *ListNode
	}

	h1 := helper.NewListNode(1, helper.NewListNode(2, nil))
	m1 := helper.NewListNode(3, nil)
	h1.Next.Next = m1
	m1.Next = helper.NewListNode(4, helper.NewListNode(5, nil))

	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{"case1", args{h1}, m1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := middleNode(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("middleNode() = %v, want %v", got, tt.want)
			}
		})
	}
}
