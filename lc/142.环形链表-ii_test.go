package lc

import (
	"reflect"
	"testing"
)

func Test_detectCycle(t *testing.T) {
	type args struct {
		head *ListNode
	}

	h1 := NewListNode(3, nil)
	s1 := NewListNode(2, nil)
	h1.Next = s1
	s1.Next = NewListNode(0, NewListNode(-4, s1))

	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{"case1", args{h1}, s1},
		{"case2", args{NewListNode(1, NewListNode(2, nil))}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := detectCycle(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("detectCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}
