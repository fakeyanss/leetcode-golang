package helper

import "testing"

func TestArrayEqual(t *testing.T) {
	type args struct {
		a []int
		b []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"case1", args{[]int{1}, []int{2}}, false},
		{"case2", args{[]int{1}, []int{1}}, true},
		{"case3", args{nil, []int{1}}, false},
		{"case4", args{[]int{1, 2}, []int{1}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ArrayEqual(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("ArrayEqual() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPrintListNode(t *testing.T) {
	type args struct {
		l *ListNode
	}
	tests := []struct {
		name string
		args args
	}{
		{"case1", args{NewListNode(1, NewListNode(2, NewListNode(3, nil)))}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			PrintListNode(tt.args.l)
		})
	}
}
