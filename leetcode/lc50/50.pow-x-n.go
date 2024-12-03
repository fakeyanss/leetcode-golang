/*
 * @lc app=leetcode.cn id=50 lang=golang
 * @lcpr version=20002
 *
 * [50] Pow(x, n)
 *
 * https://leetcode.cn/problems/powx-n/description/
 *
 * algorithms
 * Medium (38.60%)
 * Likes:    1405
 * Dislikes: 0
 * Total Accepted:    487.4K
 * Total Submissions: 1.3M
 * Testcase Example:  '2.00000\n10'
 *
 * 实现 pow(x, n) ，即计算 x 的整数 n 次幂函数（即，x^n^ ）。
 *
 *
 *
 * 示例 1：
 *
 * 输入：x = 2.00000, n = 10
 * 输出：1024.00000
 *
 *
 * 示例 2：
 *
 * 输入：x = 2.10000, n = 3
 * 输出：9.26100
 *
 *
 * 示例 3：
 *
 * 输入：x = 2.00000, n = -2
 * 输出：0.25000
 * 解释：2^-2 = 1/2^2 = 1/4 = 0.25
 *
 *
 *
 *
 * 提示：
 *
 *
 * -100.0 < x < 100.0
 * -2^31 <= n <= 2^31-1
 * n 是一个整数
 * 要么 x 不为零，要么 n > 0 。
 * -10^4 <= x^n <= 10^4
 *
 *
 */

// @lcpr-template-start

// @lcpr-template-end
package lc50

// @lc code=start
func myPow(x float64, n int) float64 {
	var quickPow func(a float64, k int) float64
	quickPow = func(a float64, k int) float64 {
		if k == 0 {
			return 1.0
		}
		y := quickPow(a, k/2)
		if k%2 == 0 {
			return y * y
		}
		return y * y * a
	}

	if n >= 0 {
		return quickPow(x, n)
	}
	return 1.0 / quickPow(x, -n)
}

// @lc code=end

/*
// @lcpr case=start
// 2.00000\n10\n
// @lcpr case=end

// @lcpr case=start
// 2.10000\n3\n
// @lcpr case=end

// @lcpr case=start
// 2.00000\n-2\n
// @lcpr case=end

*/
