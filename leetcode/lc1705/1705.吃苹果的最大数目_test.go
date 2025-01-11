package lc1705

import "testing"

func Test_eatenApples(t *testing.T) {
	type args struct {
		apples []int
		days   []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				apples: []int{1, 2, 3, 5, 2},
				days:   []int{3, 2, 1, 4, 2},
			},
			want: 7,
		},
		{
			name: "case 2",
			args: args{
				apples: []int{3, 1, 1, 0, 0, 2},
				days:   []int{3, 1, 1, 0, 0, 2},
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := eatenApples(tt.args.apples, tt.args.days); got != tt.want {
				t.Errorf("eatenApples() = %v, want %v", got, tt.want)
			}
		})
	}
}
