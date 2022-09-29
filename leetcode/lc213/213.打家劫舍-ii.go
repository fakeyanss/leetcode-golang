/*
 * @lc app=leetcode.cn id=213 lang=golang
 *
 * [213] 打家劫舍 II
 *
 * https://leetcode.cn/problems/house-robber-ii/description/
 *
 * algorithms
 * Medium (43.92%)
 * Likes:    1167
 * Dislikes: 0
 * Total Accepted:    286.6K
 * Total Submissions: 652.4K
 * Testcase Example:  '[2,3,2]'
 *
 * 你是一个专业的小偷，计划偷窃沿街的房屋，每间房内都藏有一定的现金。这个地方所有的房屋都 围成一圈
 * ，这意味着第一个房屋和最后一个房屋是紧挨着的。同时，相邻的房屋装有相互连通的防盗系统，如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警 。
 *
 * 给定一个代表每个房屋存放金额的非负整数数组，计算你 在不触动警报装置的情况下 ，今晚能够偷窃到的最高金额。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [2,3,2]
 * 输出：3
 * 解释：你不能先偷窃 1 号房屋（金额 = 2），然后偷窃 3 号房屋（金额 = 2）, 因为他们是相邻的。
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums = [1,2,3,1]
 * 输出：4
 * 解释：你可以先偷窃 1 号房屋（金额 = 1），然后偷窃 3 号房屋（金额 = 3）。
 * 偷窃到的最高金额 = 1 + 3 = 4 。
 *
 * 示例 3：
 *
 *
 * 输入：nums = [1,2,3]
 * 输出：3
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= nums.length <= 100
 * 0 <= nums[i] <= 1000
 *
 *
 */
package lc213

// @lc code=start
func rob(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}
	return maxInt(robRange(nums, 0, n-2), robRange(nums, 1, n-1))
}

func robRange(nums []int, start, end int) int {
	n := len(nums)
	// dp[i]=x 表示 抢劫nums[i..end]的房子，获取的最大金额为x
	dp := make([]int, n+2)
	for i := end; i >= start; i-- {
		dp[i] = maxInt(dp[i+1], nums[i]+dp[i+2])
	}
	return dp[start]
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
