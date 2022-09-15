/*
 * @lc app=leetcode.cn id=1312 lang=golang
 *
 * [1312] 让字符串成为回文串的最少插入次数
 *
 * https://leetcode.cn/problems/minimum-insertion-steps-to-make-a-string-palindrome/description/
 *
 * algorithms
 * Hard (68.64%)
 * Likes:    156
 * Dislikes: 0
 * Total Accepted:    19K
 * Total Submissions: 27.6K
 * Testcase Example:  '"zzazz"'
 *
 * 给你一个字符串 s ，每一次操作你都可以在字符串的任意位置插入任意字符。
 *
 * 请你返回让 s 成为回文串的 最少操作次数 。
 *
 * 「回文串」是正读和反读都相同的字符串。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：s = "zzazz"
 * 输出：0
 * 解释：字符串 "zzazz" 已经是回文串了，所以不需要做任何插入操作。
 *
 *
 * 示例 2：
 *
 *
 * 输入：s = "mbadm"
 * 输出：2
 * 解释：字符串可变为 "mbdadbm" 或者 "mdbabdm" 。
 *
 *
 * 示例 3：
 *
 *
 * 输入：s = "leetcode"
 * 输出：5
 * 解释：插入 5 个字符后字符串变为 "leetcodocteel" 。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= s.length <= 500
 * s 中所有字符都是小写字母。
 *
 *
 */
package lc1312

// @lc code=start
func minInsertions(s string) int {
	// n := len(s)
	// memo := make([][]int, n)
	// for i := range memo {
	// 	memo[i] = make([]int, n)
	// }
	// return dpMemo(memo, s, 0, n-1)
	return dpTable(s)
}

// dp(s, i, j) 定义：字符串 s 在区间 [i...j] 中让字符串成为回文串的最少插入次数
func dpMemo(memo [][]int, s string, i, j int) int {
	if i > j {
		return 0
	}
	if memo[i][j] != 0 {
		return memo[i][j]
	}
	if s[i] == s[j] {
		memo[i][j] = dpMemo(memo, s, i+1, j-1)
	} else {
		memo[i][j] = 1 + minInt(dpMemo(memo, s, i+1, j), dpMemo(memo, s, i, j-1))
	}
	return memo[i][j]
}

func dpTable(s string) int {
	n := len(s)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	for i := n - 1; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1]
			} else {
				dp[i][j] = minInt(dp[i+1][j], dp[i][j-1]) + 1
			}
		}
	}
	return dp[0][n-1]
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end
