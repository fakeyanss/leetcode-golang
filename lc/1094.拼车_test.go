package lc

import "testing"

func Test_carPooling(t *testing.T) {
	type args struct {
		trips    [][]int
		capacity int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"case1", args{[][]int{{2, 1, 5}, {3, 3, 7}}, 4}, false},
		{"case1", args{[][]int{{2, 1, 5}, {3, 3, 7}}, 5}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := carPooling(tt.args.trips, tt.args.capacity); got != tt.want {
				t.Errorf("carPooling() = %v, want %v", got, tt.want)
			}
		})
	}
}
