/*
 * @lc app=leetcode.cn id=64 lang=golang
 * @lcpr version=20004
 *
 * [64] 最小路径和
 *
 * https://leetcode.cn/problems/minimum-path-sum/description/
 *
 * algorithms
 * Medium (71.34%)
 * Likes:    1754
 * Dislikes: 0
 * Total Accepted:    695.6K
 * Total Submissions: 975K
 * Testcase Example:  '[[1,3,1],[1,5,1],[4,2,1]]'
 *
 * 给定一个包含非负整数的 m x n 网格 grid ，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。
 *
 * 说明：每次只能向下或者向右移动一步。
 *
 *
 *
 * 示例 1：
 *
 * 输入：grid = [[1,3,1],[1,5,1],[4,2,1]]
 * 输出：7
 * 解释：因为路径 1→3→1→1→1 的总和最小。
 *
 *
 * 示例 2：
 *
 * 输入：grid = [[1,2,3],[4,5,6]]
 * 输出：12
 *
 *
 *
 *
 * 提示：
 *
 *
 * m == grid.length
 * n == grid[i].length
 * 1 <= m, n <= 200
 * 0 <= grid[i][j] <= 200
 *
 *
 */
package lc64

import "math"

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func minPathSum(grid [][]int) int {
	// dp[i][j] 表示从(0,0)到(i,j)的最小路径和
	m, n := len(grid), len(grid[0])
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
		dp[i][0] = math.MaxInt
	}
	for j := range dp[0] {
		dp[0][j] = math.MaxInt
	}
	dp[1][1] = grid[0][0]
	for i := range m {
		for j := range n {
			if i == 0 && j == 0 {
				continue
			}
			dp[i+1][j+1] = grid[i][j] + min(dp[i+1][j], dp[i][j+1])
		}
	}
	return dp[m][n]
}

// @lc code=end

/*
// @lcpr case=start
// [[1,3,1],[1,5,1],[4,2,1]]\n
// @lcpr case=end

// @lcpr case=start
// [[1,2,3],[4,5,6]]\n
// @lcpr case=end

*/
