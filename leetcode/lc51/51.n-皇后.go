/*
 * @lc app=leetcode.cn id=51 lang=golang
 *
 * [51] N 皇后
 *
 * https://leetcode.cn/problems/n-queens/description/
 *
 * algorithms
 * Hard (74.12%)
 * Likes:    1511
 * Dislikes: 0
 * Total Accepted:    252.8K
 * Total Submissions: 341K
 * Testcase Example:  '4'
 *
 * 按照国际象棋的规则，皇后可以攻击与之处在同一行或同一列或同一斜线上的棋子。
 *
 * n 皇后问题 研究的是如何将 n 个皇后放置在 n×n 的棋盘上，并且使皇后彼此之间不能相互攻击。
 *
 * 给你一个整数 n ，返回所有不同的 n 皇后问题 的解决方案。
 *
 *
 *
 * 每一种解法包含一个不同的 n 皇后问题 的棋子放置方案，该方案中 'Q' 和 '.' 分别代表了皇后和空位。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：n = 4
 * 输出：[[".Q..","...Q","Q...","..Q."],["..Q.","Q...","...Q",".Q.."]]
 * 解释：如上图所示，4 皇后问题存在两个不同的解法。
 *
 *
 * 示例 2：
 *
 *
 * 输入：n = 1
 * 输出：[["Q"]]
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
package lc51

// @lc code=start
var res [][]string

func solveNQueens(n int) [][]string {
	res = make([][]string, 0)

	// 初始化棋盘
	board := make([][]byte, n)
	for i := range board {
		board[i] = make([]byte, n)
		for j := range board[i] {
			board[i][j] = '.'
		}
	}
	backtrack(board, 0)
	return res
}

// 路径：board 中小于 row 的行，都已经成功放置了皇后
// 选择列表：第 row  行的所有列都是放置皇后的选择
// 结束条件：row 超过 board 的最后一行
func backtrack(board [][]byte, row int) {
	if row == len(board) {
		b := make([]string, 0)
		for i := range board {
			b = append(b, string(board[i]))
		}
		res = append(res, b)
		return
	}

	n := len(board[row])
	for col := 0; col < n; col++ {
		if !isValid(board, row, col) {
			continue
		}

		// 做选择
		board[row][col] = 'Q'
		// 下一行
		backtrack(board, row+1)
		// 撤销选择
		board[row][col] = '.'
	}
}

// 是否可以在 board[i][j] 放置皇后
func isValid(board [][]byte, row, col int) bool {
	n := len(board)
	// 检查列是否有皇后
	for i := 0; i <= row; i++ {
		if board[i][col] == 'Q' {
			return false
		}
	}
	// 检查右上方是否有皇后
	for i, j := row-1, col+1; i >= 0 && j < n; i-- {
		if board[i][j] == 'Q' {
			return false
		}
		j++
	}
	// 检查左上方是否有皇后
	for i, j := row-1, col-1; i >= 0 && j >= 0; i-- {
		if board[i][j] == 'Q' {
			return false
		}
		j--
	}
	return true
}

// @lc code=end
