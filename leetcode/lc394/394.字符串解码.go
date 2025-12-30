/*
 * @lc app=leetcode.cn id=394 lang=golang
 * @lcpr version=30305
 *
 * [394] 字符串解码
 *
 * https://leetcode.cn/problems/decode-string/description/
 *
 * algorithms
 * Medium (61.01%)
 * Likes:    2114
 * Dislikes: 0
 * Total Accepted:    553.7K
 * Total Submissions: 906.8K
 * Testcase Example:  '"3[a]2[bc]"\n"3[a2[c]]"\n"2[abc]3[cd]ef"'
 *
 * 给定一个经过编码的字符串，返回它解码后的字符串。
 *
 * 编码规则为: k[encoded_string]，表示其中方括号内部的 encoded_string 正好重复 k 次。注意 k 保证为正整数。
 *
 * 你可以认为输入字符串总是有效的；输入字符串中没有额外的空格，且输入的方括号总是符合格式要求的。
 *
 * 此外，你可以认为原始数据不包含数字，所有的数字只表示重复的次数 k ，例如不会出现像 3a 或 2[4] 的输入。
 *
 * 测试用例保证输出的长度不会超过 10^5。
 *
 *
 *
 * 示例 1：
 *
 * 输入：s = "3[a]2[bc]"
 * 输出："aaabcbc"
 *
 *
 * 示例 2：
 *
 * 输入：s = "3[a2[c]]"
 * 输出："accaccacc"
 *
 *
 * 示例 3：
 *
 * 输入：s = "2[abc]3[cd]ef"
 * 输出："abcabccdcdcdef"
 *
 *
 * 示例 4：
 *
 * 输入：s = "abc3[cd]xyz"
 * 输出："abccdcdcdxyz"
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= s.length <= 30
 * s 由小写英文字母、数字和方括号 '[]' 组成
 * s 保证是一个 有效 的输入。
 * s 中所有整数的取值范围为 [1, 300]
 *
 *
 */
package lc394

import (
	"strconv"
	"strings"
)

// @lc code=start
func decodeString(s string) string {
	numsMap := map[string]bool{"0": true, "1": true, "2": true, "3": true, "4": true, "5": true, "6": true, "7": true, "8": true, "9": true}

	var stk []string
	for i := range s {
		if s[i] != ']' {
			stk = append(stk, string(s[i]))
		} else {
			var tmpStr string // 暂存'[]'中间的字母
			for stk[len(stk)-1] != "[" {
				tmpStr = stk[len(stk)-1] + tmpStr
				stk = stk[:len(stk)-1]
			}
			stk = stk[:len(stk)-1] // 丢掉'['

			var tmpNumStr string // 暂存数字
			for len(stk) > 0 && numsMap[stk[len(stk)-1]] {
				tmpNumStr = stk[len(stk)-1] + tmpNumStr
				stk = stk[:len(stk)-1]
			}
			repeatNum, _ := strconv.Atoi(tmpNumStr)

			repeated := strings.Repeat(tmpStr, repeatNum)
			stk = append(stk, repeated) // 结果压回栈
		}
	}
	var res string
	for _, v := range stk {
		res += v
	}
	return res
}

// @lc code=end

/*
// @lcpr case=start
// "3[a]2[bc]"\n
// @lcpr case=end

// @lcpr case=start
// "3[a2[c]]"\n
// @lcpr case=end

// @lcpr case=start
// "2[abc]3[cd]ef"\n
// @lcpr case=end

*/
