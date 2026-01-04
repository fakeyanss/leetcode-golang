/*
* @lc app=leetcode.cn id=212 lang=golang
* @lcpr version=30305
*
* [212] 单词搜索 II
*
* https://leetcode.cn/problems/word-search-ii/description/
*
  - algorithms
  - Hard (43.46%)
  - Likes:    970
  - Dislikes: 0
  - Total Accepted:    136.9K
  - Total Submissions: 314.8K
  - Testcase Example:  '[["o","a","a","n"],["e","t","a","e"],["i","h","k","r"],["i","f","l","v"]]\n' +
    '["oath","pea","eat","rain"]\n' +
    '[["a","b"],["c","d"]]\n' +
    '["abcb"]'

*
* 给定一个 m x n 二维字符网格 board 和一个单词（字符串）列表 words， 返回所有二维网格上的单词 。
*
* 单词必须按照字母顺序，通过 相邻的单元格
* 内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母在一个单词中不允许被重复使用。
*
*
*
* 示例 1：
*
* 输入：board =
* [["o","a","a","n"],["e","t","a","e"],["i","h","k","r"],["i","f","l","v"]],
* words = ["oath","pea","eat","rain"]
* 输出：["eat","oath"]
*
*
* 示例 2：
*
* 输入：board = [["a","b"],["c","d"]], words = ["abcb"]
* 输出：[]
*
*
*
*
* 提示：
*
*
* m == board.length
* n == board[i].length
* 1 <= m, n <= 12
* board[i][j] 是一个小写英文字母
* 1 <= words.length <= 3 * 10^4
* 1 <= words[i].length <= 10
* words[i] 由小写英文字母组成
* words 中的所有字符串互不相同
*
*
*/
package lc212

// @lc code=start
func findWords(board [][]byte, words []string) []string {
    root := &TrieNode{}
    for _, w := range words {
        root.insert(w)
    }

    m, n := len(board), len(board[0])
    var res []string
    directions := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

    // dfs: 从(i,j)出发，在trie的node节点下匹配单词
    var dfs func(i, j int, node *TrieNode)
    dfs = func(i, j int, node *TrieNode) {
        if i < 0 || i >= m || j < 0 || j >= n || board[i][j] == '#' {
            return
        }
        idx := board[i][j] - 'a'
        if node.son[idx] == nil {
            return
        }

        cur := node.son[idx]
        // 找到单词，加入结果集
        if cur.word != "" {
            res = append(res, cur.word)
            cur.word = "" // 剪枝，删除该单词，避免后面重复找到
        }

        v := board[i][j]
        board[i][j] = '#' // 访问标记，直接原地修改
        for _, dir := range directions {
            x, y := i+dir[0], j+dir[1]
            dfs(x, y, cur)
        }
        board[i][j] = v // 恢复标记
    }

    for i := range m {
        for j := range n {
            dfs(i, j, root)
        }
    }
    return res
}

type TrieNode struct {
    son  [26]*TrieNode
    word string // 从end标记改为存储完整路径，方便匹配时取出结果
}

func (tn *TrieNode) insert(word string) {
    cur := tn
    for _, c := range word {
        c -= 'a'
        if cur.son[c] == nil {
            cur.son[c] = &TrieNode{}
        }
        cur = cur.son[c]
    }
    cur.word = word
}

// @lc code=end

/*
// @lcpr case=start
// [["o","a","a","n"],["e","t","a","e"],["i","h","k","r"],["i","f","l","v"]]\n["oath","pea","eat","rain"]\n
// @lcpr case=end

// @lcpr case=start
// [["a","b"],["c","d"]]\n["abcb"]\n
// @lcpr case=end

*/
