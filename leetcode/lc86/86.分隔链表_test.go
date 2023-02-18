package lc0086

import (
	"reflect"
	"testing"

	"github.com/fakeyanss/leetcode-golang/helper"
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
		{"case1", args{helper.NewListNode(1, helper.NewListNode(4, helper.NewListNode(3, helper.NewListNode(2, helper.NewListNode(5, helper.NewListNode(2, nil)))))), 3}, helper.NewListNode(1, helper.NewListNode(2, helper.NewListNode(2, helper.NewListNode(4, helper.NewListNode(3, helper.NewListNode(5, nil))))))},
		{"case2", args{helper.NewListNode(2, helper.NewListNode(1, nil)), 2}, helper.NewListNode(1, helper.NewListNode(2, nil))},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := partition(tt.args.head, tt.args.x); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("partition() = %v, want %v", got, tt.want)
			}
		})
	}
}
