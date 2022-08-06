/*
 * @lc app=leetcode.cn id=76 lang=golang
 *
 * [76] 最小覆盖子串
 *
 * https://leetcode.cn/problems/minimum-window-substring/description/
 *
 * algorithms
 * Hard (44.62%)
 * Likes:    2047
 * Dislikes: 0
 * Total Accepted:    325.1K
 * Total Submissions: 728.7K
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
 *
 * 输入：s = "ADOBECODEBANC", t = "ABC"
 * 输出："BANC"
 *
 *
 * 示例 2：
 *
 *
 * 输入：s = "a", t = "a"
 * 输出："a"
 *
 *
 * 示例 3:
 *
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
 * 1
 * s 和 t 由英文字母组成
 *
 *
 *
 * 进阶：你能设计一个在 o(n) 时间内解决此问题的算法吗？
 */
package lc

import "math"

// @lc code=start
func minWindow(s string, t string) string {
	// 存储t的每个字符出现个数
	need := make(map[byte]int)
	for c := range t {
		need[t[c]]++
	}
	left, right := 0, 0
	// 存储窗口中字符出现个数
	window := make(map[byte]int)
	// 记录窗口中满足need条件的字符个数
	valid := 0
	// 记录最小子串的起始索引和长度
	start, length := 0, math.MaxInt
	for right < len(s) {
		c := s[right]
		right++
		// 如果当前字符是need包含的，就加入到window
		if _, ok := need[c]; ok {
			window[c]++
			// 如果出现次数满足了，valid记录一次
			if window[c] == need[c] {
				valid++
			}
		}

		// valid次数与need长度一致，说明所有字符都找到了
		// 开始收缩窗口
		for valid == len(need) {
			// 更新最小覆盖子串
			if right-left < length {
				start = left
				length = right - left
			}
			d := s[left]
			left++
			if _, ok := need[d]; ok {
				if window[d] == need[d] {
					valid--
				}
				window[d]--
			}
		}
	}

	if length == math.MaxInt {
		return ""
	}
	return s[start : start+length]
}

// @lc code=end
