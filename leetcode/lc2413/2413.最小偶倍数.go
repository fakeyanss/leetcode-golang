/*
 * @lc app=leetcode.cn id=2413 lang=golang
 *
 * [2413] 最小偶倍数
 *
 * https://leetcode.cn/problems/smallest-even-multiple/description/
 *
 * algorithms
 * Easy (91.57%)
 * Likes:    5
 * Dislikes: 0
 * Total Accepted:    8.8K
 * Total Submissions: 9.7K
 * Testcase Example:  '5'
 *
 * 给你一个正整数 n ，返回 2 和 n 的最小公倍数（正整数）。
 *
 *
 * 示例 1：
 *
 * 输入：n = 5
 * 输出：10
 * 解释：5 和 2 的最小公倍数是 10 。
 *
 *
 * 示例 2：
 *
 * 输入：n = 6
 * 输出：6
 * 解释：6 和 2 的最小公倍数是 6 。注意数字会是它自身的倍数。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= n <= 150
 *
 *
 */
package lc2413

// @lc code=start
func smallestEvenMultiple(n int) int {
	return lcm(n, 2)
}

// 最小公倍数 = 两数乘积 / 两数的最大公约数(gcd)
func lcm(x, y int) int {
	return x * y / gcd(x, y)
}

func gcd(x, y int) int {
	if y == 0 {
		return x
	}
	return gcd(y, x%y)
}

// @lc code=end
