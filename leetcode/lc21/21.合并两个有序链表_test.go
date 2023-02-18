package lc0021

import (
	"github.com/fakeyanss/leetcode-golang/helper"
	"reflect"
	"testing"
)

type TreeNode = helper.ListNode

func TestMergeTwoLists(t *testing.T) {

}

func Test_mergeTwoLists(t *testing.T) {
	type args struct {
		list1 *ListNode
		list2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "case1",
			args: args{
				list1: &ListNode{
					1,
					&ListNode{
						2,
						&ListNode{
							4,
							nil,
						},
					},
				},
				list2: &ListNode{
					1,
					&ListNode{
						3,
						&ListNode{
							4,
							nil,
						},
					},
				},
			},
			want: &ListNode{
				1,
				&ListNode{
					1,
					&ListNode{
						2,
						&ListNode{
							3,
							&ListNode{
								4,
								&ListNode{
									4,
									nil,
								},
							},
						},
					},
				},
			},
		},
		{
			name: "case2",
			args: args{
				list1: nil,
				list2: nil,
			},
			want: nil,
		},
		{
			name: "case3",
			args: args{
				list1: nil,
				list2: &ListNode{
					0,
					nil,
				},
			},
			want: &ListNode{
				0,
				nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := mergeTwoLists(tt.args.list1, tt.args.list2)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeTwoLists() = %v, want %v", got, tt.want)
			}
		})
	}
}
