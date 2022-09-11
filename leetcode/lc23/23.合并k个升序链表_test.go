package lc23

import (
	"reflect"
	"testing"

	"github.com/fakeYanss/leetcode-golang/helper"
)

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
					helper.NewListNode(1, helper.NewListNode(4, helper.NewListNode(5, nil))),
					helper.NewListNode(1, helper.NewListNode(3, helper.NewListNode(4, nil))),
					helper.NewListNode(2, helper.NewListNode(6, nil)),
				},
			},
			helper.NewListNode(1, helper.NewListNode(1, helper.NewListNode(2, helper.NewListNode(3, helper.NewListNode(4, helper.NewListNode(4, helper.NewListNode(5, helper.NewListNode(6, nil)))))))),
		},
		{"case2", args{[]*ListNode{helper.NewListNode(1, nil)}}, helper.NewListNode(1, nil)},
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

func Test_mergek(t *testing.T) {
	type args struct {
		list []*ListNode
		l    int
		r    int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{"case1", args{[]*ListNode{helper.NewListNode(1, nil)}, 0, 0}, helper.NewListNode(1, nil)},
		{"case1", args{[]*ListNode{helper.NewListNode(1, nil)}, 1, 0}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergek(tt.args.list, tt.args.l, tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergek() = %v, want %v", got, tt.want)
			}
		})
	}
}
