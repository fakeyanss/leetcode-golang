/*
 * @lc app=leetcode.cn id=712 lang=golang
 *
 * [712] 两个字符串的最小ASCII删除和
 *
 * https://leetcode.cn/problems/minimum-ascii-delete-sum-for-two-strings/description/
 *
 * algorithms
 * Medium (68.21%)
 * Likes:    298
 * Dislikes: 0
 * Total Accepted:    25.6K
 * Total Submissions: 37.6K
 * Testcase Example:  '"sea"\n"eat"'
 *
 * 给定两个字符串s1 和 s2，返回 使两个字符串相等所需删除字符的 ASCII 值的最小和 。
 *
 *
 *
 * 示例 1:
 *
 *
 * 输入: s1 = "sea", s2 = "eat"
 * 输出: 231
 * 解释: 在 "sea" 中删除 "s" 并将 "s" 的值(115)加入总和。
 * 在 "eat" 中删除 "t" 并将 116 加入总和。
 * 结束时，两个字符串相等，115 + 116 = 231 就是符合条件的最小和。
 *
 *
 * 示例 2:
 *
 *
 * 输入: s1 = "delete", s2 = "leet"
 * 输出: 403
 * 解释: 在 "delete" 中删除 "dee" 字符串变成 "let"，
 * 将 100[d]+101[e]+101[e] 加入总和。在 "leet" 中删除 "e" 将 101[e] 加入总和。
 * 结束时，两个字符串都等于 "let"，结果即为 100+101+101+101 = 403 。
 * 如果改为将两个字符串转换为 "lee" 或 "eet"，我们会得到 433 或 417 的结果，比答案更大。
 *
 *
 *
 *
 * 提示:
 *
 *
 * 0 <= s1.length, s2.length <= 1000
 * s1 和 s2 由小写英文字母组成
 *
 *
 */
package lc712

// @lc code=start
func minimumDeleteSum(s1 string, s2 string) int {
	m, n := len(s1), len(s2)
	memo := make([][]int, m)
	for i := range memo {
		memo[i] = make([]int, n)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}
	return dp(memo, s1, 0, s2, 0)
}

// 定义：将 s1[i..] 和 s2[j..] 删除成相同字符串，最小的 ASCII 码之和为 dp(s1, i, s2, j)
func dp(memo [][]int, s1 string, i int, s2 string, j int) int {
	res := 0
	// base case
	if i == len(s1) {
		// s1遍历到头了，s2剩下的都得删除
		for ; j < len(s2); j++ {
			res += int(s2[j])
		}
		return res
	}
	if j == len(s2) {
		for ; i < len(s1); i++ {
			res += int(s1[i])
		}
		return res
	}

	if memo[i][j] != -1 {
		return memo[i][j]
	}

	if s1[i] == s2[j] {
		memo[i][j] = dp(memo, s1, i+1, s2, j+1)
	} else {
		memo[i][j] = minInt(
			int(s1[i])+dp(memo, s1, i+1, s2, j),
			int(s2[j])+dp(memo, s1, i, s2, j+1),
		)
	}

	return memo[i][j]
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end
