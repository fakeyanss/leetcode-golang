/*
 * @lc app=leetcode.cn id=211 lang=golang
 *
 * [211] 添加与搜索单词 - 数据结构设计
 *
 * https://leetcode.cn/problems/design-add-and-search-words-data-structure/description/
 *
 * algorithms
 * Medium (49.88%)
 * Likes:    466
 * Dislikes: 0
 * Total Accepted:    66.6K
 * Total Submissions: 133.5K
 * Testcase Example:  '["WordDictionary","addWord","addWord","addWord","search","search","search","search"]\n' +
  '[[],["bad"],["dad"],["mad"],["pad"],["bad"],[".ad"],["b.."]]'
 *
 * 请你设计一个数据结构，支持 添加新单词 和 查找字符串是否与任何先前添加的字符串匹配 。
 *
 * 实现词典类 WordDictionary ：
 *
 *
 * WordDictionary() 初始化词典对象
 * void addWord(word) 将 word 添加到数据结构中，之后可以对它进行匹配
 * bool search(word) 如果数据结构中存在字符串与 word 匹配，则返回 true ；否则，返回  false 。word 中可能包含一些
 * '.' ，每个 . 都可以表示任何一个字母。
 *
 *
 *
 *
 * 示例：
 *
 *
 * 输入：
 *
 * ["WordDictionary","addWord","addWord","addWord","search","search","search","search"]
 * [[],["bad"],["dad"],["mad"],["pad"],["bad"],[".ad"],["b.."]]
 * 输出：
 * [null,null,null,null,false,true,true,true]
 *
 * 解释：
 * WordDictionary wordDictionary = new WordDictionary();
 * wordDictionary.addWord("bad");
 * wordDictionary.addWord("dad");
 * wordDictionary.addWord("mad");
 * wordDictionary.search("pad"); // 返回 False
 * wordDictionary.search("bad"); // 返回 True
 * wordDictionary.search(".ad"); // 返回 True
 * wordDictionary.search("b.."); // 返回 True
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= word.length <= 25
 * addWord 中的 word 由小写英文字母组成
 * search 中的 word 由 '.' 或小写英文字母组成
 * 最多调用 10^4 次 addWord 和 search
 *
 *
*/
package lc211

// @lc code=start
type WordDictionary struct {
    Children [26]*WordDictionary
    Val      *struct{}
}

func Constructor() WordDictionary {
    return WordDictionary{}
}

func (t *WordDictionary) getNode(node *WordDictionary, key string) *WordDictionary {
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

func (t *WordDictionary) insert(node *WordDictionary, key string, i int) *WordDictionary {
    if node == nil {
        node = new(WordDictionary)
    }
    if i == len(key) {
        node.Val = &struct{}{}
        return node
    }
    c := key[i] - 'a'
    node.Children[c] = t.insert(node.Children[c], key, i+1)
    return node
}

func (t *WordDictionary) hasKeyWithPattern(node *WordDictionary, pattern string, i int) bool {
    if node == nil {
        // 树枝不存在，匹配失败
        return false
    }
    if i == len(pattern) {
        // 模式串匹配完，看最终节点是否是一个k
        return node.Val != nil
    }
    c := pattern[i]
    if c != '.' {
        // 如果不是通配符，遍历下一个指定子节点
        return t.hasKeyWithPattern(node.Children[c-'a'], pattern, i+1)
    }
    // 是通配符，遍历所有子节点
    for j := 0; j < 26; j++ {
        // 任何一个子节点匹配成功，就返回true
        if t.hasKeyWithPattern(node.Children[j], pattern, i+1) {
            return true
        }
    }
    return false
}

func (t *WordDictionary) AddWord(word string) {
    t.insert(t, word, 0)
}

func (t *WordDictionary) Search(word string) bool {
    return t.hasKeyWithPattern(t, word, 0)
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
// @lc code=end
