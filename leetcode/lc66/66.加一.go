/*
 * @lc app=leetcode.cn id=66 lang=golang
 * @lcpr version=20004
 *
 * [66] 加一
 *
 * https://leetcode.cn/problems/plus-one/description/
 *
 * algorithms
 * Easy (46.37%)
 * Likes:    1452
 * Dislikes: 0
 * Total Accepted:    800.8K
 * Total Submissions: 1.7M
 * Testcase Example:  '[1,2,3]'
 *
 * 给定一个由 整数 组成的 非空 数组所表示的非负整数，在该数的基础上加一。
 *
 * 最高位数字存放在数组的首位， 数组中每个元素只存储单个数字。
 *
 * 你可以假设除了整数 0 之外，这个整数不会以零开头。
 *
 *
 *
 * 示例 1：
 *
 * 输入：digits = [1,2,3]
 * 输出：[1,2,4]
 * 解释：输入数组表示数字 123。
 *
 *
 * 示例 2：
 *
 * 输入：digits = [4,3,2,1]
 * 输出：[4,3,2,2]
 * 解释：输入数组表示数字 4321。
 *
 *
 * 示例 3：
 *
 * 输入：digits = [9]
 * 输出：[1,0]
 * 解释：输入数组表示数字 9。
 * 加 1 得到了 9 + 1 = 10。
 * 因此，结果应该是 [1,0]。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= digits.length <= 100
 * 0 <= digits[i] <= 9
 *
 *
 */
package lc66

import "slices"

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func plusOne(digits []int) []int {
	for i, d := range slices.Backward(digits) {
		if d < 9 {
			digits[i]++ // 找到右侧第一个小于9的数，进位
			return digits
		}
		digits[i] = 0 // 将进位数字右边的9都变成0
	}
	// 如果digits全是9，加1后多了1位，首位是0后面全是0
	digits = append(digits, 0)
	digits[0] = 1
	return digits
}

// @lc code=end

/*
// @lcpr case=start
// [1,2,3]\n
// @lcpr case=end

// @lcpr case=start
// [4,3,2,1]\n
// @lcpr case=end

// @lcpr case=start
// [9]\n
// @lcpr case=end

*/
