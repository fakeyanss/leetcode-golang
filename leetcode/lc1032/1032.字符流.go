/*
* @lc app=leetcode.cn id=1032 lang=golang
*
* [1032] 字符流
*
* https://leetcode.cn/problems/stream-of-characters/description/
*
  - algorithms
  - Hard (47.53%)
  - Likes:    156
  - Dislikes: 0
  - Total Accepted:    17.3K
  - Total Submissions: 30.9K
  - Testcase Example:  '["StreamChecker","query","query","query","query","query","query","query","query","query","query","query","query"]\n' +
    '[[["cd","f","kl"]],["a"],["b"],["c"],["d"],["e"],["f"],["g"],["h"],["i"],["j"],["k"],["l"]]'

*
* 设计一个算法：接收一个字符流，并检查这些字符的后缀是否是字符串数组 words 中的一个字符串。
*
* 例如，words = ["abc", "xyz"] 且字符流中逐个依次加入 4 个字符 'a'、'x'、'y' 和 'z'
* ，你所设计的算法应当可以检测到 "axyz" 的后缀 "xyz" 与 words 中的字符串 "xyz" 匹配。
*
* 按下述要求实现 StreamChecker 类：
*
*
* StreamChecker(String[] words) ：构造函数，用字符串数组 words 初始化数据结构。
* boolean query(char letter)：从字符流中接收一个新字符，如果字符流中的任一非空后缀能匹配 words 中的某一字符串，返回
* true ；否则，返回 false。
*
*
*
*
* 示例：
*
*
* 输入：
* ["StreamChecker", "query", "query", "query", "query", "query", "query",
* "query", "query", "query", "query", "query", "query"]
* [[["cd", "f", "kl"]], ["a"], ["b"], ["c"], ["d"], ["e"], ["f"], ["g"],
* ["h"], ["i"], ["j"], ["k"], ["l"]]
* 输出：
* [null, false, false, false, true, false, true, false, false, false, false,
* false, true]
*
* 解释：
* StreamChecker streamChecker = new StreamChecker(["cd", "f", "kl"]);
* streamChecker.query("a"); // 返回 False
* streamChecker.query("b"); // 返回 False
* streamChecker.query("c"); // 返回n False
* streamChecker.query("d"); // 返回 True ，因为 'cd' 在 words 中
* streamChecker.query("e"); // 返回 False
* streamChecker.query("f"); // 返回 True ，因为 'f' 在 words 中
* streamChecker.query("g"); // 返回 False
* streamChecker.query("h"); // 返回 False
* streamChecker.query("i"); // 返回 False
* streamChecker.query("j"); // 返回 False
* streamChecker.query("k"); // 返回 False
* streamChecker.query("l"); // 返回 True ，因为 'kl' 在 words 中
*
*
*
*
* 提示：
*
*
* 1 <= words.length <= 2000
* 1 <= words[i].length <= 200
* words[i] 由小写英文字母组成
* letter 是一个小写英文字母
* 最多调用查询 4 * 10^4 次
*
*
*/
package lc1032

// @lc code=start
type StreamChecker struct {
    t trie
    s []byte
}

func Constructor(words []string) StreamChecker {
    t := trie{}
    for _, w := range words {
        t.insert(w)
    }
    return StreamChecker{t: t, s: []byte{}}
}

func (sc *StreamChecker) Query(letter byte) bool {
    sc.s = append(sc.s, letter)
    return sc.t.search(sc.s)
}

type trie struct {
    children [26]*trie
    isEnd    bool
}

// 倒序插入
func (t *trie) insert(word string) {
    node := t
    for i := len(word) - 1; i >= 0; i-- {
        idx := word[i] - 'a'
        if node.children[idx] == nil {
            node.children[idx] = &trie{}
        }
        node = node.children[idx]
    }
    node.isEnd = true
}

// 从后往前遍历字符流，对于每个字符 c，我们在前缀树中查找是否存在以 c 为结尾的字符串，如果存在，返回 true，否则返回 false
func (t *trie) search(word []byte) bool {
    node := t
    for i, j := len(word)-1, 0; i >= 0 && j < 201; i, j = i-1, j+1 {
        idx := word[i] - 'a'
        if node.children[idx] == nil {
            return false
        }
        node = node.children[idx]
        if node.isEnd {
            return true
        }
    }
    return false
}

/**
 * Your StreamChecker object will be instantiated and called as such:
 * obj := Constructor(words);
 * param_1 := obj.Query(letter);
 */
// @lc code=end
