/*
 * @lc app=leetcode.cn id=516 lang=golang
 *
 * [516] 最长回文子序列
 *
 * https://leetcode.cn/problems/longest-palindromic-subsequence/description/
 *
 * algorithms
 * Medium (66.90%)
 * Likes:    886
 * Dislikes: 0
 * Total Accepted:    147.2K
 * Total Submissions: 220K
 * Testcase Example:  '"bbbab"'
 *
 * 给你一个字符串 s ，找出其中最长的回文子序列，并返回该序列的长度。
 *
 * 子序列定义为：不改变剩余字符顺序的情况下，删除某些字符或者不删除任何字符形成的一个序列。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：s = "bbbab"
 * 输出：4
 * 解释：一个可能的最长回文子序列为 "bbbb" 。
 *
 *
 * 示例 2：
 *
 *
 * 输入：s = "cbbd"
 * 输出：2
 * 解释：一个可能的最长回文子序列为 "bb" 。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1
 * s 仅由小写英文字母组成
 *
 *
 */
package lc516

// @lc code=start
func longestPalindromeSubseq(s string) int {
	n := len(s)
	memo := make([][]int, n)
	for i := range memo {
		memo[i] = make([]int, n)
	}
	return dpMemo(memo, s, 0, n-1)
	// return dpTable(s)
}

func dpMemo(memo [][]int, s string, i, j int) int {
	if i > j {
		return 0
	}
	if i == j {
		return 1
	}
	if memo[i][j] != 0 {
		return memo[i][j]
	}
	if s[i] == s[j] {
		memo[i][j] = 2 + dpMemo(memo, s, i+1, j-1)
	} else {
		memo[i][j] = maxInt(dpMemo(memo, s, i, j-1), dpMemo(memo, s, i+1, j))
	}
	return memo[i][j]
}

func dpTable(s string) int {
	// dp[i][j] 定义：在子串 s[i..j] 中，最长回文子序列的长度为 dp[i][j]
	n := len(s)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
		// base case
		dp[i][i] = 1
	}

	// dp[i][...]可能依赖dp[i+1][...]，所以i得反着遍历
	for i := n - 1; i >= 0; i-- {
		// dp[...][j]可能依赖dp[...][j-1], 所以j得正着遍历
		// 只需要遍历二位矩阵的右上区域
		for j := i + 1; j < n; j++ {
			if s[i] == s[j] {
				dp[i][j] = 2 + dp[i+1][j-1]
			} else {
				dp[i][j] = maxInt(dp[i+1][j], dp[i][j-1])
			}
		}
	}
	return dp[0][n-1]
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
