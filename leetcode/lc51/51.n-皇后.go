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

import "strings"

// @lc code=start

func solveNQueens(n int) [][]string {
	var res [][]string
	queens := make([]int, n)     // 皇后放在[idx,val], idx即row，val即col
	col := make([]bool, n)       // 记录每列是否可以放置
	diag1 := make([]bool, n*2-1) // 记录左上角方向是否已经放置，同一斜线上的row-col不变
	diag2 := make([]bool, n*2-1) // 记录右上角方向是否已经放置，同一斜线上的row+col不变

	var dfs func(int)
	dfs = func(row int) {
		if row == n {
			board := make([]string, n)
			for i, c := range queens {
				board[i] = strings.Repeat(".", c) + "Q" + strings.Repeat(".", n-1-c)
			}
			res = append(res, board)
		}

		for c, ok := range col {
			rc := row - c + n - 1 // 偏移量n-1，避免负数
			if !ok && !diag1[rc] && !diag2[row+c] {
				// 可以放置皇后
				queens[row] = c // 直接覆盖
				col[c], diag1[rc], diag2[row+c] = true, true, true
				dfs(row + 1)                                          // 深入下一行
				col[c], diag1[rc], diag2[row+c] = false, false, false // 恢复
			}
		}
	}

	dfs(0)
	return res
}

// @lc code=end
