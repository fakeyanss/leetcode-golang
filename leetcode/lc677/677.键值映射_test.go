package lc677

import (
	"testing"
)

func TestMapSum(t *testing.T) {
	trie := Constructor()
	trie.Insert("apple", 3)
	val := trie.Sum("ap")
	if val != 3 {
		t.Errorf("TestMap.Sum = %v, want %v", val, 3)
	}
	trie.Insert("app", 2)
	val = trie.Sum("ap")
	if val != 5 {
		t.Errorf("TestMap.Sum = %v, want %v", val, 5)
	}
}
