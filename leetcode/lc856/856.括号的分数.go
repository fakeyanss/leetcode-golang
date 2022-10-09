/*
 * @lc app=leetcode.cn id=856 lang=golang
 *
 * [856] 括号的分数
 *
 * https://leetcode.cn/problems/score-of-parentheses/description/
 *
 * algorithms
 * Medium (63.33%)
 * Likes:    437
 * Dislikes: 0
 * Total Accepted:    46.8K
 * Total Submissions: 68.5K
 * Testcase Example:  '"()"'
 *
 * 给定一个平衡括号字符串 S，按下述规则计算该字符串的分数：
 *
 *
 * () 得 1 分。
 * AB 得 A + B 分，其中 A 和 B 是平衡括号字符串。
 * (A) 得 2 * A 分，其中 A 是平衡括号字符串。
 *
 *
 *
 *
 * 示例 1：
 *
 * 输入： "()"
 * 输出： 1
 *
 *
 * 示例 2：
 *
 * 输入： "(())"
 * 输出： 2
 *
 *
 * 示例 3：
 *
 * 输入： "()()"
 * 输出： 2
 *
 *
 * 示例 4：
 *
 * 输入： "(()(()))"
 * 输出： 6
 *
 *
 *
 *
 * 提示：
 *
 *
 * S 是平衡括号字符串，且只含有 ( 和 ) 。
 * 2 <= S.length <= 50
 *
 *
 */
package lc856

// @lc code=start
func scoreOfParentheses(s string) int {
	// 对于每个 (，我们将深度加一，对于每个 )，我们将深度减一。当我们遇到 () 时，我们将 2^d 加到答案中
	depth := 0
	res := 0
	for i, v := range s {
		if v == '(' {
			depth++
		} else {
			depth--
			if s[i-1] == '(' {
				res += 1 << depth
			}
		}
	}
	return res
}

// @lc code=end
