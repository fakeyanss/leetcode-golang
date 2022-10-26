package helper

import (
	"reflect"
	"testing"
)

func TestNewListNode(t *testing.T) {
	type args struct {
		val  int
		next *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{"case1", args{1, nil}, &ListNode{1, nil}},
		{"case2", args{1, &ListNode{2, nil}}, &ListNode{1, &ListNode{2, nil}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewListNode(tt.args.val, tt.args.next); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("helper.NewListNode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewTreeNode(t *testing.T) {
	type args struct {
		val int
		l   *TreeNode
		r   *TreeNode
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{"case1", args{1, &TreeNode{2, nil, nil}, &TreeNode{3, nil, nil}}, &TreeNode{1, &TreeNode{2, nil, nil}, &TreeNode{3, nil, nil}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewTreeNode(tt.args.val, tt.args.l, tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("helper.NewTreeNode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTrie(t *testing.T) {
	words := make(map[string]int)
	words["them"] = 1
	words["zip"] = 2
	words["team"] = 3
	words["the"] = 4
	words["app"] = 5
	words["that"] = 6
	trieMap := new(TrieMap)
	for k, v := range words {
		trieMap.Put(k, v)
	}

	size := trieMap.Size
	if size != 6 {
		t.Errorf("helper.TrieMap.Size = %v, want %v", size, 6)
	}

	for k := range words {
		exist := trieMap.ContainsKey(k)
		if !exist {
			t.Errorf("helper.TrieMap.ContainsKey() = %v, want %v", exist, true)
		}
	}

	trieMap.Remove("them")
	exist := trieMap.ContainsKey("them")
	if exist {
		t.Errorf("helper.TrieMap.ContainsKey() = %v after remove, want %v", exist, false)
	}
	trieMap.Put("them", 1) // recover

	val := trieMap.Get("them")
	if val != 1 {
		t.Errorf("helper.TrieMap.Get() = %v, want %v", val, 1)
	}

	has := trieMap.HasKeyWithPrefix("th")
	if !has {
		t.Errorf("helper.TrieMap.HasKeyWithPrefix() = %v, want %v", has, true)
	}

	shortest := trieMap.ShortestPrefixOf("themxyz")
	if shortest != "the" {
		t.Errorf("helper.TrieMap.ShortestPrefixOf() = %v, want %v", shortest, "t")
	}

	longest := trieMap.LongestPrefixOf("themxyz")
	if longest != "them" {
		t.Errorf("helper.TrieMap.LongestPrefixOf() = %v, want %v", longest, "t")
	}

	keys := trieMap.KeysWithPrefix("th")
	if !reflect.DeepEqual(keys, []string{"that", "the", "them"}) {
		t.Errorf("helper.TrieMap.KeysWithPrefix() = %v, want %v", keys, []string{"them", "the", "that"})
	}

	keys = trieMap.KeysWithPattern("t.a.")
	if !reflect.DeepEqual(keys, []string{"team", "that"}) {
		t.Errorf("helper.TrieMap.KeysWithPattern() = %v, want %v", keys, []string{"team", "that"})
	}

	has = trieMap.HasKeyWithPattern(".ip")
	if !has {
		t.Errorf("helper.TrieMap.HasKeyWithPattern() = %v, want %v", has, true)
	}
	has = trieMap.HasKeyWithPattern(".i")
	if has {
		t.Errorf("helper.TrieMap.HasKeyWithPattern() = %v, want %v", has, false)
	}
}
