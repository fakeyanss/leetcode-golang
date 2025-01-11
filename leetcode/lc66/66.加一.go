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

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func plusOne(digits []int) []int {
	if digits[len(digits)-1] < 9 {
		digits[len(digits)-1]++
		return digits
	}
	res := make([]int, len(digits)+1)
	res[len(res)-1] = 0
	flag := 1
	for i := len(digits) - 2; i >= 0; i-- {
		if flag == 1 {
			if digits[i] == 9 {
				res[i+1] = 0
				flag = 1
			} else {
				res[i+1] = digits[i] + 1
				flag = 0
			}
		} else {
			res[i+1] = digits[i]
		}
	}
	if flag == 1 && res[1] == 0 {
		res[0] = 1
	} else {
		res = res[1:]
	}
	return res
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
