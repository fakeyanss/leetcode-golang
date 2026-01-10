/*
 * @lc app=leetcode.cn id=238 lang=golang
 * @lcpr version=30305
 *
 * [238] 除自身以外数组的乘积
 *
 * https://leetcode.cn/problems/product-of-array-except-self/description/
 *
 * algorithms
 * Medium (77.94%)
 * Likes:    2161
 * Dislikes: 0
 * Total Accepted:    873.1K
 * Total Submissions: 1.1M
 * Testcase Example:  '[1,2,3,4]\n[-1,1,0,-3,3]'
 *
 * 给你一个整数数组 nums，返回 数组 answer ，其中 answer[i] 等于 nums 中除了 nums[i] 之外其余各元素的乘积 。
 *
 * 题目数据 保证 数组 nums之中任意元素的全部前缀元素和后缀的乘积都在  32 位 整数范围内。
 *
 * 请 不要使用除法，且在 O(n) 时间复杂度内完成此题。
 *
 *
 *
 * 示例 1:
 *
 * 输入: nums = [1,2,3,4]
 * 输出: [24,12,8,6]
 *
 *
 * 示例 2:
 *
 * 输入: nums = [-1,1,0,-3,3]
 * 输出: [0,0,9,0,0]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 2 <= nums.length <= 10^5
 * -30 <= nums[i] <= 30
 * 输入 保证 数组 answer[i] 在  32 位 整数范围内
 *
 *
 *
 *
 * 进阶：你可以在 O(1) 的额外空间复杂度内完成这个题目吗？（ 出于对空间复杂度分析的目的，输出数组 不被视为 额外空间。）
 *
 */
package lc238

// @lc code=start
// 思路：前缀和
func productExceptSelf(nums []int) []int {
	n := len(nums)
	// premul[i]记录nums[i]左边所有数的成绩
	premul := make([]int, n)
	premul[0] = 1
	for i := 1; i < n; i++ {
		premul[i] = premul[i-1] * nums[i-1]
	}
	// postmul[i]记录nums[i]右边所有数的成绩
	postmul := make([]int, n)
	postmul[n-1] = 1
	for i := n - 2; i >= 0; i-- {
		postmul[i] = postmul[i+1] * nums[i+1]
	}

	res := make([]int, n)
	for i := range n {
		res[i] = premul[i] * postmul[i]
	}
	return res
}

// @lc code=end

/*
// @lcpr case=start
// [1,2,3,4]\n
// @lcpr case=end

// @lcpr case=start
// [-1,1,0,-3,3]\n
// @lcpr case=end

*/
