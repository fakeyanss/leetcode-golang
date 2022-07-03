package lc

import "testing"

func Test_romanToInt(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{"case1", args{"III"}, 3},
		{"case1", args{"IV"}, 4},
		{"case1", args{"IX"}, 9},
		{"case1", args{"LVIII"}, 58},
		{"case1", args{"MCMXCIV"}, 1994},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := romanToInt(tt.args.s); gotAns != tt.wantAns {
				t.Errorf("romanToInt() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
