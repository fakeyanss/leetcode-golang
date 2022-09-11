package lc232

import (
	"reflect"
	"testing"
)

func TestConstructor(t *testing.T) {
	tests := []struct {
		name string
		want MyQueue
	}{
		{"case1", Constructor()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Constructor(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Constructor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMyQueue_Push(t *testing.T) {
	type args struct {
		x int
	}

	tests := []struct {
		name string
		q    *MyQueue
		args args
	}{
		{"case1", &MyQueue{}, args{1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.q.Push(tt.args.x)
		})
	}
}

func TestMyQueue_Pop(t *testing.T) {
	tests := []struct {
		name string
		q    *MyQueue
		want int
	}{
		{"case1", &MyQueue{[]int{1}, []int{}}, 1},
		{"case2", &MyQueue{[]int{}, []int{1}}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.q.Pop(); got != tt.want {
				t.Errorf("MyQueue.Pop() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMyQueue_Peek(t *testing.T) {
	tests := []struct {
		name string
		q    *MyQueue
		want int
	}{
		{"case1", &MyQueue{[]int{1}, []int{}}, 1},
		{"case2", &MyQueue{[]int{}, []int{1}}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.q.Peek(); got != tt.want {
				t.Errorf("MyQueue.Peek() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMyQueue_Empty(t *testing.T) {
	tests := []struct {
		name string
		q    *MyQueue
		want bool
	}{
		{"case1", &MyQueue{[]int{1}, []int{}}, false},
		{"case2", &MyQueue{}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.q.Empty(); got != tt.want {
				t.Errorf("MyQueue.Empty() = %v, want %v", got, tt.want)
			}
		})
	}
}
