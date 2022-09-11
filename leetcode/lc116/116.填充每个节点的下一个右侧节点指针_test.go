package lc116

import (
	"reflect"
	"testing"
)

func Test_connect(t *testing.T) {
	type args struct {
		root *Node
	}

	n1 := &Node{Val: 1, Left: nil, Right: nil, Next: nil}
	n2 := &Node{Val: 2, Left: nil, Right: nil, Next: nil}
	n3 := &Node{Val: 3, Left: nil, Right: nil, Next: nil}
	n4 := &Node{Val: 4, Left: nil, Right: nil, Next: nil}
	n5 := &Node{Val: 5, Left: nil, Right: nil, Next: nil}
	n6 := &Node{Val: 6, Left: nil, Right: nil, Next: nil}
	n7 := &Node{Val: 7, Left: nil, Right: nil, Next: nil}
	n1.Left, n1.Right = n2, n3
	n2.Left, n2.Right = n4, n5
	n3.Left, n3.Right = n6, n7

	w1 := &Node{Val: 1, Left: nil, Right: nil, Next: nil}
	w2 := &Node{Val: 2, Left: nil, Right: nil, Next: nil}
	w3 := &Node{Val: 3, Left: nil, Right: nil, Next: nil}
	w4 := &Node{Val: 4, Left: nil, Right: nil, Next: nil}
	w5 := &Node{Val: 5, Left: nil, Right: nil, Next: nil}
	w6 := &Node{Val: 6, Left: nil, Right: nil, Next: nil}
	w7 := &Node{Val: 7, Left: nil, Right: nil, Next: nil}
	w1.Left, w1.Right, w1.Next = w2, w3, nil
	w2.Left, w2.Right, w2.Next = w4, w5, w3
	w3.Left, w3.Right, w3.Next = w6, w7, nil
	w4.Next = w5
	w5.Next = w6
	w6.Next = w7

	tests := []struct {
		name string
		args args
		want *Node
	}{
		{
			"case1",
			args{n1},
			w1,
		},
		{
			"case2",
			args{nil},
			nil,
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
