/*
 * @lc app=leetcode.cn id=10 lang=golang
 *
 * [10] 正则表达式匹配
 *
 * https://leetcode.cn/problems/regular-expression-matching/description/
 *
 * algorithms
 * Hard (31.66%)
 * Likes:    3227
 * Dislikes: 0
 * Total Accepted:    318.7K
 * Total Submissions: 1M
 * Testcase Example:  '"aa"\n"a"'
 *
 * 给你一个字符串 s 和一个字符规律 p，请你来实现一个支持 '.' 和 '*' 的正则表达式匹配。
 *
 *
 * '.' 匹配任意单个字符
 * '*' 匹配零个或多个前面的那一个元素
 *
 *
 * 所谓匹配，是要涵盖 整个 字符串 s的，而不是部分字符串。
 *
 *
 * 示例 1：
 *
 *
 * 输入：s = "aa", p = "a"
 * 输出：false
 * 解释："a" 无法匹配 "aa" 整个字符串。
 *
 *
 * 示例 2:
 *
 *
 * 输入：s = "aa", p = "a*"
 * 输出：true
 * 解释：因为 '*' 代表可以匹配零个或多个前面的那一个元素, 在这里前面的元素就是 'a'。因此，字符串 "aa" 可被视为 'a' 重复了一次。
 *
 *
 * 示例 3：
 *
 *
 * 输入：s = "ab", p = ".*"
 * 输出：true
 * 解释：".*" 表示可匹配零个或多个（'*'）任意字符（'.'）。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= s.length <= 20
 * 1 <= p.length <= 30
 * s 只包含从 a-z 的小写字母。
 * p 只包含从 a-z 的小写字母，以及字符 . 和 *。
 * 保证每次出现字符 * 时，前面都匹配到有效的字符
 *
 *
 */
package lc10

// @lc code=start
var memo [][]int

func isMatch(s string, p string) bool {
	memo = make([][]int, len(s))
	for i := range memo {
		memo[i] = make([]int, len(p))
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}
	return dp(s, 0, p, 0)
}

// dp(s,i,p,j) = true, 表示s[i..]可以匹配p[j..]，否则表示无法匹配
func dp(s string, i int, p string, j int) bool {
	m, n := len(s), len(p)
	// base case
	if j == n {
		// 都匹配到字符串最右，结束
		return i == m
	}
	if i == m {
		// 只有s匹配到最右，需要看p[j..]的是否是成对的 ".*"，是的话可以匹配成功
		if (n-j)%2 == 1 {
			return false
		}
		for ; j+1 < n; j += 2 {
			if p[j+1] != '*' {
				return false
			}
		}
		return true
	}

	if memo[i][j] != -1 {
		return memo[i][j] == 1
	}

	res := false
	if s[i] == p[j] || p[j] == '.' {
		// 如果可以匹配
		if j < n-1 && p[j+1] == '*' {
			// 如果p[j+1] = "*"，则可以匹配多个s的字符
			res = dp(s, i, p, j+2) || dp(s, i+1, p, j)
		} else {
			// 如果只是p[j+1] != "*"，则只匹配一个s的字符
			res = dp(s, i+1, p, j+1)
		}
	} else {
		if j < n-1 && p[j+1] == '*' {
			// p[j+1]='*'，可以允许匹配0次
			res = dp(s, i, p, j+2)
		} else {
			res = false
		}
	}

	if res {
		memo[i][j] = 1
	} else {
		memo[i][j] = 0
	}
	return res
}

// @lc code=end
