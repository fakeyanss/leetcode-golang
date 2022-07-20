package lc

import (
	"reflect"
	"testing"
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
					NewListNode(1, NewListNode(4, NewListNode(5, nil))),
					NewListNode(1, NewListNode(3, NewListNode(4, nil))),
					NewListNode(2, NewListNode(6, nil)),
				},
			},
			NewListNode(1, NewListNode(1, NewListNode(2, NewListNode(3, NewListNode(4, NewListNode(4, NewListNode(5, NewListNode(6, nil)))))))),
		},
		{"case2", args{[]*ListNode{NewListNode(1, nil)}}, NewListNode(1, nil)},
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
		{"case1", args{[]*ListNode{NewListNode(1, nil)}, 0, 0}, NewListNode(1, nil)},
		{"case1", args{[]*ListNode{NewListNode(1, nil)}, 1, 0}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergek(tt.args.list, tt.args.l, tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergek() = %v, want %v", got, tt.want)
			}
		})
	}
}
