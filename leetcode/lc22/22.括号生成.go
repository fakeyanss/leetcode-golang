/*
 * @lc app=leetcode.cn id=22 lang=golang
 *
 * [22] 括号生成
 *
 * https://leetcode.cn/problems/generate-parentheses/description/
 *
 * algorithms
 * Medium (77.54%)
 * Likes:    2877
 * Dislikes: 0
 * Total Accepted:    593.7K
 * Total Submissions: 765.6K
 * Testcase Example:  '3'
 *
 * 数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：n = 3
 * 输出：["((()))","(()())","(())()","()(())","()()()"]
 *
 *
 * 示例 2：
 *
 *
 * 输入：n = 1
 * 输出：["()"]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= n <= 8
 *
 *
 */
package lc22

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start

func generateParenthesis(n int) []string {
	var res []string
	track := make([]byte, 2*n)
	var dfs func(int, int)
	dfs = func(left, right int) {
		if right == n {
			res = append(res, string(track))
			return
		}

		if left < n {
			track[left+right] = '(' // 直接覆盖
			dfs(left+1, right)
		}
		if right < left {
			track[left+right] = ')' // 直接覆盖
			dfs(left, right+1)
		}
	}
	dfs(0, 0)
	return res
}

// @lc code=end
