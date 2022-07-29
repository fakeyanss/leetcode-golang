package lc

import (
	"reflect"
	"testing"
)

func TestConstructor304(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name string
		args args
		want NumMatrix
	}{
		{"case1", args{[][]int{{1, 2}, {3, 4}}}, NumMatrix{[][]int{{0, 0, 0}, {0, 1, 3}, {0, 4, 10}}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Constructor304(tt.args.matrix); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Constructor304() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNumMatrix_SumRegion(t *testing.T) {
	type fields struct {
		preSum [][]int
	}
	type args struct {
		row1 int
		col1 int
		row2 int
		col2 int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{"case1", fields{[][]int{{0, 0, 0}, {0, 1, 3}, {0, 4, 10}}}, args{1, 1, 1, 1}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nu := &NumMatrix{
				preSum: tt.fields.preSum,
			}
			if got := nu.SumRegion(tt.args.row1, tt.args.col1, tt.args.row2, tt.args.col2); got != tt.want {
				t.Errorf("NumMatrix.SumRegion() = %v, want %v", got, tt.want)
			}
		})
	}
}
