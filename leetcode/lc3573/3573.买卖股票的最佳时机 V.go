/*
 * @lc app=leetcode.cn id=3573 lang=golang
 * @lcpr version=30305
 *
 * [3573] 买卖股票的最佳时机 V
 *
 * https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-v/description/
 *
 * algorithms
 * Medium (52.21%)
 * Likes:    30
 * Dislikes: 0
 * Total Accepted:    9K
 * Total Submissions: 14.6K
 * Testcase Example:  '[1,7,9,8,2]\n2\n[12,16,19,19,8,1,19,13,9]\n3'
 *
 * 给你一个整数数组 prices，其中 prices[i] 是第 i 天股票的价格（美元），以及一个整数 k。
 *
 * 你最多可以进行 k 笔交易，每笔交易可以是以下任一类型：
 *
 *
 *
 * 普通交易：在第 i 天买入，然后在之后的第 j 天卖出，其中 i < j。你的利润是 prices[j] -
 * prices[i]。
 *
 *
 * 做空交易：在第 i 天卖出，然后在之后的第 j 天买回，其中 i < j。你的利润是 prices[i] - prices[j]。
 *
 *
 *
 * 注意：你必须在开始下一笔交易之前完成当前交易。此外，你不能在已经进行买入或卖出操作的同一天再次进行买入或卖出操作。
 *
 * 通过进行 最多 k 笔交易，返回你可以获得的最大总利润。
 *
 *
 *
 * 示例 1:
 *
 *
 * 输入: prices = [1,7,9,8,2], k = 2
 *
 * 输出: 14
 *
 * 解释:
 * 我们可以通过 2 笔交易获得 14 美元的利润：
 *
 *
 * 一笔普通交易：第 0 天以 1 美元买入，第 2 天以 9 美元卖出。
 * 一笔做空交易：第 3 天以 8 美元卖出，第 4 天以 2 美元买回。
 *
 *
 *
 * 示例 2:
 *
 *
 * 输入: prices = [12,16,19,19,8,1,19,13,9], k = 3
 *
 * 输出: 36
 *
 * 解释:
 * 我们可以通过 3 笔交易获得 36 美元的利润：
 *
 *
 * 一笔普通交易：第 0 天以 12 美元买入，第 2 天以 19 美元卖出。
 * 一笔做空交易：第 3 天以 19 美元卖出，第 4 天以 8 美元买回。
 * 一笔普通交易：第 5 天以 1 美元买入，第 6 天以 19 美元卖出。
 *
 *
 *
 *
 *
 * 提示:
 *
 *
 * 2 <= prices.length <= 10^3
 * 1 <= prices[i] <= 10^9
 * 1 <= k <= prices.length / 2
 *
 *
 */
package lc3573

// @lc code=start
func maximumProfit(prices []int, k int) int64 {
	// dp[i][j][0|1|2], 表示第i天最大j次交易，当前(0不持有股票|1持有股票|2卖出做空)状态下的最大利润
	n := len(prices)
	dp := make([][][]int64, n)
	for i := range dp {
		dp[i] = make([][]int64, k+1)
		for j := range dp[i] {
			dp[i][j] = make([]int64, 3)
		}
	}

	for j := 1; j <= k; j++ {
		dp[0][j][0] = 0
		dp[0][j][1] = int64(-prices[0])
		dp[0][j][2] = int64(prices[0])
	}

	for i := 1; i < n; i++ {
		for j := k; j >= 1; j-- {
			dp[i][j][0] = max(dp[i-1][j][0], max(dp[i-1][j][1]+int64(prices[i]), dp[i-1][j][2]-int64(prices[i])))
			dp[i][j][1] = max(dp[i-1][j][1], dp[i-1][j-1][0]-int64(prices[i]))
			dp[i][j][2] = max(dp[i-1][j][2], dp[i-1][j-1][0]+int64(prices[i]))
		}
	}
	return dp[n-1][k][0]
}

// @lc code=end

/*
// @lcpr case=start
// [1,7,9,8,2]\n2\n
// @lcpr case=end

// @lcpr case=start
// [12,16,19,19,8,1,19,13,9]\n3\n
// @lcpr case=end

*/
