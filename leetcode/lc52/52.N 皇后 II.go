/*
 * @lc app=leetcode.cn id=52 lang=golang
 * @lcpr version=20004
 *
 * [52] N 皇后 II
 *
 * https://leetcode.cn/problems/n-queens-ii/description/
 *
 * algorithms
 * Hard (82.49%)
 * Likes:    556
 * Dislikes: 0
 * Total Accepted:    178.4K
 * Total Submissions: 214.9K
 * Testcase Example:  '4'
 *
 * n 皇后问题 研究的是如何将 n 个皇后放置在 n × n 的棋盘上，并且使皇后彼此之间不能相互攻击。
 *
 * 给你一个整数 n ，返回 n 皇后问题 不同的解决方案的数量。
 *
 *
 *
 *
 *
 * 示例 1：
 *
 * 输入：n = 4
 * 输出：2
 * 解释：如上图所示，4 皇后问题存在两个不同的解法。
 *
 *
 * 示例 2：
 *
 * 输入：n = 1
 * 输出：1
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= n <= 9
 *
 *
 *
 *
 */
package lc52

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func totalNQueens(n int) int {
	board := make([][]int, n)
	for i := range board {
		board[i] = make([]int, n)
	}
	var res int

	isValid := func(board [][]int, row, col int) bool {
		for i := 0; i < n; i++ {
			// if board[row][i] == 1 { // 同行
			// 	return false
			// }
			if board[i][col] == 1 { // 同列
				return false
			}
		}
		for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 { // 左上方斜线
			if board[i][j] == 1 {
				return false
			}
		}
		// for i, j := row+1, col+1; i < n && j < n; i, j = i+1, j+1 { // 右下方斜线
		// 	if board[i][j] == 1 {
		// 		return false
		// 	}
		// }
		for i, j := row-1, col+1; i >= 0 && j < n; i, j = i-1, j+1 { // 右上方斜线
			if board[i][j] == 1 {
				return false
			}
		}
		// for i, j := row+1, col-1; i < n && j >= 0; i, j = i+1, j-1 { // 左下方斜线
		// 	if board[i][j] == 1 {
		// 		return false
		// 	}
		// }
		return true
	}

	var backtrack func(row int)
	backtrack = func(row int) {
		if row == n {
			res++
			return
		}

		for col := 0; col < n; col++ {
			if !isValid(board, row, col) {
				continue
			}
			board[row][col] = 1
			backtrack(row + 1)
			board[row][col] = 0
		}
	}

	backtrack(0)
	return res
}

// @lc code=end

/*
// @lcpr case=start
// 4\n
// @lcpr case=end

// @lcpr case=start
// 1\n
// @lcpr case=end

*/
