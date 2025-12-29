/*
 * @lc app=leetcode.cn id=17 lang=golang
 * @lcpr version=20003
 *
 * [17] 电话号码的字母组合
 *
 * https://leetcode.cn/problems/letter-combinations-of-a-phone-number/description/
 *
 * algorithms
 * Medium (60.92%)
 * Likes:    2955
 * Dislikes: 0
 * Total Accepted:    985.8K
 * Total Submissions: 1.6M
 * Testcase Example:  '"23"'
 *
 * 给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。答案可以按 任意顺序 返回。
 *
 * 给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。
 *
 *
 *
 *
 *
 * 示例 1：
 *
 * 输入：digits = "23"
 * 输出：["ad","ae","af","bd","be","bf","cd","ce","cf"]
 *
 *
 * 示例 2：
 *
 * 输入：digits = ""
 * 输出：[]
 *
 *
 * 示例 3：
 *
 * 输入：digits = "2"
 * 输出：["a","b","c"]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 0 <= digits.length <= 4
 * digits[i] 是范围 ['2', '9'] 的一个数字。
 *
 *
 */
package lc17

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func letterCombinations(digits string) []string {
	mapping := [10]string{"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}

	var res []string
	n := len(digits)
	if n == 0 {
		return res
	}

	track := make([]byte, n)

	var dfs func(int)
	dfs = func(i int) {
		if i == n {
			res = append(res, string(track))
			return
		}
		for _, c := range mapping[digits[i]-'0'] {
			track[i] = byte(c) // 直接覆盖
			dfs(i + 1)
		}
	}
	dfs(0)
	return res
}

// @lc code=end

/*
// @lcpr case=start
// "23"\n
// @lcpr case=end

// @lcpr case=start
// ""\n
// @lcpr case=end

// @lcpr case=start
// "2"\n
// @lcpr case=end

*/
