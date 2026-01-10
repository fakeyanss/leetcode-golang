/*
 * @lc app=leetcode.cn id=50 lang=golang
 * @lcpr version=20004
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
package lc50

// @lcpr-template-end
// @lc code=start
// 思路：快速幂
// 数学原理：
// 1. 任何正整数n可唯一分解为2的幂次之和（二进制）：
// n = a0×2^0 + a1×2^1 + ... + ak×2^k（ai∈{0,1}）
// 2. 幂的乘法法则：
// x^n = x^(a0×2^0 + a1×2^1 + ... + ak×2^k) = x^(a0×2^0) × x^(a1×2^1) × ... × x^(ak×2^k)
// 3. 当ai=0时，x^(ai×2^i)=x^0=1（乘1不影响结果，可忽略）
// 4. 当ai=1时，x^(ai×2^i)=x^(2^i)（需乘入结果）
func myPow(x float64, n int) float64 {
	res := 1.0
	if n < 0 { // x^-n=(1/x)^n
		n = -n
		x = 1 / x
	}
	for n > 0 { // 从低到高枚举n的每个比特位，即找到ai是0还是1
		if n&1 == 1 {
			res *= x // 这个比特位是1，把累积的x乘到res中
		}
		x *= x  // x累积
		n >>= 1 // n右移除以2
	}
	return res
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
