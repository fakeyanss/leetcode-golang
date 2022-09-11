package lc344

import (
	"reflect"
	"testing"
)

func Test_reverseString(t *testing.T) {
	type args struct {
		s []byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{"case1", args{[]byte{'h', 'e', 'l', 'l', 'o'}}, []byte{'o', 'l', 'l', 'e', 'h'}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if reverseString(tt.args.s); !reflect.DeepEqual(tt.args.s, tt.want) {
				t.Errorf("reverseString() = %v, want %v", tt.args.s, tt.want)
			}
		})
	}
}
