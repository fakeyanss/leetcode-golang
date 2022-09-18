// https://ghproxy.com/https://raw.githubusercontent.com/fakeYanss/imgplace/master//2022/2022-09-18_20220918143300.png

package lc6184

import (
	"testing"
)

func Test_countDaysTogether(t *testing.T) {
	type args struct {
		arriveAlice string
		leaveAlice  string
		arriveBob   string
		leaveBob    string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"case1", args{arriveAlice: "08-15", leaveAlice: "08-18", arriveBob: "08-16", leaveBob: "08-19"}, 3},
		{"case1", args{arriveAlice: "10-01", leaveAlice: "10-31", arriveBob: "11-01", leaveBob: "12-31"}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countDaysTogether(tt.args.arriveAlice, tt.args.leaveAlice, tt.args.arriveBob, tt.args.leaveBob); got != tt.want {
				t.Errorf("countDaysTogether() = %v, want %v", got, tt.want)
			}
		})
	}
	// fmt.Println(calcDays("10-31"))
	// fmt.Println(calcDays("11-01"))
}
