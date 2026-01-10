/*
 * @lc app=leetcode.cn id=1411 lang=golang
 * @lcpr version=30305
 *
 * [1411] 给 N x 3 网格图涂色的方案数
 *
 * https://leetcode.cn/problems/number-of-ways-to-paint-n-3-grid/description/
 *
 * algorithms
 * Hard (59.92%)
 * Likes:    156
 * Dislikes: 0
 * Total Accepted:    19.8K
 * Total Submissions: 28.8K
 * Testcase Example:  '1\n5000'
 *
 * 你有一个 n x 3 的网格图 grid ，你需要用 红，黄，绿
 * 三种颜色之一给每一个格子上色，且确保相邻格子颜色不同（也就是有相同水平边或者垂直边的格子颜色不同）。
 *
 * 给你网格图的行数 n 。
 *
 * 请你返回给 grid 涂色的方案数。由于答案可能会非常大，请你返回答案对 10^9 + 7 取余的结果。
 *
 *
 *
 * 示例 1：
 *
 * 输入：n = 1
 * 输出：12
 * 解释：总共有 12 种可行的方法：
 *
 *
 *
 * 示例 2：
 *
 * 输入：n = 2
 * 输出：54
 *
 *
 * 示例 3：
 *
 * 输入：n = 3
 * 输出：246
 *
 *
 * 示例 4：
 *
 * 输入：n = 7
 * 输出：106494
 *
 *
 * 示例 5：
 *
 * 输入：n = 5000
 * 输出：30228214
 *
 *
 *
 *
 * 提示：
 *
 *
 * n == grid.length
 * grid[i].length == 3
 * 1 <= n <= 5000
 *
 *
 */
package lc1411

// @lc code=start
const mod = 1e9 + 7

// 思路：DP
func numOfWays(n int) int {
	// 可以排列的颜色只有两类，ABA或ABC，各有6种排列
	// ABA：010，020，101，121，202，212
	// ABC：012，021，102，120，201，210
	// 状态转移：
	// 当第i-1行是ABC，第i行是ABC：可行方案数为2
	// 当第i-1行是ABC，第i行是ABA：可行方案数为2
	// 当第i-1行是ABA，第i行是ABC：可行方案数为2
	// 当第i-1行是ABA，第i行是ABA：可行方案数为3
	// dp[i][0|1]表示一共i行，第i行为ABC或ABA排列的方案数
	// dp[i][0]=2*dp[i-1][0]+2*dp[i-1][1]
	// dp[i][1]=2*dp[i-1][0]+3*dp[i-1][1]
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, 2)
	}
	dp[1][0] = 6
	dp[1][1] = 6
	for i := 2; i <= n; i++ {
		dp[i][0] = (2*dp[i-1][0] + 2*dp[i-1][1]) % mod
		dp[i][1] = (2*dp[i-1][0] + 3*dp[i-1][1]) % mod
	}
	return (dp[n][0] + dp[n][1]) % mod
}

// @lc code=end

/*
// @lcpr case=start
// 1\n
// @lcpr case=end

// @lcpr case=start
// 5000\n
// @lcpr case=end

*/
