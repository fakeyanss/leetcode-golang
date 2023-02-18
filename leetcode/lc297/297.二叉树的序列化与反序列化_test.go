package lc297

import (
	"reflect"
	"testing"

	"github.com/fakeyanss/leetcode-golang/helper"
)

func TestCodec_serialize(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		this *Codec
		args args
		want string
	}{
		{
			"case1",
			&Codec{},
			args{
				helper.NewTreeNode(2,
					helper.NewTreeNode(1, nil, helper.NewTreeNode(6, nil, nil)),
					helper.NewTreeNode(3, nil, nil),
				),
			},
			"2,1,null,6,null,null,3,null,null,",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.this.serialize(tt.args.root); got != tt.want {
				t.Errorf("Codec.serialize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCodec_deserialize(t *testing.T) {
	type args struct {
		data string
	}
	tests := []struct {
		name string
		this *Codec
		args args
		want *TreeNode
	}{
		{
			"case1",
			&Codec{},
			args{"2,1,null,6,null,null,3,null,null,"},
			helper.NewTreeNode(2,
				helper.NewTreeNode(1, nil, helper.NewTreeNode(6, nil, nil)),
				helper.NewTreeNode(3, nil, nil),
			),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &Codec{}
			if got := this.deserialize(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Codec.deserialize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConstructor(t *testing.T) {
	tests := []struct {
		name string
		want Codec
	}{
		{"case1", Codec{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Constructor(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Constructor() = %v, want %v", got, tt.want)
			}
		})
	}
}
