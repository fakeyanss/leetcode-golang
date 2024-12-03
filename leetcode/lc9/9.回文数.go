/*
 * @lc app=leetcode.cn id=9 lang=golang
 * @lcpr version=20003
 *
 * [9] 回文数
 *
 * https://leetcode.cn/problems/palindrome-number/description/
 *
 * algorithms
 * Easy (56.24%)
 * Likes:    2924
 * Dislikes: 0
 * Total Accepted:    1.7M
 * Total Submissions: 3.1M
 * Testcase Example:  '121'
 *
 * 给你一个整数 x ，如果 x 是一个回文整数，返回 true ；否则，返回 false 。
 *
 * 回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。
 *
 *
 * 例如，121 是回文，而 123 不是。
 *
 *
 *
 *
 * 示例 1：
 *
 * 输入：x = 121
 * 输出：true
 *
 *
 * 示例 2：
 *
 * 输入：x = -121
 * 输出：false
 * 解释：从左向右读, 为 -121 。 从右向左读, 为 121- 。因此它不是一个回文数。
 *
 *
 * 示例 3：
 *
 * 输入：x = 10
 * 输出：false
 * 解释：从右向左读, 为 01 。因此它不是一个回文数。
 *
 *
 *
 *
 * 提示：
 *
 *
 * -2^31 <= x <= 2^31 - 1
 *
 *
 *
 *
 * 进阶：你能不将整数转为字符串来解决这个问题吗？
 *
 */
package lc9

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	tmp, y := x, 0
	for tmp > 0 {
		y = y*10 + tmp%10
		tmp = tmp / 10
	}

	return y == x
}

// @lc code=end

/*
// @lcpr case=start
// 121\n
// @lcpr case=end

// @lcpr case=start
// -121\n
// @lcpr case=end

// @lcpr case=start
// 10\n
// @lcpr case=end

*/
