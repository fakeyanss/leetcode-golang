package lc1104

import (
	"reflect"
	"testing"
)

func Test_pathInZigZagTree(t *testing.T) {
	type args struct {
		label int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test1",
			args: args{label: 14},
			want: []int{1, 3, 4, 14},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pathInZigZagTree(tt.args.label); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("pathInZigZagTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
