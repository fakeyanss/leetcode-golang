/*
 * @lc app=leetcode.cn id=2110 lang=golang
 * @lcpr version=30305
 *
 * [2110] 股票平滑下跌阶段的数目
 *
 * https://leetcode.cn/problems/number-of-smooth-descent-periods-of-a-stock/description/
 *
 * algorithms
 * Medium (60.52%)
 * Likes:    58
 * Dislikes: 0
 * Total Accepted:    25.4K
 * Total Submissions: 42.1K
 * Testcase Example:  '[3,2,1,4]\n[8,6,7,7]\n[1]'
 *
 * 给你一个整数数组 prices ，表示一支股票的历史每日股价，其中 prices[i] 是这支股票第 i 天的价格。
 *
 * 一个 平滑下降的阶段 定义为：对于 连续一天或者多天 ，每日股价都比 前一日股价恰好少 1 ，这个阶段第一天的股价没有限制。
 *
 * 请你返回 平滑下降阶段 的数目。
 *
 *
 *
 * 示例 1：
 *
 * 输入：prices = [3,2,1,4]
 * 输出：7
 * 解释：总共有 7 个平滑下降阶段：
 * [3], [2], [1], [4], [3,2], [2,1] 和 [3,2,1]
 * 注意，仅一天按照定义也是平滑下降阶段。
 *
 *
 * 示例 2：
 *
 * 输入：prices = [8,6,7,7]
 * 输出：4
 * 解释：总共有 4 个连续平滑下降阶段：[8], [6], [7] 和 [7]
 * 由于 8 - 6 ≠ 1 ，所以 [8,6] 不是平滑下降阶段。
 *
 *
 * 示例 3：
 *
 * 输入：prices = [1]
 * 输出：1
 * 解释：总共有 1 个平滑下降阶段：[1]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= prices.length <= 10^5
 * 1 <= prices[i] <= 10^5
 *
 *
 */
package lc2110

// @lc code=start
func getDescentPeriods(prices []int) int64 {
	// dp[i]表示以i结尾的平滑下降的个数
	// dp[0] = 1
	// dp[i] = 1 or 1+dp[i-1]
	n := len(prices)
	dp := make([]int64, n)
	dp[0] = 1
	for i := 1; i < n; i++ {
		dp[i] = 1
		if prices[i] == prices[i-1]-1 {
			dp[i] += dp[i-1]
		}
	}
	var res int64
	for _, v := range dp {
		res += v
	}
	return res
}

// @lc code=end

/*
// @lcpr case=start
// [3,2,1,4]\n
// @lcpr case=end

// @lcpr case=start
// [8,6,7,7]\n
// @lcpr case=end

// @lcpr case=start
// [1]\n
// @lcpr case=end

*/
