package lc

import (
	"reflect"
	"testing"
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
		{"case1", args{NewListNode(1, NewListNode(4, NewListNode(3, NewListNode(2, NewListNode(5, NewListNode(2, nil)))))), 3}, NewListNode(1, NewListNode(2, NewListNode(2, NewListNode(4, NewListNode(3, NewListNode(5, nil))))))},
		{"case2", args{NewListNode(2, NewListNode(1, nil)), 2}, NewListNode(1, NewListNode(2, nil))},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := partition(tt.args.head, tt.args.x); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("partition() = %v, want %v", got, tt.want)
			}
		})
	}
}
