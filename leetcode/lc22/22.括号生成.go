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

// @lc code=start
var (
	res   []string
	track []byte
)

func generateParenthesis(n int) []string {
	res = make([]string, 0)
	track = make([]byte, 0)
	if n == 0 {
		return res
	}
	backtrack(n, n)
	return res
}

func backtrack(left, right int) {
	// 生成括号过程中，剩余的 { 一定比 } 少
	if left > right {
		return
	}
	if left < 0 || right < 0 {
		return
	}
	if left == 0 && right == 0 {
		res = append(res, string(track))
		return
	}

	track = append(track, '(')
	backtrack(left-1, right)
	track = track[:len(track)-1]

	track = append(track, ')')
	backtrack(left, right-1)
	track = track[:len(track)-1]
}

// @lc code=end
