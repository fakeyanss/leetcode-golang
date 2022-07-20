package lc

import (
	"reflect"
	"testing"
)

func Test_removeNthFromEnd(t *testing.T) {
	type args struct {
		head *ListNode
		n    int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{"case1", args{NewListNode(1, NewListNode(2, NewListNode(3, NewListNode(4, NewListNode(5, nil))))), 2}, NewListNode(1, NewListNode(2, NewListNode(3, NewListNode(5, nil))))},
		{"case2", args{NewListNode(1, nil), 1}, nil},
		{"case3", args{NewListNode(1, NewListNode(2, nil)), 1}, NewListNode(1, nil)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeNthFromEnd(tt.args.head, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeNthFromEnd() = %v, want %v", got, tt.want)
			}
		})
	}
}
