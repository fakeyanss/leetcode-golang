package lc2037

import "testing"

func Test_minMovesToSeat(t *testing.T) {
	type args struct {
		seats    []int
		students []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{
				seats:    []int{3, 1, 5},
				students: []int{2, 7, 4},
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minMovesToSeat(tt.args.seats, tt.args.students); got != tt.want {
				t.Errorf("minMovesToSeat() = %v, want %v", got, tt.want)
			}
		})
	}
}
