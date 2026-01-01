/*
 * @lc app=leetcode.cn id=5 lang=golang
 *
 * [5] 最长回文子串
 *
 * https://leetcode.cn/problems/longest-palindromic-substring/description/
 *
 * algorithms
 * Medium (36.88%)
 * Likes:    5465
 * Dislikes: 0
 * Total Accepted:    1.1M
 * Total Submissions: 3M
 * Testcase Example:  '"babad"'
 *
 * 给你一个字符串 s，找到 s 中最长的回文子串。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：s = "babad"
 * 输出："bab"
 * 解释："aba" 同样是符合题意的答案。
 *
 *
 * 示例 2：
 *
 *
 * 输入：s = "cbbd"
 * 输出："bb"
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= s.length <= 1000
 * s 仅由数字和英文字母组成
 *
 *
 */
package lc5

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func longestPalindrome(s string) string {
	n := len(s)
	if n < 2 {
		return s
	}

	// dp[i][j] 表示s[i..j]是否是回文串
	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, n)
		dp[i][i] = true // 长度1的串都是回文串
	}

	begin, maxLen := 0, 1
	for l := 2; l <= n; l++ { // 枚举所有的子串长度
		for i := 0; i <= n-l; i++ { // 枚举子串的左边界
			j := i + l - 1 // 子串右边界
			if j >= n {
				break
			}

			if s[i] == s[j] {
				if j-i < 3 {
					dp[i][j] = true
				} else {
					// 因为l的遍历顺序是从小到大，计算长的子串时，短的子串一定都计算过
					dp[i][j] = dp[i+1][j-1]
				}
			}

			if dp[i][j] && l > maxLen { // 更新最长回文串的位置
				begin, maxLen = i, l
			}
		}
	}
	return s[begin : begin+maxLen]
}

// @lc code=end
