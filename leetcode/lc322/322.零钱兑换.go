/*
 * @lc app=leetcode.cn id=322 lang=golang
 *
 * [322] 零钱兑换
 *
 * https://leetcode.cn/problems/coin-change/description/
 *
 * algorithms
 * Medium (46.00%)
 * Likes:    2121
 * Dislikes: 0
 * Total Accepted:    519.2K
 * Total Submissions: 1.1M
 * Testcase Example:  '[1,2,5]\n11'
 *
 * 给你一个整数数组 coins ，表示不同面额的硬币；以及一个整数 amount ，表示总金额。
 *
 * 计算并返回可以凑成总金额所需的 最少的硬币个数 。如果没有任何一种硬币组合能组成总金额，返回 -1 。
 *
 * 你可以认为每种硬币的数量是无限的。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：coins = [1, 2, 5], amount = 11
 * 输出：3
 * 解释：11 = 5 + 5 + 1
 *
 * 示例 2：
 *
 *
 * 输入：coins = [2], amount = 3
 * 输出：-1
 *
 * 示例 3：
 *
 *
 * 输入：coins = [1], amount = 0
 * 输出：0
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= coins.length <= 12
 * 1 <= coins[i] <= 2^31 - 1
 * 0 <= amount <= 10^4
 *
 *
 */
package lc322

import "sort"

// @lc code=start
func coinChange(coins []int, amount int) int {
	if amount < 1 {
		return 0
	}
	// dp 备忘录, 备忘录长度大于最大硬币个数
	// memo := make([]int, amount+1)
	// for i := 0; i <= amount; i++ {
	// 	memo[i] = -2
	// }

	// var dp func(amount int) int
	// dp = func(amount int) int {
	// 	if amount == 0 {
	// 		return 0
	// 	}
	// 	if amount < 0 {
	// 		return -1
	// 	}
	// 	if memo[amount] != -2 {
	// 		return memo[amount]
	// 	}

	// 	res := math.MaxInt
	// 	for _, v := range coins {
	// 		sub := dp(amount - v)
	// 		if sub >= 0 && res > sub {
	// 			res = sub + 1
	// 		}
	// 	}

	// 	if res == math.MaxInt {
	// 		memo[amount] = -1
	// 	} else {
	// 		memo[amount] = res
	// 	}
	// 	return memo[amount]
	// }

	// return dp(amount)

	// dp table
	// 定义：要凑出金额 n，至少要 dp[n] 个硬币
	dp := make([]int, amount+1)
	dp[0] = 0
	sort.Ints(coins)
	// 计算所有取值
	for i := 0; i <= amount; i++ {
		if i != 0 {
			dp[i] = amount + 1
		}
		for _, v := range coins {
			// 子问题无解，跳过
			if i-v < 0 {
				continue
			}
			if dp[i] > dp[i-v] {
				dp[i] = 1 + dp[i-v]
			}
		}
	}
	if dp[amount] == amount+1 {
		return -1
	}
	return dp[amount]
}

// @lc code=end
