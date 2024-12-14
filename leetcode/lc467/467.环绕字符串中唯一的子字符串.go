/*
 * @lc app=leetcode.cn id=467 lang=golang
 * @lcpr version=20004
 *
 * [467] 环绕字符串中唯一的子字符串
 *
 * https://leetcode.cn/problems/unique-substrings-in-wraparound-string/description/
 *
 * algorithms
 * Medium (51.92%)
 * Likes:    425
 * Dislikes: 0
 * Total Accepted:    44.6K
 * Total Submissions: 85.9K
 * Testcase Example:  '"a"'
 *
 * 定义字符串 base 为一个 "abcdefghijklmnopqrstuvwxyz" 无限环绕的字符串，所以 base 看起来是这样的：
 *
 *
 * "...zabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcd....".
 *
 *
 * 给你一个字符串 s ，请你统计并返回 s 中有多少 不同非空子串 也在 base 中出现。
 *
 *
 *
 * 示例 1：
 *
 * 输入：s = "a"
 * 输出：1
 * 解释：字符串 s 的子字符串 "a" 在 base 中出现。
 *
 *
 * 示例 2：
 *
 * 输入：s = "cac"
 * 输出：2
 * 解释：字符串 s 有两个子字符串 ("a", "c") 在 base 中出现。
 *
 *
 * 示例 3：
 *
 * 输入：s = "zab"
 * 输出：6
 * 解释：字符串 s 有六个子字符串 ("z", "a", "b", "za", "ab", and "zab") 在 base 中出现。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= s.length <= 10^5
 * s 由小写英文字母组成
 *
 *
 */
package lc467

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func findSubstringInWraproundString(s string) int {
	dp := [26]int{} // dp[i] 定义：s串中以字母i结尾的子串个数
	k := 0
	for i := 0; i < len(s); i++ {
		c := s[i]
		if i > 0 && (c-s[i-1] == 1 || s[i-1]-c == 25) {
			k++
		} else {
			k = 1
		}
		dp[c-'a'] = max(dp[c-'a'], k)
	}
	res := 0
	for _, v := range dp {
		res += v
	}
	return res
}

// @lc code=end

/*
// @lcpr case=start
// "a"\n
// @lcpr case=end

// @lcpr case=start
// "cac"\n
// @lcpr case=end

// @lcpr case=start
// "zab"\n
// @lcpr case=end

*/
