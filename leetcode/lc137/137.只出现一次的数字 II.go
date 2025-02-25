/*
 * @lc app=leetcode.cn id=137 lang=golang
 * @lcpr version=20004
 *
 * [137] 只出现一次的数字 II
 *
 * https://leetcode.cn/problems/single-number-ii/description/
 *
 * algorithms
 * Medium (72.34%)
 * Likes:    1275
 * Dislikes: 0
 * Total Accepted:    228.2K
 * Total Submissions: 315.4K
 * Testcase Example:  '[2,2,3,2]'
 *
 * 给你一个整数数组 nums ，除某个元素仅出现 一次 外，其余每个元素都恰出现 三次 。请你找出并返回那个只出现了一次的元素。
 *
 * 你必须设计并实现线性时间复杂度的算法且使用常数级空间来解决此问题。
 *
 *
 *
 * 示例 1：
 *
 * 输入：nums = [2,2,3,2]
 * 输出：3
 *
 *
 * 示例 2：
 *
 * 输入：nums = [0,1,0,1,0,1,99]
 * 输出：99
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= nums.length <= 3 * 10^4
 * -2^31 <= nums[i] <= 2^31 - 1
 * nums 中，除某个元素仅出现 一次 外，其余每个元素都恰出现 三次
 *
 *
 */
package lc137

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func singleNumber(nums []int) int {
	// a, b := 0, 0
	// for _, x := range nums {
	// 	a = (a ^ x) & ^b
	// 	b = (b ^ x) & ^a
	// }
	// return a

	res := int32(0)
	for i := 0; i < 32; i++ {
		cnt1 := int32(0)
		for _, x := range nums {
			cnt1 += int32(x) >> i & 1 // 取出第i位的bit相加
		}
		res |= cnt1 % 3 << i // 模3取余结果就是目标数字在第i位的bit
	}
	return int(res)
}

// @lc code=end

/*
// @lcpr case=start
// [2,2,3,2]\n
// @lcpr case=end

// @lcpr case=start
// [0,1,0,1,0,1,99]\n
// @lcpr case=end

*/
