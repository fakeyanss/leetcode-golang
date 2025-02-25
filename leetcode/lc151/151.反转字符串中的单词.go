/*
 * @lc app=leetcode.cn id=151 lang=golang
 * @lcpr version=20004
 *
 * [151] 反转字符串中的单词
 *
 * https://leetcode.cn/problems/reverse-words-in-a-string/description/
 *
 * algorithms
 * Medium (57.32%)
 * Likes:    1256
 * Dislikes: 0
 * Total Accepted:    687.9K
 * Total Submissions: 1.2M
 * Testcase Example:  '"the sky is blue"'
 *
 * 给你一个字符串 s ，请你反转字符串中 单词 的顺序。
 *
 * 单词 是由非空格字符组成的字符串。s 中使用至少一个空格将字符串中的 单词 分隔开。
 *
 * 返回 单词 顺序颠倒且 单词 之间用单个空格连接的结果字符串。
 *
 * 注意：输入字符串 s中可能会存在前导空格、尾随空格或者单词间的多个空格。返回的结果字符串中，单词间应当仅用单个空格分隔，且不包含任何额外的空格。
 *
 *
 *
 * 示例 1：
 *
 * 输入：s = "the sky is blue"
 * 输出："blue is sky the"
 *
 *
 * 示例 2：
 *
 * 输入：s = "  hello world  "
 * 输出："world hello"
 * 解释：反转后的字符串中不能存在前导空格和尾随空格。
 *
 *
 * 示例 3：
 *
 * 输入：s = "a good   example"
 * 输出："example good a"
 * 解释：如果两个单词间有多余的空格，反转后的字符串需要将单词间的空格减少到仅有一个。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= s.length <= 10^4
 * s 包含英文大小写字母、数字和空格 ' '
 * s 中 至少存在一个 单词
 *
 *
 *
 *
 *
 *
 *
 * 进阶：如果字符串在你使用的编程语言中是一种可变数据类型，请尝试使用 O(1) 额外空间复杂度的 原地 解法。
 *
 */
package lc151

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func reverseWords(s string) string {
	// 倒序遍历s，栈存储每个单词，遇到空格出栈
	n := len(s)
	var stack []byte
	var res []byte
	for i := n - 1; i >= 0; i-- {
		if s[i] == ' ' && len(stack) == 0 {
			continue
		}
		if s[i] == ' ' || i == 0 {
			if i == 0 && s[i] != ' ' {
				stack = append(stack, s[i])
			}
			// 出栈
			for len(stack) > 0 {
				v := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				res = append(res, v)
			}
			res = append(res, ' ')
		} else {
			stack = append(stack, s[i])
		}
	}

	return string(res[:len(res)-1]) // 去除末尾的空格

	// 如果要原地修改，就反转字符串s，再挨个反转单词
}

// @lc code=end

/*
// @lcpr case=start
// "the sky is blue"\n
// @lcpr case=end

// @lcpr case=start
// "  hello world  "\n
// @lcpr case=end

// @lcpr case=start
// "a good   example"\n
// @lcpr case=end

*/
