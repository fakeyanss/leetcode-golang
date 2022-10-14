/*
 * @lc app=leetcode.cn id=940 lang=golang
 *
 * [940] 不同的子序列 II
 *
 * https://leetcode.cn/problems/distinct-subsequences-ii/description/
 *
 * algorithms
 * Hard (43.47%)
 * Likes:    223
 * Dislikes: 0
 * Total Accepted:    15.8K
 * Total Submissions: 31.1K
 * Testcase Example:  '"abc"'
 *
 * 给定一个字符串 s，计算 s 的 不同非空子序列 的个数。因为结果可能很大，所以返回答案需要对 10^9 + 7 取余 。
 *
 * 字符串的 子序列 是经由原字符串删除一些（也可能不删除）字符但不改变剩余字符相对位置的一个新字符串。
 *
 *
 * 例如，"ace" 是 "abcde" 的一个子序列，但 "aec" 不是。
 *
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：s = "abc"
 * 输出：7
 * 解释：7 个不同的子序列分别是 "a", "b", "c", "ab", "ac", "bc", 以及 "abc"。
 *
 *
 * 示例 2：
 *
 *
 * 输入：s = "aba"
 * 输出：6
 * 解释：6 个不同的子序列分别是 "a", "b", "ab", "ba", "aa" 以及 "aba"。
 *
 *
 * 示例 3：
 *
 *
 * 输入：s = "aaa"
 * 输出：3
 * 解释：3 个不同的子序列分别是 "a", "aa" 以及 "aaa"。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= s.length <= 2000
 * s 仅由小写英文字母组成
 *
 *
 *
 *
 */
package lc940

// @lc code=start
func distinctSubseqII(s string) int {
	// dp[i] 为以s[i]结尾的不同子序列个数
	// 如果s[i]不是重复字符，dp[i] = 2*dp[i-1]+1，
	// 解释：把s[i]添加到以s[i-1]结尾的每一个子序列末尾，以及s[i]本身单独为一个子序列
	// 如果s[i]是重复字符，设上一次出现的位置下标j，dp[i] = 2*dp[i-1]-dp[j-1]
	// 解释：对于s[j-1]之前的子序列，默认添加s[j]与s[i]的结果是重复的
	const mod int = 1e9 + 7
	n := len(s)
	dp := make([]int, n+1)
	last := make([]int, 26)
	for i := 0; i < len(last); i++ {
		last[i] = -1
	}

	for i := 0; i < n; i++ {
		ch := s[i] - 'a'
		dp[i+1] = (2*dp[i] + 1) % mod
		if last[ch] >= 0 {
			dp[i+1] = (dp[i+1] - 1 - dp[last[ch]] + mod) % mod
		}
		last[ch] = i
	}
	return dp[n]
}

// @lc code=end
