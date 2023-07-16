package lc2679

import "testing"

func Test_matrixSum(t *testing.T) {
	type args struct {
		nums [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{nums: [][]int{{7, 2, 1}, {6, 4, 2}, {6, 5, 3}, {3, 2, 1}}},
			want: 15,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := matrixSum(tt.args.nums); got != tt.want {
				t.Errorf("matrixSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
