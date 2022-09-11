package lc125

import (
	"testing"
)

func Test_isPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"case1", args{"A man, a plan, a canal: Panama"}, true},
		{"case2", args{"race a car"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.args.s); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isAlnum(t *testing.T) {
	type args struct {
		ch byte
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"case1", args{'A'}, true},
		{"case2", args{'a'}, true},
		{"case1", args{'1'}, true},
		{"case1", args{'-'}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isAlnum(tt.args.ch); got != tt.want {
				t.Errorf("isAlnum() = %v, want %v", got, tt.want)
			}
		})
	}
}
