/*
 * @lc app=leetcode.cn id=72 lang=golang
 *
 * [72] 编辑距离
 *
 * https://leetcode.cn/problems/edit-distance/description/
 *
 * algorithms
 * Hard (62.72%)
 * Likes:    2591
 * Dislikes: 0
 * Total Accepted:    303.3K
 * Total Submissions: 483.6K
 * Testcase Example:  '"horse"\n"ros"'
 *
 * 给你两个单词 word1 和 word2， 请返回将 word1 转换成 word2 所使用的最少操作数  。
 *
 * 你可以对一个单词进行如下三种操作：
 *
 *
 * 插入一个字符
 * 删除一个字符
 * 替换一个字符
 *
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：word1 = "horse", word2 = "ros"
 * 输出：3
 * 解释：
 * horse -> rorse (将 'h' 替换为 'r')
 * rorse -> rose (删除 'r')
 * rose -> ros (删除 'e')
 *
 *
 * 示例 2：
 *
 *
 * 输入：word1 = "intention", word2 = "execution"
 * 输出：5
 * 解释：
 * intention -> inention (删除 't')
 * inention -> enention (将 'i' 替换为 'e')
 * enention -> exention (将 'n' 替换为 'x')
 * exention -> exection (将 'n' 替换为 'c')
 * exection -> execution (插入 'u')
 *
 *
 *
 *
 * 提示：
 *
 *
 * 0 <= word1.length, word2.length <= 500
 * word1 和 word2 由小写英文字母组成
 *
 *
 */
package lc72

import "math"

// @lc code=start
func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	memo := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		memo[i] = make([]int, n+1)
	}
	return dpMemo(memo, word1, len(word1)-1, word2, len(word2)-1)
	// return dpTable(word1, word2)
}

// 自顶向下，dp+备忘录
// dp(str1, i, str2, j), 表示str1[0...i], str2[0...j]的最小编辑距离
func dpMemo(memo [][]int, str1 string, i int, str2 string, j int) int {
	// base case
	if i < 0 {
		return j + 1
	}
	if j < 0 {
		return i + 1
	}

	if memo[i][j] != 0 {
		return memo[i][j]
	}

	if str1[i] == str2[j] {
		memo[i][j] = dpMemo(memo, str1, i-1, str2, j-1)
	} else {
		memo[i][j] = minInt(
			dpMemo(memo, str1, i-1, str2, j)+1,   // 删除
			dpMemo(memo, str1, i, str2, j-1)+1,   // 插入
			dpMemo(memo, str1, i-1, str2, j-1)+1, // 替换
		)
	}
	return memo[i][j]
}

// 自底向上，dp table
func dpTable(word1, word2 string) int {
	m, n := len(word1), len(word2)
	dp := make([][]int, m+1)

	// base case
	for i := range dp {
		dp[i] = make([]int, n+1)
		dp[i][0] = i
	}
	for j := range dp[0] {
		dp[0][j] = j
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = minInt(
					dp[i-1][j]+1,
					dp[i][j-1]+1,
					dp[i-1][j-1]+1,
				)
			}
		}
	}
	return dp[m][n]
}

func minInt(nums ...int) int {
	min := math.MaxInt
	for _, v := range nums {
		if min > v {
			min = v
		}
	}
	return min
}

// @lc code=end
