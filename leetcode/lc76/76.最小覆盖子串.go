/*
 * @lc app=leetcode.cn id=76 lang=golang
 * @lcpr version=20004
 *
 * [76] 最小覆盖子串
 *
 * https://leetcode.cn/problems/minimum-window-substring/description/
 *
 * algorithms
 * Hard (46.66%)
 * Likes:    3085
 * Dislikes: 0
 * Total Accepted:    690.1K
 * Total Submissions: 1.5M
 * Testcase Example:  '"ADOBECODEBANC"\n"ABC"'
 *
 * 给你一个字符串 s 、一个字符串 t 。返回 s 中涵盖 t 所有字符的最小子串。如果 s 中不存在涵盖 t 所有字符的子串，则返回空字符串 ""
 * 。
 *
 *
 *
 * 注意：
 *
 *
 * 对于 t 中重复字符，我们寻找的子字符串中该字符数量必须不少于 t 中该字符数量。
 * 如果 s 中存在这样的子串，我们保证它是唯一的答案。
 *
 *
 *
 *
 * 示例 1：
 *
 * 输入：s = "ADOBECODEBANC", t = "ABC"
 * 输出："BANC"
 * 解释：最小覆盖子串 "BANC" 包含来自字符串 t 的 'A'、'B' 和 'C'。
 *
 *
 * 示例 2：
 *
 * 输入：s = "a", t = "a"
 * 输出："a"
 * 解释：整个字符串 s 是最小覆盖子串。
 *
 *
 * 示例 3:
 *
 * 输入: s = "a", t = "aa"
 * 输出: ""
 * 解释: t 中两个字符 'a' 均应包含在 s 的子串中，
 * 因此没有符合条件的子字符串，返回空字符串。
 *
 *
 *
 * 提示：
 *
 *
 * ^m == s.length
 * ^n == t.length
 * 1 <= m, n <= 10^5
 * s 和 t 由英文字母组成
 *
 *
 *
 * 进阶：你能设计一个在 o(m+n) 时间内解决此问题的算法吗？
 */
package lc0076

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func minWindow(s string, t string) string {
	win := [128]int{}
	diff := 0
	for _, c := range t {
		if win[c] == 0 {
			diff++
		}
		win[c]++
	}

	l, r := -1, len(s)
	i := 0
	for j, c := range s {
		win[c]--
		if win[c] == 0 {
			diff--
		}
		for diff == 0 {
			if r-l > j-i {
				l, r = i, j
			}
			d := s[i]
			if win[d] == 0 {
				diff++
			}
			win[d]++
			i++
		}
	}
	if l < 0 {
		return ""
	}
	return s[l : r+1]
}

// @lc code=end

/*
// @lcpr case=start
// "ADOBECODEBANC"\n"ABC"\n
// @lcpr case=end

// @lcpr case=start
// "a"\n"a"\n
// @lcpr case=end

// @lcpr case=start
// "a"\n"aa"\n
// @lcpr case=end

*/
