/*
 * @lc app=leetcode.cn id=1800 lang=golang
 *
 * [1800] 最大升序子数组和
 *
 * https://leetcode.cn/problems/maximum-ascending-subarray-sum/description/
 *
 * algorithms
 * Easy (67.20%)
 * Likes:    79
 * Dislikes: 0
 * Total Accepted:    35.4K
 * Total Submissions: 51.3K
 * Testcase Example:  '[10,20,30,5,10,50]'
 *
 * 给你一个正整数组成的数组 nums ，返回 nums 中一个 升序 子数组的最大可能元素和。
 *
 * 子数组是数组中的一个连续数字序列。
 *
 * 已知子数组 [numsl, numsl+1, ..., numsr-1, numsr] ，若对所有 i（l ），numsi  < numsi+1
 * 都成立，则称这一子数组为 升序 子数组。注意，大小为 1 的子数组也视作 升序 子数组。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [10,20,30,5,10,50]
 * 输出：65
 * 解释：[5,10,50] 是元素和最大的升序子数组，最大元素和为 65 。
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums = [10,20,30,40,50]
 * 输出：150
 * 解释：[10,20,30,40,50] 是元素和最大的升序子数组，最大元素和为 150 。
 *
 *
 * 示例 3：
 *
 *
 * 输入：nums = [12,17,15,13,10,11,12]
 * 输出：33
 * 解释：[10,11,12] 是元素和最大的升序子数组，最大元素和为 33 。
 *
 *
 * 示例 4：
 *
 *
 * 输入：nums = [100,10,1]
 * 输出：100
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1
 * 1
 *
 *
 */
package lc1800

// @lc code=start
func maxAscendingSum(nums []int) int {
	// 简单dp
	// dp[i] = x，表示nums[0..i]中以nums[i]结尾的升序子数组最大和为x
	n := len(nums)
	dp := make([]int, n)
	dp[0] = nums[0]
	res := dp[0]
	for i := 1; i < n; i++ {
		if nums[i] > nums[i-1] {
			dp[i] = dp[i-1] + nums[i]
		} else {
			dp[i] = nums[i]
		}
		res = maxInt(res, dp[i])
	}
	return res
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
