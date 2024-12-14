package lc954

import "testing"

func Test_canReorderDoubled(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case1",
			args: args{
				arr: []int{4, -2, 2, -4},
			},
			want: true,
		},
		{
			name: "case2",
			args: args{
				arr: []int{2, 1, 2, 1, 1, 1, 2, 2},
			},
			want: true,
		},
		{
			name: "case3",
			args: args{
				arr: []int{3, 1, 3, 6},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canReorderDoubled(tt.args.arr); got != tt.want {
				t.Errorf("canReorderDoubled() = %v, want %v", got, tt.want)
			}
		})
	}
}
