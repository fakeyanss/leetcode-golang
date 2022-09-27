/*
 * @lc app=leetcode.cn id=312 lang=golang
 *
 * [312] 戳气球
 *
 * https://leetcode.cn/problems/burst-balloons/description/
 *
 * algorithms
 * Hard (69.77%)
 * Likes:    1109
 * Dislikes: 0
 * Total Accepted:    89.1K
 * Total Submissions: 127.7K
 * Testcase Example:  '[3,1,5,8]'
 *
 * 有 n 个气球，编号为0 到 n - 1，每个气球上都标有一个数字，这些数字存在数组 nums 中。
 *
 * 现在要求你戳破所有的气球。戳破第 i 个气球，你可以获得 nums[i - 1] * nums[i] * nums[i + 1] 枚硬币。 这里的 i
 * - 1 和 i + 1 代表和 i 相邻的两个气球的序号。如果 i - 1或 i + 1 超出了数组的边界，那么就当它是一个数字为 1 的气球。
 *
 * 求所能获得硬币的最大数量。
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [3,1,5,8]
 * 输出：167
 * 解释：
 * nums = [3,1,5,8] --> [3,5,8] --> [3,8] --> [8] --> []
 * coins =  3*1*5    +   3*5*8   +  1*3*8  + 1*8*1 = 167
 *
 * 示例 2：
 *
 *
 * 输入：nums = [1,5]
 * 输出：10
 *
 *
 *
 *
 * 提示：
 *
 *
 * n == nums.length
 * 1 <= n <= 300
 * 0 <= nums[i] <= 100
 *
 *
 */
package lc312

// @lc code=start
func maxCoins(nums []int) int {
	n := len(nums)
	// 给nums前后补充边
	points := make([]int, n+2)
	points[0] = 1
	points[n+1] = 1
	for i := 0; i < n; i++ {
		points[i+1] = nums[i]
	}

	// dp[i][j] = x表示，戳破气球i和气球j之间（开区间，不包括i和j）的所有气球，可以获得的最高分数为x
	// dp[i][j] = dp[i][k] + dp[k][j] + points[i] * points[k] * points[j], i<k<j
	// i == j 时，dp[i][j] = 0
	dp := make([][]int, n+2)
	for i := range dp {
		dp[i] = make([]int, n+2)
	}

	// base case, 即左上-右下对角线为0, 最终目标是右上角的值
	// 从下往上，遍历
	for i := n; i >= 0; i-- {
		// 从左往右
		for j := i + 1; j < n+2; j++ {
			// 最后选择戳破(i,j)区间的哪个气球
			for k := i + 1; k < j; k++ {
				dp[i][j] = maxInt(dp[i][j],
					dp[i][k]+dp[k][j]+points[i]*points[k]*points[j])
			}
		}
	}

	return dp[0][n+1]
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
