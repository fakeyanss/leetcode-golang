/*
 * @lc app=leetcode.cn id=72 lang=golang
 * @lcpr version=20004
 *
 * [72] 编辑距离
 *
 * https://leetcode.cn/problems/edit-distance/description/
 *
 * algorithms
 * Medium (63.28%)
 * Likes:    3551
 * Dislikes: 0
 * Total Accepted:    577.1K
 * Total Submissions: 911.9K
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

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func minDistance(word1 string, word2 string) int {
	// dp[i][j] 表示 word1[:i]到word2[:j]的最小编辑距离
	m, n := len(word1), len(word2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
		dp[i][0] = i
	}
	for j := range dp[0] {
		dp[0][j] = j
	}

	for i := range m {
		for j := range n {
			if word1[i] == word2[j] {
				dp[i+1][j+1] = dp[i][j]
			} else {
				// dp[i][j+1] <- dp[i+1][j+1] : 删除
				// dp[i+1][j] <- dp[i+1][j+1] : 插入
				// dp[i][j] <- dp[i+1][j+1] : 替换
				dp[i+1][j+1] = 1 + min(dp[i][j+1], dp[i+1][j], dp[i][j])
			}
		}
	}
	return dp[m][n]
}

// @lc code=end

/*
// @lcpr case=start
// "horse"\n"ros"\n
// @lcpr case=end

// @lcpr case=start
// "intention"\n"execution"\n
// @lcpr case=end

*/
