/*
 * @lc app=leetcode.cn id=560 lang=golang
 * @lcpr version=30305
 *
 * [560] 和为 K 的子数组
 *
 * https://leetcode.cn/problems/subarray-sum-equals-k/description/
 *
 * algorithms
 * Medium (45.89%)
 * Likes:    2983
 * Dislikes: 0
 * Total Accepted:    935K
 * Total Submissions: 2M
 * Testcase Example:  '[1,1,1]\n2\n[1,2,3]\n3'
 *
 * 给你一个整数数组 nums 和一个整数 k ，请你统计并返回 该数组中和为 k 的子数组的个数 。
 *
 * 子数组是数组中元素的连续非空序列。
 *
 *
 *
 * 示例 1：
 *
 * 输入：nums = [1,1,1], k = 2
 * 输出：2
 *
 *
 * 示例 2：
 *
 * 输入：nums = [1,2,3], k = 3
 * 输出：2
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= nums.length <= 2 * 10^4
 * -1000 <= nums[i] <= 1000
 * -10^7 <= k <= 10^7
 *
 *
 */
package lc560

// @lc code=start
func subarraySum(nums []int, k int) int {
	presum := map[int]int{0: 1}
	res, cursum := 0, 0
	for _, v := range nums {
		cursum += v
		res += presum[cursum-k]
		presum[cursum]++
	}
	return res
}

// @lc code=end

/*
// @lcpr case=start
// [1,1,1]\n2\n
// @lcpr case=end

// @lcpr case=start
// [1,2,3]\n3\n
// @lcpr case=end

*/
