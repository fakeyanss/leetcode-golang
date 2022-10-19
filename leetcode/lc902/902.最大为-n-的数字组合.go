/*
 * @lc app=leetcode.cn id=902 lang=golang
 *
 * [902] 最大为 N 的数字组合
 *
 * https://leetcode.cn/problems/numbers-at-most-n-given-digit-set/description/
 *
 * algorithms
 * Hard (45.80%)
 * Likes:    223
 * Dislikes: 0
 * Total Accepted:    22.2K
 * Total Submissions: 48.5K
 * Testcase Example:  '["1","3","5","7"]\n100'
 *
 * 给定一个按 非递减顺序 排列的数字数组 digits 。你可以用任意次数 digits[i] 来写的数字。例如，如果 digits =
 * ['1','3','5']，我们可以写数字，如 '13', '551', 和 '1351315'。
 *
 * 返回 可以生成的小于或等于给定整数 n 的正整数的个数 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：digits = ["1","3","5","7"], n = 100
 * 输出：20
 * 解释：
 * 可写出的 20 个数字是：
 * 1, 3, 5, 7, 11, 13, 15, 17, 31, 33, 35, 37, 51, 53, 55, 57, 71, 73, 75,
 * 77.
 *
 *
 * 示例 2：
 *
 *
 * 输入：digits = ["1","4","9"], n = 1000000000
 * 输出：29523
 * 解释：
 * 我们可以写 3 个一位数字，9 个两位数字，27 个三位数字，
 * 81 个四位数字，243 个五位数字，729 个六位数字，
 * 2187 个七位数字，6561 个八位数字和 19683 个九位数字。
 * 总共，可以使用D中的数字写出 29523 个整数。
 *
 * 示例 3:
 *
 *
 * 输入：digits = ["7"], n = 8
 * 输出：1
 *
 *
 *
 *
 * 提示：
 *
 *
 *
 * 1 <= digits.length <= 9
 * digits[i].length == 1
 * digits[i] 是从 '1' 到 '9' 的数
 * digits 中的所有值都 不同
 * digits 按 非递减顺序 排列
 * 1 <= n <= 10^9
 *
 *
 */
package lc902

import "strconv"

// @lc code=start
func atMostNGivenDigitSet(digits []string, n int) int {
	s := strconv.Itoa(n)
	m := len(s)
	dp := make([]int, m)
	for i := range dp {
		dp[i] = -1 // dp[i] = -1 表示 i 这个状态还没被计算出来
	}
	var f func(int, bool, bool) int
	f = func(i int, isLimit, isNum bool) (res int) {
		if i == m {
			if isNum { // 如果填了数字，则为 1 种合法方案
				return 1
			}
			return
		}
		if !isLimit && isNum { // 在不受到任何约束的情况下，返回记录的结果，避免重复运算
			dv := &dp[i]
			if *dv >= 0 {
				return *dv
			}
			defer func() { *dv = res }()
		}
		if !isNum { // 前面不填数字，那么可以跳过当前数位，也不填数字
			// isLimit 改为 false，因为没有填数字，位数都比 n 要短，自然不会受到 n 的约束
			// isNum 仍然为 false，因为没有填任何数字
			res += f(i+1, false, false)
		}
		// 根据是否受到约束，决定可以填的数字的上限
		up := byte('9')
		if isLimit {
			up = s[i]
		}
		// 注意：对于一般的题目而言，如果此时 isNum 为 false，则必须从 1 开始枚举，由于本题 digits 没有 0，所以无需处理这种情况
		for _, d := range digits { // 枚举要填入的数字 d
			if d[0] > up { // d 超过上限，由于 digits 是有序的，后面的 d 都会超过上限，故退出循环
				break
			}
			// isLimit：如果当前受到 n 的约束，且填的数字等于上限，那么后面仍然会受到 n 的约束
			// isNum 为 true，因为填了数字
			res += f(i+1, isLimit && d[0] == up, true)
		}
		return
	}
	return f(0, true, false)
}

// @lc code=end
