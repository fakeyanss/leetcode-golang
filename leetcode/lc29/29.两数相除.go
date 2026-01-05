/*
 * @lc app=leetcode.cn id=29 lang=golang
 * @lcpr version=30305
 *
 * [29] 两数相除
 *
 * https://leetcode.cn/problems/divide-two-integers/description/
 *
 * algorithms
 * Medium (22.50%)
 * Likes:    1323
 * Dislikes: 0
 * Total Accepted:    270.3K
 * Total Submissions: 1.2M
 * Testcase Example:  '10\n3\n7\n-3'
 *
 * 给你两个整数，被除数 dividend 和除数 divisor。将两数相除，要求 不使用 乘法、除法和取余运算。
 *
 * 整数除法应该向零截断，也就是截去（truncate）其小数部分。例如，8.345 将被截断为 8 ，-2.7335 将被截断至 -2 。
 *
 * 返回被除数 dividend 除以除数 divisor 得到的 商 。
 *
 * 注意：假设我们的环境只能存储 32 位 有符号整数，其数值范围是 [−2^31,  2^31 − 1] 。本题中，如果商 严格大于 2^31 − 1
 * ，则返回 2^31 − 1 ；如果商 严格小于 -2^31 ，则返回 -2^31^ 。
 *
 *
 *
 * 示例 1:
 *
 * 输入: dividend = 10, divisor = 3
 * 输出: 3
 * 解释: 10/3 = 3.33333.. ，向零截断后得到 3 。
 *
 * 示例 2:
 *
 * 输入: dividend = 7, divisor = -3
 * 输出: -2
 * 解释: 7/-3 = -2.33333.. ，向零截断后得到 -2 。
 *
 *
 *
 * 提示：
 *
 *
 * -2^31 <= dividend, divisor <= 2^31 - 1
 * divisor != 0
 *
 *
 */
package lc29

// @lc code=start
func divide(dividend int, divisor int) int {
	if dividend == 0 {
		return 0
	}
	// 唯一溢出情况，-2^31 / -1 = 2^31（超出int范围）
	if dividend == -1<<31 && divisor == -1 {
		return 1<<31 - 1
	}
	if divisor == -1 {
		return -dividend
	}
	if divisor == 1 {
		return dividend
	}

	// 记录结果符号
	sign := 1
	if (dividend < 0) != (divisor < 0) {
		sign = -1
	}

	// 转换为负数，避免正数2^31溢出
	if dividend > 0 {
		dividend = -dividend
	}
	if divisor > 0 {
		divisor = -divisor
	}

	// 倍增快速除法
	var quotient int // 商
	limit := -1 << 31
	limit >>= 1 // MIN的一半
	for dividend <= divisor {
		tmp, cnt := divisor, 1
		// tmp>=limit避免倍增后溢出MIN，tmp+tmp>=dividend，避免倍增后tmp绝对值大于被除数绝对值
		for tmp >= limit && tmp+tmp >= dividend {
			tmp += tmp
			cnt += cnt
		}
		dividend -= tmp
		quotient += cnt
	}

	if sign == 1 {
		return quotient
	}
	return -quotient
}

// @lc code=end

/*
// @lcpr case=start
// 10\n3\n
// @lcpr case=end

// @lcpr case=start
// 7\n-3\n
// @lcpr case=end

*/
