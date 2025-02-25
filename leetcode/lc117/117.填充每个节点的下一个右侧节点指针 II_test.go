package lc117

import (
	"reflect"
	"testing"
)

func Test_connect(t *testing.T) {
	type args struct {
		root *Node
	}
	tests := []struct {
		name string
		args args
		want *Node
	}{
		{
			name: "case1",
			args: args{
				root: newNode(
					1,
					newNode(
						2,
						newNode(4, nil, nil, nil),
						newNode(5, nil, nil, nil),
						nil,
					),
					newNode(
						3,
						nil,
						newNode(7, nil, nil, nil),
						nil,
					),
					nil,
				),
			},
			want: func() *Node {
				n1 := newNode(1, nil, nil, nil)
				n2 := newNode(2, nil, nil, nil)
				n3 := newNode(3, nil, nil, nil)
				n4 := newNode(4, nil, nil, nil)
				n5 := newNode(5, nil, nil, nil)
				n7 := newNode(7, nil, nil, nil)
				n1.Left, n1.Right = n2, n3
				n2.Left, n2.Right, n2.Next = n4, n5, n3
				n3.Right = n7
				n4.Next = n5
				n5.Next = n7
				return n1
			}(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := connect(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("connect() = %v, want %v", got, tt.want)
			}
		})
	}
}

func newNode(val int, left, right, next *Node) *Node {
	return &Node{
		Val:   val,
		Left:  left,
		Right: right,
		Next:  next,
	}
}
