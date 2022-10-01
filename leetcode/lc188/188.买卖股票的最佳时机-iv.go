/*
 * @lc app=leetcode.cn id=188 lang=golang
 *
 * [188] 买卖股票的最佳时机 IV
 *
 * https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-iv/description/
 *
 * algorithms
 * Hard (42.65%)
 * Likes:    815
 * Dislikes: 0
 * Total Accepted:    153.7K
 * Total Submissions: 360.3K
 * Testcase Example:  '2\n[2,4,1]'
 *
 * 给定一个整数数组 prices ，它的第 i 个元素 prices[i] 是一支给定的股票在第 i 天的价格。
 *
 * 设计一个算法来计算你所能获取的最大利润。你最多可以完成 k 笔交易。
 *
 * 注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：k = 2, prices = [2,4,1]
 * 输出：2
 * 解释：在第 1 天 (股票价格 = 2) 的时候买入，在第 2 天 (股票价格 = 4) 的时候卖出，这笔交易所能获得利润 = 4-2 = 2 。
 *
 * 示例 2：
 *
 *
 * 输入：k = 2, prices = [3,2,6,5,0,3]
 * 输出：7
 * 解释：在第 2 天 (股票价格 = 2) 的时候买入，在第 3 天 (股票价格 = 6) 的时候卖出, 这笔交易所能获得利润 = 6-2 = 4 。
 * ⁠    随后，在第 5 天 (股票价格 = 0) 的时候买入，在第 6 天 (股票价格 = 3) 的时候卖出, 这笔交易所能获得利润 = 3-0 =
 * 3 。
 *
 *
 *
 * 提示：
 *
 *
 * 0
 * 0
 * 0
 *
 *
 */
package lc188

import "math"

// @lc code=start
func maxProfit(k int, prices []int) int {
	return dpWithK(k, prices)
}

// dp[i][k][0 or 1]
// 0 <= i <= n-1, 1 <= k <= K
// n 为天数，K 为交易数的上限，0 和 1 代表是否持有股票
// 此问题共有 n*K*2 种状态

// base case：
// dp[-1][...][0] = dp[...][0][0] = 0
// dp[-1][...][1] = dp[...][0][1] = -infifity

// 状态转移
// dp[i][k][0] = max(dp[i-1][k][0], dp[i-1][k][1] + prices[i])
// dp[i][k][1] = max(dp[i-1][k][1], dp[i-1][k-1][0] - prices[i])

func dpWithK(k int, prices []int) int {
	n := len(prices)
	// 确定k的情况下，如果k>n/2，即无论如何都用不完k的额度，等于不限制k
	if k > n/2 {
		return dpWithInfinityK(prices)
	}
	dp := make([][][]int, n)
	for i := range dp {
		dp[i] = make([][]int, k+1)
		for j := range dp[i] {
			dp[i][j] = make([]int, 2)
			// k = 0 的base case
			dp[i][0][0] = 0
			dp[i][0][1] = math.MinInt
		}
	}

	for i := 0; i < n; i++ {
		for j := k; j > 0; j-- {
			// base case
			if i-1 == -1 {
				dp[i][j][0] = 0
				dp[i][j][1] = -prices[i]
				continue
			}
			dp[i][j][0] = maxInt(dp[i-1][j][0], dp[i-1][j][1]+prices[i])
			dp[i][j][1] = maxInt(dp[i-1][j][1], dp[i-1][j-1][0]-prices[i])
		}
	}
	return dp[n-1][k][0]
}

func dpWithInfinityK(prices []int) int {
	// 相当于k正无穷，k和k-1相同
	n := len(prices)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, 2)
	}
	for i := 0; i < n; i++ {
		if i-1 == -1 {
			dp[i][0] = 0
			dp[i][1] = -prices[i]
			continue
		}
		dp[i][0] = maxInt(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = maxInt(dp[i-1][1], dp[i-1][0]-prices[i])
	}
	return dp[n-1][0]
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
