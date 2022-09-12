/*
 * @lc app=leetcode.cn id=931 lang=golang
 *
 * [931] 下降路径最小和
 *
 * https://leetcode.cn/problems/minimum-falling-path-sum/description/
 *
 * algorithms
 * Medium (67.18%)
 * Likes:    190
 * Dislikes: 0
 * Total Accepted:    46.1K
 * Total Submissions: 68.6K
 * Testcase Example:  '[[2,1,3],[6,5,4],[7,8,9]]'
 *
 * 给你一个 n x n 的 方形 整数数组 matrix ，请你找出并返回通过 matrix 的下降路径 的 最小和 。
 *
 * 下降路径
 * 可以从第一行中的任何元素开始，并从每一行中选择一个元素。在下一行选择的元素和当前行所选元素最多相隔一列（即位于正下方或者沿对角线向左或者向右的第一个元素）。具体来说，位置
 * (row, col) 的下一个元素应当是 (row + 1, col - 1)、(row + 1, col) 或者 (row + 1, col + 1)
 * 。
 *
 *
 *
 * 示例 1：
 *
 *
 *
 *
 * 输入：matrix = [[2,1,3],[6,5,4],[7,8,9]]
 * 输出：13
 * 解释：如图所示，为和最小的两条下降路径
 *
 *
 * 示例 2：
 *
 *
 *
 *
 * 输入：matrix = [[-19,57],[-40,-5]]
 * 输出：-59
 * 解释：如图所示，为和最小的下降路径
 *
 *
 *
 *
 * 提示：
 *
 *
 * n == matrix.length == matrix[i].length
 * 1 <= n <= 100
 * -100 <= matrix[i][j] <= 100
 *
 *
 */
package lc931

import (
	"math"
)

// @lc code=start
var memo [][]int

func minFallingPathSum(matrix [][]int) int {
	n := len(matrix)
	res := math.MaxInt
	memo = make([][]int, n)
	for i := 0; i < n; i++ {
		memo[i] = make([]int, n)
		for j := 0; j < n; j++ {
			memo[i][j] = math.MaxInt
		}
	}

	for j := 0; j < n; j++ {
		// 终点在matrix[n-1]的任意一列
		res = minInt(res, dp(matrix, n-1, j))
	}
	return res
}

// 从第一行（matrix[0][..]）向下落，落到位置 matrix[i][j] 的最小路径和为 dp(matrix, i, j)
func dp(matrix [][]int, i, j int) int {
	// 边界条件
	if i < 0 || j < 0 || i >= len(matrix) || j >= len(matrix) {
		return math.MaxInt
	}
	// base case
	if i == 0 {
		return matrix[0][j]
	}
	// 备忘录查询
	if memo[i][j] != math.MaxInt {
		return memo[i][j]
	}

	memo[i][j] = matrix[i][j] + minInt(dp(matrix, i-1, j-1), dp(matrix, i-1, j), dp(matrix, i-1, j+1))

	return memo[i][j]
}

func minInt(nums ...int) int {
	res := nums[0]
	for i := 1; i < len(nums); i++ {
		if res > nums[i] {
			res = nums[i]
		}
	}
	return res
}

// @lc code=end
