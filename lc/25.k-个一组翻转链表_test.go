package lc

import (
	"reflect"
	"testing"
)

func Test_reverseKGroup(t *testing.T) {
	type args struct {
		head *ListNode
		k    int
	}

	h1 := NewListNode(1, nil)
	r1 := NewListNode(2, nil)
	h1.Next = r1
	r1.Next = NewListNode(3, NewListNode(4, NewListNode(5, nil)))

	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{"case1", args{h1, 2}, r1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseKGroup(tt.args.head, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseKGroup() = %v, want %v", got, tt.want)
			}
		})
	}
}
