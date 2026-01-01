/*
 * @lc app=leetcode.cn id=152 lang=golang
 * @lcpr version=30305
 *
 * [152] 乘积最大子数组
 *
 * https://leetcode.cn/problems/maximum-product-subarray/description/
 *
 * algorithms
 * Medium (43.71%)
 * Likes:    2515
 * Dislikes: 0
 * Total Accepted:    670.5K
 * Total Submissions: 1.5M
 * Testcase Example:  '[2,3,-2,4]\n[-2,0,-1]'
 *
 * 给你一个整数数组 nums ，请你找出数组中乘积最大的非空连续 子数组（该子数组中至少包含一个数字），并返回该子数组所对应的乘积。
 *
 * 测试用例的答案是一个 32-位 整数。
 *
 * 请注意，一个只包含一个元素的数组的乘积是这个元素的值。
 *
 *
 *
 * 示例 1:
 *
 * 输入: nums = [2,3,-2,4]
 * 输出: 6
 * 解释: 子数组 [2,3] 有最大乘积 6。
 *
 *
 * 示例 2:
 *
 * 输入: nums = [-2,0,-1]
 * 输出: 0
 * 解释: 结果不能为 2, 因为 [-2,-1] 不是子数组。
 *
 *
 *
 * 提示:
 *
 *
 * 1 <= nums.length <= 2 * 10^4
 * -10 <= nums[i] <= 10
 * nums 的任何子数组的乘积都 保证 是一个 32-位 整数
 *
 *
 */
package lc152

// @lc code=start
func maxProduct(nums []int) int {
	// fmax[i],fmin[i]分别表示以nums[i]结尾的子数组乘积最大值、最小值
	// 因为存在负数，所以当前乘积最大值，可能依赖前一个状态的最小值
	n := len(nums)
	fmax, fmin := make([]int, n), make([]int, n)
	fmax[0], fmin[0] = nums[0], nums[0]
	res := nums[0]
	for i := 1; i < n; i++ {
		fmax[i] = max(nums[i], fmax[i-1]*nums[i], fmin[i-1]*nums[i])
		fmin[i] = min(nums[i], fmax[i-1]*nums[i], fmin[i-1]*nums[i])
		res = max(res, fmax[i])
	}
	return res
}

// @lc code=end

/*
// @lcpr case=start
// [2,3,-2,4]\n
// @lcpr case=end

// @lcpr case=start
// [-2,0,-1]\n
// @lcpr case=end

*/
