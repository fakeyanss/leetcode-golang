package lc

import (
	"reflect"
	"testing"
)

func Test_deleteDuplicates(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{"case1", args{NewListNode(1, NewListNode(1, NewListNode(2, nil)))}, NewListNode(1, NewListNode(2, nil))},
		{"case2", args{NewListNode(1, NewListNode(1, NewListNode(1, NewListNode(2, NewListNode(3, NewListNode(3, nil))))))}, NewListNode(1, NewListNode(2, NewListNode(3, nil)))},
		{"case3", args{nil}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := deleteDuplicates(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("deleteDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}
