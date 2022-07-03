/*
 * @lc app=leetcode.cn id=28 lang=golang
 *
 * [28] 实现 strStr()
 *
 * https://leetcode.cn/problems/implement-strstr/description/
 *
 * algorithms
 * Easy (40.87%)
 * Likes:    1481
 * Dislikes: 0
 * Total Accepted:    674.9K
 * Total Submissions: 1.7M
 * Testcase Example:  '"hello"\n"ll"'
 *
 * 实现 strStr() 函数。
 *
 * 给你两个字符串 haystack 和 needle ，请你在 haystack 字符串中找出 needle 字符串出现的第一个位置（下标从 0
 * 开始）。如果不存在，则返回  -1 。
 *
 * 说明：
 *
 * 当 needle 是空字符串时，我们应当返回什么值呢？这是一个在面试中很好的问题。
 *
 * 对于本题而言，当 needle 是空字符串时我们应当返回 0 。这与 C 语言的 strstr() 以及 Java 的 indexOf()
 * 定义相符。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：haystack = "hello", needle = "ll"
 * 输出：2
 *
 *
 * 示例 2：
 *
 *
 * 输入：haystack = "aaaaa", needle = "bba"
 * 输出：-1
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= haystack.length, needle.length <= 10^4
 * haystack 和 needle 仅由小写英文字符组成
 *
 *
 */
package lc

// @lc code=start
func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	// 暴力求解
	p1, p2 := len(haystack), len(needle)
	for i := 0; i <= p1-p2; i++ {
		if haystack[i:i+p2] == needle {
			return i
		}
	}
	return -1
	// KMP, too hard
}

// @lc code=end
