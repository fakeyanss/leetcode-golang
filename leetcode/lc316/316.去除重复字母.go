/*
 * @lc app=leetcode.cn id=316 lang=golang
 *
 * [316] 去除重复字母
 *
 * https://leetcode.cn/problems/remove-duplicate-letters/description/
 *
 * algorithms
 * Medium (47.94%)
 * Likes:    789
 * Dislikes: 0
 * Total Accepted:    96.7K
 * Total Submissions: 201.6K
 * Testcase Example:  '"bcabc"'
 *
 * 给你一个字符串 s ，请你去除字符串中重复的字母，使得每个字母只出现一次。需保证 返回结果的字典序最小（要求不能打乱其他字符的相对位置）。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：s = "bcabc"
 * 输出："abc"
 *
 *
 * 示例 2：
 *
 *
 * 输入：s = "cbacdcbc"
 * 输出："acdb"
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= s.length <= 10^4
 * s 由小写英文字母组成
 *
 *
 *
 *
 * 注意：该题与 1081
 * https://leetcode-cn.com/problems/smallest-subsequence-of-distinct-characters
 * 相同
 *
 */
package lc316

// @lc code=start
func removeDuplicateLetters(s string) string {
	// 维护一个计数器记录字符串中字符的数量
	// 因为输入为 ASCII 字符，大小 256 够用了
	cnt := make([]byte, 256)
	for i := range s {
		cnt[s[i]]++
	}

	// 非重复字符存储到stk
	stk := make([]byte, 0)
	// 维护字符是否在栈中
	inStack := make([]bool, 256)

	for i := range s {
		c := s[i]
		cnt[c]--
		// 如果字符 c 存在栈中，直接跳过
		if inStack[c] {
			continue
		}
		// 插入之前，和之前的元素比较一下大小
		// 如果字典序比前面的小，pop 前面的元素
		for len(stk) > 0 && stk[len(stk)-1] > c {
			// 如果后面不存在栈顶元素了，则不pop
			if cnt[stk[len(stk)-1]] == 0 {
				break
			}
			// 如果后面还有，就pop
			inStack[stk[len(stk)-1]] = false
			stk = stk[:len(stk)-1]
		}
		// 若不存在，则插入栈顶并标记为存在
		stk = append(stk, c)
		inStack[c] = true
	}

	return string(stk)
}

// @lc code=end
