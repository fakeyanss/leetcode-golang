package lc303

import (
	"reflect"
	"testing"
)

func TestConstructor303(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want NumArray
	}{
		{"case1", args{[]int{1, 2, 3}}, NumArray{[]int{0, 1, 3, 6}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Constructor303(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Constructor303() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNumArray_SumRange(t *testing.T) {
	type fields struct {
		preSum []int
	}
	type args struct {
		left  int
		right int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{"case1", fields{[]int{0, -2, 0, 1, -4, -2, -3}}, args{0, 2}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nu := &NumArray{
				preSum: tt.fields.preSum,
			}
			if got := nu.SumRange(tt.args.left, tt.args.right); got != tt.want {
				t.Errorf("NumArray.SumRange() = %v, want %v", got, tt.want)
			}
		})
	}
}
