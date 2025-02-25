/*
 * @lc app=leetcode.cn id=115 lang=golang
 * @lcpr version=20004
 *
 * [115] 不同的子序列
 *
 * https://leetcode.cn/problems/distinct-subsequences/description/
 *
 * algorithms
 * Hard (52.76%)
 * Likes:    1307
 * Dislikes: 0
 * Total Accepted:    207.8K
 * Total Submissions: 392.9K
 * Testcase Example:  '"rabbbit"\n"rabbit"'
 *
 * 给你两个字符串 s 和 t ，统计并返回在 s 的 子序列 中 t 出现的个数，结果需要对 10^9 + 7 取模。
 *
 *
 *
 * 示例 1：
 *
 * 输入：s = "rabbbit", t = "rabbit"
 * 输出：3
 * 解释：
 * 如下所示, 有 3 种可以从 s 中得到 "rabbit" 的方案。
 * rabbbit
 * rabbbit
 * rabbbit
 *
 * 示例 2：
 *
 * 输入：s = "babgbag", t = "bag"
 * 输出：5
 * 解释：
 * 如下所示, 有 5 种可以从 s 中得到 "bag" 的方案。
 * babgbag
 * babgbag
 * babgbag
 * babgbag
 * babgbag
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= s.length, t.length <= 1000
 * s 和 t 由英文字母组成
 *
 *
 */
package lc115

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func numDistinct(s string, t string) int {
	// dp[i][j] 表示 s[:i-1] 和 t[:j-1] 下的子序列个数
	// dp[i][0]=1, dp[0][j]=0
	// 区分 s[i] == t[j] 和 s[i] != t[j] 的情况来讨论
	m, n := len(s), len(t)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for i := 0; i < m+1; i++ {
		for j := 0; j < n+1; j++ {
			if j == 0 {
				dp[i][j] = 1
			} else if i == 0 {
				dp[i][j] = 0
			} else {
				if s[i-1] == t[j-1] {
					dp[i][j] = dp[i-1][j-1] + dp[i-1][j] // 选择匹配 + 选择不匹配
				} else {
					dp[i][j] = dp[i-1][j] // 没法匹配
				}
			}
		}
	}
	return dp[m][n]
}

// @lc code=end

/*
// @lcpr case=start
// "rabbbit"\n"rabbit"\n
// @lcpr case=end

// @lcpr case=start
// "babgbag"\n"bag"\n
// @lcpr case=end

*/
