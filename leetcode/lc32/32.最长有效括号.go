/*
 * @lc app=leetcode.cn id=32 lang=golang
 * @lcpr version=30305
 *
 * [32] 最长有效括号
 *
 * https://leetcode.cn/problems/longest-valid-parentheses/description/
 *
 * algorithms
 * Hard (41.03%)
 * Likes:    2814
 * Dislikes: 0
 * Total Accepted:    665.5K
 * Total Submissions: 1.6M
 * Testcase Example:  '"(()"\n")()())"\n""'
 *
 * 给你一个只包含 '(' 和 ')' 的字符串，找出最长有效（格式正确且连续）括号 子串 的长度。
 *
 * 左右括号匹配，即每个左括号都有对应的右括号将其闭合的字符串是格式正确的，比如 "(()())"。
 *
 *
 *
 *
 *
 * 示例 1：
 *
 * 输入：s = "(()"
 * 输出：2
 * 解释：最长有效括号子串是 "()"
 *
 *
 * 示例 2：
 *
 * 输入：s = ")()())"
 * 输出：4
 * 解释：最长有效括号子串是 "()()"
 *
 *
 * 示例 3：
 *
 * 输入：s = ""
 * 输出：0
 *
 *
 *
 *
 * 提示：
 *
 *
 * 0 <= s.length <= 3 * 10^4
 * s[i] 为 '(' 或 ')'
 *
 *
 *
 *
 */
package lc32

// @lc code=start
func longestValidParentheses(s string) int {
	// dp[i]表示以s[i]结尾的子串的最长有效括号长度
	n := len(s)
	dp := make([]int, n)
	var res int
	for i := 1; i < n; i++ {
		if s[i] == ')' {
			if s[i-1] == '(' { // 组成一对'()'
				if i >= 2 {
					dp[i] = 2 + dp[i-2]
				} else {
					dp[i] = 2
				}
			} else { // 连续两个'))'
				// dp[i-1]为前一个s[i-1]和s[i-dp[i-1]-1]组成的有效括号长度
				if i-dp[i-1] > 0 && s[i-dp[i-1]-1] == '(' {
					if i-dp[i-1] >= 2 {
						dp[i] = dp[i-1] + dp[i-dp[i-1]-2] + 2
					} else {
						dp[i] = dp[i-1] + 2
					}
				}
			}
			res = max(res, dp[i])
		}
	}
	return res
}

// @lc code=end

/*
// @lcpr case=start
// "(()"\n
// @lcpr case=end

// @lcpr case=start
// ")()())"\n
// @lcpr case=end

// @lcpr case=start
// ""\n
// @lcpr case=end

*/
