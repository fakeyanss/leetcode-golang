/*
 * @lc app=leetcode.cn id=122 lang=golang
 *
 * [122] 买卖股票的最佳时机 II
 *
 * https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-ii/description/
 *
 * algorithms
 * Medium (71.16%)
 * Likes:    1846
 * Dislikes: 0
 * Total Accepted:    740.6K
 * Total Submissions: 1M
 * Testcase Example:  '[7,1,5,3,6,4]'
 *
 * 给你一个整数数组 prices ，其中 prices[i] 表示某支股票第 i 天的价格。
 *
 * 在每一天，你可以决定是否购买和/或出售股票。你在任何时候 最多 只能持有 一股 股票。你也可以先购买，然后在 同一天 出售。
 *
 * 返回 你能获得的 最大 利润 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：prices = [7,1,5,3,6,4]
 * 输出：7
 * 解释：在第 2 天（股票价格 = 1）的时候买入，在第 3 天（股票价格 = 5）的时候卖出, 这笔交易所能获得利润 = 5 - 1 = 4
 * 。
 * 随后，在第 4 天（股票价格 = 3）的时候买入，在第 5 天（股票价格 = 6）的时候卖出, 这笔交易所能获得利润 = 6 - 3 = 3 。
 * ⁠    总利润为 4 + 3 = 7 。
 *
 * 示例 2：
 *
 *
 * 输入：prices = [1,2,3,4,5]
 * 输出：4
 * 解释：在第 1 天（股票价格 = 1）的时候买入，在第 5 天 （股票价格 = 5）的时候卖出, 这笔交易所能获得利润 = 5 - 1 = 4
 * 。
 * 总利润为 4 。
 *
 * 示例 3：
 *
 *
 * 输入：prices = [7,6,4,3,1]
 * 输出：0
 * 解释：在这种情况下, 交易无法获得正利润，所以不参与交易可以获得最大利润，最大利润为 0 。
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= prices.length <= 3 * 10^4
 * 0 <= prices[i] <= 10^4
 *
 *
 */
package lc122

// @lc code=start
func maxProfit(prices []int) int {
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
