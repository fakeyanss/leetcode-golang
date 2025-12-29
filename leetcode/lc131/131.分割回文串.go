/*
 * @lc app=leetcode.cn id=131 lang=golang
 * @lcpr version=30305
 *
 * [131] 分割回文串
 *
 * https://leetcode.cn/problems/palindrome-partitioning/description/
 *
 * algorithms
 * Medium (75.39%)
 * Likes:    2137
 * Dislikes: 0
 * Total Accepted:    697.3K
 * Total Submissions: 925.2K
 * Testcase Example:  '"aab"\n"a"'
 *
 * 给你一个字符串 s，请你将 s 分割成一些 子串，使每个子串都是 回文串 。返回 s 所有可能的分割方案。
 *
 *
 *
 * 示例 1：
 *
 * 输入：s = "aab"
 * 输出：[["a","a","b"],["aa","b"]]
 *
 *
 * 示例 2：
 *
 * 输入：s = "a"
 * 输出：[["a"]]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= s.length <= 16
 * s 仅由小写英文字母组成
 *
 *
 */
package lc131

// @lc code=start
func partition(s string) [][]string {
	n := len(s)
	// 优化：预处理s[i..j]是否为回文串
	dp := make([][]bool, n) // dp[i][j]表示s[i..j]是否为回文串
	for i := range dp {
		dp[i] = make([]bool, n)
		dp[i][i] = true
	}
	for length := 2; length <= n; length++ { // 遍历所有子串长度2到n
		for i := 0; i+length <= n; i++ { // 遍历所有的区间[i,j]
			j := i + length - 1
			if length == 2 {
				dp[i][j] = s[i] == s[j]
			} else {
				dp[i][j] = (s[i] == s[j]) && dp[i+1][j-1]
			}
		}
	}

	var res [][]string
	var track []string

	var dfs func(int, int)
	// i为要分割的位置，start为回文串的起点
	dfs = func(i, start int) {
		if i == n { // 分割完成
			res = append(res, append([]string{}, track...))
			return
		}

		if i < n-1 { // i=n-1 时只能分割
			dfs(i+1, start) // 不选s[i]为分割点
		}

		substr := s[start : i+1] // 选s[i]为分割点
		if dp[start][i] {
			track = append(track, substr)
			dfs(i+1, i+1) // 新的分割点和回文串从s[i+1]起始
			track = track[:len(track)-1]
		}
	}
	dfs(0, 0)
	return res
}

// @lc code=end

/*
// @lcpr case=start
// "aab"\n
// @lcpr case=end

// @lcpr case=start
// "a"\n
// @lcpr case=end

*/
