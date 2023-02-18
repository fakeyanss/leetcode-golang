package lc83

import (
	"reflect"
	"testing"

	"github.com/fakeyanss/leetcode-golang/helper"
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
		{"case1", args{helper.NewListNode(1, helper.NewListNode(1, helper.NewListNode(2, nil)))}, helper.NewListNode(1, helper.NewListNode(2, nil))},
		{"case2", args{helper.NewListNode(1, helper.NewListNode(1, helper.NewListNode(1, helper.NewListNode(2, helper.NewListNode(3, helper.NewListNode(3, nil))))))}, helper.NewListNode(1, helper.NewListNode(2, helper.NewListNode(3, nil)))},
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
