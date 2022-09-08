package lc

import (
	"reflect"
	"testing"
)

func Test_connect(t *testing.T) {
	type args struct {
		root *Node
	}

	n1 := &Node{1, nil, nil, nil}
	n2 := &Node{2, nil, nil, nil}
	n3 := &Node{3, nil, nil, nil}
	n4 := &Node{4, nil, nil, nil}
	n5 := &Node{5, nil, nil, nil}
	n6 := &Node{6, nil, nil, nil}
	n7 := &Node{7, nil, nil, nil}
	n1.Left, n1.Right = n2, n3
	n2.Left, n2.Right = n4, n5
	n3.Left, n3.Right = n6, n7

	w1 := &Node{1, nil, nil, nil}
	w2 := &Node{2, nil, nil, nil}
	w3 := &Node{3, nil, nil, nil}
	w4 := &Node{4, nil, nil, nil}
	w5 := &Node{5, nil, nil, nil}
	w6 := &Node{6, nil, nil, nil}
	w7 := &Node{7, nil, nil, nil}
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
