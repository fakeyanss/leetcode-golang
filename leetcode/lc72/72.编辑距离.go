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
	m, n := len(word1), len(word2)
	memo := make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		memo[i] = make([]int, n+1)
	}
	// dp[i,j]表示word1[0..i]到word2[0..j]的最小编辑距离
	var dp func(i, j int) int
	dp = func(i, j int) int {
		if i < 0 {
			return j + 1
		}
		if j < 0 {
			return i + 1
		}
		if memo[i][j] != 0 {
			return memo[i][j]
		}
		if word1[i] == word2[j] {
			memo[i][j] = dp(i-1, j-1)
		} else {
			memo[i][j] = min(
				dp(i, j-1),   // 插入
				dp(i-1, j-1), // 替换
				dp(i-1, j),   // 删除
			) + 1
		}
		return memo[i][j]
	}

	return dp(m-1, n-1)
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
