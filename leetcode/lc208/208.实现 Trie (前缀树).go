/*
* @lc app=leetcode.cn id=208 lang=golang
* @lcpr version=20003
*
* [208] 实现 Trie (前缀树)
*
* https://leetcode.cn/problems/implement-trie-prefix-tree/description/
*
  - algorithms
  - Medium (72.37%)
  - Likes:    1724
  - Dislikes: 0
  - Total Accepted:    378.7K
  - Total Submissions: 523.1K
  - Testcase Example:  '["Trie","insert","search","search","startsWith","insert","search"]\n' +
    '[[],["apple"],["apple"],["app"],["app"],["app"],["app"]]'

*
* Trie（发音类似 "try"）或者说 前缀树
* 是一种树形数据结构，用于高效地存储和检索字符串数据集中的键。这一数据结构有相当多的应用情景，例如自动补全和拼写检查。
*
* 请你实现 Trie 类：
*
*
* Trie() 初始化前缀树对象。
* void insert(String word) 向前缀树中插入字符串 word 。
* boolean search(String word) 如果字符串 word 在前缀树中，返回 true（即，在检索之前已经插入）；否则，返回
* false 。
* boolean startsWith(String prefix) 如果之前已经插入的字符串 word 的前缀之一为 prefix ，返回 true
* ；否则，返回 false 。
*
*
*
*
* 示例：
*
* 输入
* ["Trie", "insert", "search", "search", "startsWith", "insert", "search"]
* [[], ["apple"], ["apple"], ["app"], ["app"], ["app"], ["app"]]
* 输出
* [null, null, true, false, true, null, true]
*
* 解释
* Trie trie = new Trie();
* trie.insert("apple");
* trie.search("apple");   // 返回 True
* trie.search("app");     // 返回 False
* trie.startsWith("app"); // 返回 True
* trie.insert("app");
* trie.search("app");     // 返回 True
*
*
*
*
* 提示：
*
*
* 1 <= word.length, prefix.length <= 2000
* word 和 prefix 仅由小写英文字母组成
* insert、search 和 startsWith 调用次数 总计 不超过 3 * 10^4 次
*
*
*/
package lc208

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
type Trie struct {
	Children [26]*Trie
	Val      *struct{}
}

func Constructor() Trie {
	return Trie{}
}

func (t *Trie) getNode(node *Trie, key string) *Trie {
	p, n := node, len(key)
	for i := 0; i < n; i++ {
		if p == nil {
			return nil
		}
		c := key[i] - 'a'
		p = p.Children[c]
	}
	return p
}

func (t *Trie) Insert(word string) {
	t.insert(t, word, 0)
}

func (t *Trie) insert(node *Trie, key string, i int) *Trie {
	if node == nil {
		node = new(Trie)
	}
	if i == len(key) {
		node.Val = &struct{}{}
		return node
	}
	c := key[i] - 'a'
	node.Children[c] = t.insert(node.Children[c], key, i+1)
	return node
}

func (t *Trie) Search(word string) bool {
	x := t.getNode(t, word)
	if x == nil || x.Val == nil {
		return false
	}
	return true
}

func (t *Trie) StartsWith(prefix string) bool {
	return t.getNode(t, prefix) != nil
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
// @lc code=end
