package lc142

import (
	"reflect"
	"testing"

	"github.com/fakeyanss/leetcode-golang/helper"
)

func Test_detectCycle(t *testing.T) {
	type args struct {
		head *ListNode
	}

	h1 := helper.NewListNode(3, nil)
	s1 := helper.NewListNode(2, nil)
	h1.Next = s1
	s1.Next = helper.NewListNode(0, helper.NewListNode(-4, s1))

	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{"case1", args{h1}, s1},
		{"case2", args{helper.NewListNode(1, helper.NewListNode(2, nil))}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := detectCycle(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("detectCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}
