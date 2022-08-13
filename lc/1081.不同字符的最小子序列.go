/*
 * @lc app=leetcode.cn id=1081 lang=golang
 *
 * [1081] 不同字符的最小子序列
 *
 * https://leetcode.cn/problems/smallest-subsequence-of-distinct-characters/description/
 *
 * algorithms
 * Medium (58.18%)
 * Likes:    161
 * Dislikes: 0
 * Total Accepted:    21.7K
 * Total Submissions: 37.2K
 * Testcase Example:  '"bcabc"'
 *
 * 返回 s 字典序最小的子序列，该子序列包含 s 的所有不同字符，且只包含一次。
 *
 * 注意：该题与 316 https://leetcode.com/problems/remove-duplicate-letters/ 相同
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
 * 1
 * s 由小写英文字母组成
 *
 *
 */
package lc

// @lc code=start
func smallestSubsequence(s string) string {
	cnt := make([]rune, 26)
	for _, c := range s {
		cnt[c-'a']++
	}

	stk := make([]rune, 0)
	vis := make([]bool, 26)

	for _, c := range s {
		cnt[c-'a']--
		if vis[c-'a'] {
			continue
		}

		for len(stk) > 0 && stk[len(stk)-1] > c {
			if cnt[stk[len(stk)-1]-'a'] == 0 {
				break
			}
			vis[stk[len(stk)-1]-'a'] = false
			stk = stk[:len(stk)-1]
		}

		stk = append(stk, c)
		vis[c-'a'] = true
	}

	return string(stk)
}

// @lc code=end
