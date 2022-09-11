package lc5

import (
	"testing"
)

func Test_longestPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"case1", args{"babad"}, "bab"},
		{"case2", args{"cbbd"}, "bb"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestPalindrome(tt.args.s); got != tt.want {
				t.Errorf("longestPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_palindrome(t *testing.T) {
	type args struct {
		s string
		l int
		r int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"case1", args{"abba", 1, 2}, "abba"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := palindrome(tt.args.s, tt.args.l, tt.args.r); got != tt.want {
				t.Errorf("palindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
