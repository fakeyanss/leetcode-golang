package lc92

import (
	"reflect"
	"testing"

	"github.com/fakeyanss/leetcode-golang/helper"
)

func Test_reverseBetween(t *testing.T) {
	type args struct {
		head  *ListNode
		left  int
		right int
	}

	h1 := helper.NewListNode(1, helper.NewListNode(2, helper.NewListNode(3, helper.NewListNode(4, helper.NewListNode(5, nil)))))
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{"case1", args{h1, 2, 4}, h1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseBetween(tt.args.head, tt.args.left, tt.args.right); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseBetween() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_reverseK(t *testing.T) {
	type args struct {
		head *ListNode
		k    int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseK(tt.args.head, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseK() = %v, want %v", got, tt.want)
			}
		})
	}
}
