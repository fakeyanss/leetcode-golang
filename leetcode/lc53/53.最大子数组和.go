/*
 * @lc app=leetcode.cn id=53 lang=golang
 * @lcpr version=20004
 *
 * [53] 最大子数组和
 *
 * https://leetcode.cn/problems/maximum-subarray/description/
 *
 * algorithms
 * Medium (55.74%)
 * Likes:    6870
 * Dislikes: 0
 * Total Accepted:    1.9M
 * Total Submissions: 3.5M
 * Testcase Example:  '[-2,1,-3,4,-1,2,1,-5,4]'
 *
 * 给你一个整数数组 nums ，请你找出一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
 *
 * 子数组 是数组中的一个连续部分。
 *
 *
 *
 * 示例 1：
 *
 * 输入：nums = [-2,1,-3,4,-1,2,1,-5,4]
 * 输出：6
 * 解释：连续子数组 [4,-1,2,1] 的和最大，为 6 。
 *
 *
 * 示例 2：
 *
 * 输入：nums = [1]
 * 输出：1
 *
 *
 * 示例 3：
 *
 * 输入：nums = [5,4,-1,7,8]
 * 输出：23
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= nums.length <= 10^5
 * -10^4 <= nums[i] <= 10^4
 *
 *
 *
 *
 * 进阶：如果你已经实现复杂度为 O(n) 的解法，尝试使用更为精妙的 分治法 求解。
 *
 */
package lc0053

import "math"

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func maxSubArray(nums []int) int {
	// // o(n)解法，双指针遍历，r向右移动到末尾后，再l向右移动到末尾

	// // 优化：dp(n)=max(dp(n-1)+nums[n], nums[n]), dp(n)为以nums[i]为结尾的「最大子数组和」
	// dp := make([]int, len(nums))
	// dp[0] = nums[0]
	// for i := 1; i < len(nums); i++ {
	// 	dp[i] = max(nums[i], dp[i-1]+nums[i])
	// }
	// res := math.MinInt
	// for i := 0; i < len(nums); i++ {
	// 	res = max(res, dp[i])
	// }
	// return res

	// 优化：空间复杂度o(1)
	curSum, maxSum := 0, math.MinInt
	for _, v := range nums {
		curSum = max(v, curSum+v)
		maxSum = max(maxSum, curSum)
	}
	return maxSum
}

// @lc code=end

/*
// @lcpr case=start
// [-2,1,-3,4,-1,2,1,-5,4]\n
// @lcpr case=end

// @lcpr case=start
// [1]\n
// @lcpr case=end

// @lcpr case=start
// [5,4,-1,7,8]\n
// @lcpr case=end

*/
