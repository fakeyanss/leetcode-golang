package helper

import (
	"reflect"
	"testing"
)

func TestNewListNode(t *testing.T) {
	type args struct {
		val  int
		next *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{"case1", args{1, nil}, &ListNode{1, nil}},
		{"case2", args{1, &ListNode{2, nil}}, &ListNode{1, &ListNode{2, nil}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewListNode(tt.args.val, tt.args.next); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("helper.NewListNode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewTreeNode(t *testing.T) {
	type args struct {
		val int
		l   *TreeNode
		r   *TreeNode
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{"case1", args{1, &TreeNode{2, nil, nil}, &TreeNode{3, nil, nil}}, &TreeNode{1, &TreeNode{2, nil, nil}, &TreeNode{3, nil, nil}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewTreeNode(tt.args.val, tt.args.l, tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("helper.NewTreeNode() = %v, want %v", got, tt.want)
			}
		})
	}
}
