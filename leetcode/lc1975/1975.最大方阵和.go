/*
 * @lc app=leetcode.cn id=1975 lang=golang
 * @lcpr version=30305
 *
 * [1975] 最大方阵和
 *
 * https://leetcode.cn/problems/maximum-matrix-sum/description/
 *
 * algorithms
 * Medium (43.90%)
 * Likes:    36
 * Dislikes: 0
 * Total Accepted:    9.8K
 * Total Submissions: 20.3K
 * Testcase Example:  '[[1,-1],[-1,1]]\n[[1,2,3],[-1,-2,-3],[1,2,3]]'
 *
 * 给你一个 n x n 的整数方阵 matrix 。你可以执行以下操作 任意次 ：
 *
 *
 * 选择 matrix 中 相邻 两个元素，并将它们都 乘以 -1 。
 *
 *
 * 如果两个元素有 公共边 ，那么它们就是 相邻 的。
 *
 * 你的目的是 最大化 方阵元素的和。请你在执行以上操作之后，返回方阵的 最大 和。
 *
 *
 *
 * 示例 1：
 *
 * 输入：matrix = [[1,-1],[-1,1]]
 * 输出：4
 * 解释：我们可以执行以下操作使和等于 4 ：
 * - 将第一行的 2 个元素乘以 -1 。
 * - 将第一列的 2 个元素乘以 -1 。
 *
 *
 * 示例 2：
 *
 * 输入：matrix = [[1,2,3],[-1,-2,-3],[1,2,3]]
 * 输出：16
 * 解释：我们可以执行以下操作使和等于 16 ：
 * - 将第二行的最后 2 个元素乘以 -1 。
 *
 *
 *
 *
 * 提示：
 *
 *
 * n == matrix.length == matrix[i].length
 * 2 <= n <= 250
 * -10^5 <= matrix[i][j] <= 10^5
 *
 *
 */
package lc1975

import "math"

// @lc code=start
func maxMatrixSum(matrix [][]int) int64 {
	// 任意有两个负数，可以通过连线路径上的多次转换，将这两个负数变成正数。
	// 所以当matrix有偶数个负数时，变换最大总和为所有元素的绝对值之和
	// 当matrix有奇数个负数时，丢掉绝对值最小的负数，结果为其余元素绝对值之和
	total, negCnt, minN := 0, 0, math.MaxInt
	for _, row := range matrix {
		for _, x := range row {
			if x < 0 {
				negCnt++
				x = -x
			}
			minN = min(minN, x)
			total += x
		}
	}
	if negCnt%2 == 1 {
		total -= minN * 2
	}
	return int64(total)
}

// @lc code=end

/*
// @lcpr case=start
// [[1,-1],[-1,1]]\n
// @lcpr case=end

// @lcpr case=start
// [[1,2,3],[-1,-2,-3],[1,2,3]]\n
// @lcpr case=end

*/
