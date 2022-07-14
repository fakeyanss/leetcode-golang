package lc

import (
	"reflect"
	"testing"
)

func Test_deleteNode(t *testing.T) {
	type args struct {
		node *ListNode
	}

	c1 := NewListNode(5, nil)
	r1 := NewListNode(4, c1)
	c1.Next = NewListNode(1, NewListNode(9, nil))

	c2 := NewListNode(1, nil)
	r2 := NewListNode(4, NewListNode(5, c2))
	c2.Next = NewListNode(9, nil)

	tests := []struct {
		name string
		args args
		got  *ListNode
		want *ListNode
	}{
		{"case1", args{c1}, r1, NewListNode(4, NewListNode(1, NewListNode(9, nil)))},
		{"case2", args{c2}, r2, NewListNode(4, NewListNode(5, NewListNode(9, nil)))},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if deleteNode(tt.args.node); !reflect.DeepEqual(tt.got, tt.want) {
				t.Errorf("deleteNode(), got = %v, want %v", tt.got, tt.want)
			}
		})
	}
}
