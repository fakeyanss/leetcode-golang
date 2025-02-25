/*
 * @lc app=leetcode.cn id=172 lang=golang
 * @lcpr version=20004
 *
 * [172] 阶乘后的零
 *
 * https://leetcode.cn/problems/factorial-trailing-zeroes/description/
 *
 * algorithms
 * Medium (50.61%)
 * Likes:    862
 * Dislikes: 0
 * Total Accepted:    219.3K
 * Total Submissions: 433K
 * Testcase Example:  '3'
 *
 * 给定一个整数 n ，返回 n! 结果中尾随零的数量。
 *
 * 提示 n! = n * (n - 1) * (n - 2) * ... * 3 * 2 * 1
 *
 *
 *
 * 示例 1：
 *
 * 输入：n = 3
 * 输出：0
 * 解释：3! = 6 ，不含尾随 0
 *
 *
 * 示例 2：
 *
 * 输入：n = 5
 * 输出：1
 * 解释：5! = 120 ，有一个尾随 0
 *
 *
 * 示例 3：
 *
 * 输入：n = 0
 * 输出：0
 *
 *
 *
 *
 * 提示：
 *
 *
 * 0 <= n <= 10^4
 *
 *
 *
 *
 * 进阶：你可以设计并实现对数时间复杂度的算法来解决此问题吗？
 *
 */
package lc172

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func trailingZeroes(n int) int {
	// 两个数相乘结果末尾有 0，一定是因为两个数中有因子 2 和 5.
	// 而最多可以分解出多少个因子 2 和 5，主要取决于能分解出几个因子 5，因为每个偶数都能分解出因子 2，因子 2 肯定比因子 5 多得多
	var res int
	divisor := 5
	for divisor <= n {
		res += n / divisor
		divisor *= 5
	}
	return res
}

// @lc code=end

/*
// @lcpr case=start
// 3\n
// @lcpr case=end

// @lcpr case=start
// 5\n
// @lcpr case=end

// @lcpr case=start
// 0\n
// @lcpr case=end

*/
