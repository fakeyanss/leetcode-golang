/*
 * @lc app=leetcode.cn id=121 lang=golang
 *
 * [121] 买卖股票的最佳时机
 *
 * https://leetcode.cn/problems/best-time-to-buy-and-sell-stock/description/
 *
 * algorithms
 * Easy (57.96%)
 * Likes:    2429
 * Dislikes: 0
 * Total Accepted:    837.8K
 * Total Submissions: 1.4M
 * Testcase Example:  '[7,1,5,3,6,4]'
 *
 * 给定一个数组 prices ，它的第 i 个元素 prices[i] 表示一支给定股票第 i 天的价格。
 *
 * 你只能选择 某一天 买入这只股票，并选择在 未来的某一个不同的日子 卖出该股票。设计一个算法来计算你所能获取的最大利润。
 *
 * 返回你可以从这笔交易中获取的最大利润。如果你不能获取任何利润，返回 0 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：[7,1,5,3,6,4]
 * 输出：5
 * 解释：在第 2 天（股票价格 = 1）的时候买入，在第 5 天（股票价格 = 6）的时候卖出，最大利润 = 6-1 = 5 。
 * ⁠    注意利润不能是 7-1 = 6, 因为卖出价格需要大于买入价格；同时，你不能在买入前卖出股票。
 *
 *
 * 示例 2：
 *
 *
 * 输入：prices = [7,6,4,3,1]
 * 输出：0
 * 解释：在这种情况下, 没有交易完成, 所以最大利润为 0。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1
 * 0
 *
 *
 */
package lc121

// @lc code=start
func maxProfit(prices []int) int {
	// min := math.MaxInt
	// max := 0
	// for i := 0; i < len(prices); i++ {
	// 	sub := prices[i] - min
	// 	if min > prices[i] {
	// 		min = prices[i]
	// 	} else if sub > max {
	// 		max = sub
	// 	}
	// }
	// return max

	return dp(prices)
}

// dp[i][k][0 or 1]
// 0 <= i <= n - 1, 1 <= k <= K
// n 为天数，大 K 为交易数的上限，0 和 1 代表是否持有股票。
// 此问题共 n × K × 2 种状态，全部穷举就能搞定。

// base case：
// dp[-1][...][0] = dp[...][0][0] = 0
// dp[-1][...][1] = dp[...][0][1] = -infinity

// 状态转移方程：
// dp[i][k][0] = max(dp[i-1][k][0], dp[i-1][k][1] + prices[i])
// dp[i][k][1] = max(dp[i-1][k][1], dp[i-1][k-1][0] - prices[i])

func dp(prices []int) int {
	// 相当于k=1
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
		dp[i][1] = maxInt(dp[i-1][1], -prices[i])
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
