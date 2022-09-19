/*
 * @lc app=leetcode.cn id=494 lang=golang
 *
 * [494] 目标和
 *
 * https://leetcode.cn/problems/target-sum/description/
 *
 * algorithms
 * Medium (49.08%)
 * Likes:    1377
 * Dislikes: 0
 * Total Accepted:    278.9K
 * Total Submissions: 568.4K
 * Testcase Example:  '[1,1,1,1,1]\n3'
 *
 * 给你一个整数数组 nums 和一个整数 target 。
 *
 * 向数组中的每个整数前添加 '+' 或 '-' ，然后串联起所有整数，可以构造一个 表达式 ：
 *
 *
 * 例如，nums = [2, 1] ，可以在 2 之前添加 '+' ，在 1 之前添加 '-' ，然后串联起来得到表达式 "+2-1" 。
 *
 *
 * 返回可以通过上述方法构造的、运算结果等于 target 的不同 表达式 的数目。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [1,1,1,1,1], target = 3
 * 输出：5
 * 解释：一共有 5 种方法让最终目标和为 3 。
 * -1 + 1 + 1 + 1 + 1 = 3
 * +1 - 1 + 1 + 1 + 1 = 3
 * +1 + 1 - 1 + 1 + 1 = 3
 * +1 + 1 + 1 - 1 + 1 = 3
 * +1 + 1 + 1 + 1 - 1 = 3
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums = [1], target = 1
 * 输出：1
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1
 * 0
 * 0
 * -1000
 *
 *
 */
package lc494

import "math"

// @lc code=start
func findTargetSumWays(nums []int, target int) int {
	// return backtrack(nums, target)
	return dp(nums, target)
}

// 回溯，暴力穷举
func backtrack(nums []int, target int) int {
	count := 0
	var f func(nums []int, id, remain int)
	f = func(nums []int, id int, remain int) {
		// base case
		if id == len(nums) {
			if remain == 0 {
				// 恰好凑出 target
				count++
			}
			return
		}

		// 选择-号
		remain += nums[id]
		// 穷举 nums[i + 1]
		f(nums, id+1, remain)
		// 撤销选择
		remain -= nums[id]

		// 选择+号
		remain -= nums[id]
		// 穷举 nums[i + 1]
		f(nums, id+1, remain)
		// 撤销选择
		remain += nums[id]
	}

	f(nums, 0, target)
	return count
}

// 如果我们把 nums 划分成两个子集 A 和 B，分别代表分配 + 的数和分配 - 的数，那么他们和 target 存在如下关系：
// sum(A) - sum(B) = target
// sum(A) = target + sum(B)
// sum(A) + sum(A) = target + sum(B) + sum(A)
// 2 * sum(A) = target + sum(nums)
// 可以推出 sum(A) = (target + sum(nums)) / 2，也就是把原问题转化成：nums 中存在几个子集 A，使得 A 中元素的和为 (target + sum(nums)) / 2
func dp(nums []int, target int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	if sum < int(math.Abs(float64(target))) || (sum+target)%2 == 1 {
		return 0
	}
	sum = (sum + target) / 2

	// dp[i][j] = x 表示，若只在前 i 个物品中选择，若当前背包的容量为 j，则最多有 x 种方法可以恰好装满背包
	n := len(nums)
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, sum+1)
	}
	// base case
	dp[0][0] = 1

	for i := 1; i <= n; i++ {
		for j := 0; j <= sum; j++ {
			if j-nums[i-1] < 0 {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j] + dp[i-1][j-nums[i-1]]
			}
		}
	}
	return dp[n][sum]
}

// @lc code=end
