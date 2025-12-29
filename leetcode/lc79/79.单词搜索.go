/*
 * @lc app=leetcode.cn id=79 lang=golang
 * @lcpr version=20004
 *
 * [79] 单词搜索
 *
 * https://leetcode.cn/problems/word-search/description/
 *
 * algorithms
 * Medium (47.97%)
 * Likes:    1916
 * Dislikes: 0
 * Total Accepted:    609.9K
 * Total Submissions: 1.3M
 * Testcase Example:  '[["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]]\n"ABCCED"'
 *
 * 给定一个 m x n 二维字符网格 board 和一个字符串单词 word 。如果 word 存在于网格中，返回 true ；否则，返回 false
 * 。
 *
 * 单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用。
 *
 *
 *
 * 示例 1：
 *
 * 输入：board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word =
 * "ABCCED"
 * 输出：true
 *
 *
 * 示例 2：
 *
 * 输入：board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word =
 * "SEE"
 * 输出：true
 *
 *
 * 示例 3：
 *
 * 输入：board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word =
 * "ABCB"
 * 输出：false
 *
 *
 *
 *
 * 提示：
 *
 *
 * m == board.length
 * n = board[i].length
 * 1 <= m, n <= 6
 * 1 <= word.length <= 15
 * board 和 word 仅由大小写英文字母组成
 *
 *
 *
 *
 * 进阶：你可以使用搜索剪枝的技术来优化解决方案，使其在 board 更大的情况下可以更快解决问题？
 *
 */
package lc79

import "slices"

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func exist(board [][]byte, word string) bool {
	m, n := len(board), len(board[0])
	directions := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

	// 优化，预处理字母计数
	cnt := map[byte]int{}
	for i := range m {
		for j := range n {
			cnt[board[i][j]]++
		}
	}

	// 剪枝1: word字母计数大于board中的字母计数
	w := []byte(word)
	wordCnt := map[byte]int{}
	for _, c := range w {
		wordCnt[c]++
		if wordCnt[c] > cnt[c] {
			return false
		}
	}

	// 剪枝2: word可以从左往右搜索或者从右往左搜索，选字母计数小的一端开始
	if wordCnt[w[0]] > wordCnt[w[len(w)-1]] {
		slices.Reverse(w)
	}

	var dfs func(int, int, int) bool
	dfs = func(i, j, k int) bool {
		if w[k] != board[i][j] {
			return false
		}
		if k == len(w)-1 {
			return true
		}

		board[i][j] = '0' // 访问过标记
		for _, dir := range directions {
			x, y := i+dir[0], j+dir[1]
			if x >= 0 && x < m && y >= 0 && y < n && dfs(x, y, k+1) {
				return true
			}
		}
		board[i][j] = w[k]
		return false
	}

	for i := range m {
		for j := range n {
			if dfs(i, j, 0) {
				return true
			}
		}
	}
	return false
}

// @lc code=end

/*
// @lcpr case=start
// [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]]\n"ABCCED"\n
// @lcpr case=end

// @lcpr case=start
// [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]]\n"SEE"\n
// @lcpr case=end

// @lcpr case=start
// [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]]\n"ABCB"\n
// @lcpr case=end

*/
