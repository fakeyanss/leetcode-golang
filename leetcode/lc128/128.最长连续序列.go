/*
 * @lc app=leetcode.cn id=128 lang=golang
 * @lcpr version=20004
 *
 * [128] 最长连续序列
 *
 * https://leetcode.cn/problems/longest-consecutive-sequence/description/
 *
 * algorithms
 * Medium (51.42%)
 * Likes:    2262
 * Dislikes: 0
 * Total Accepted:    823.2K
 * Total Submissions: 1.6M
 * Testcase Example:  '[100,4,200,1,3,2]'
 *
 * 给定一个未排序的整数数组 nums ，找出数字连续的最长序列（不要求序列元素在原数组中连续）的长度。
 *
 * 请你设计并实现时间复杂度为 O(n) 的算法解决此问题。
 *
 *
 *
 * 示例 1：
 *
 * 输入：nums = [100,4,200,1,3,2]
 * 输出：4
 * 解释：最长数字连续序列是 [1, 2, 3, 4]。它的长度为 4。
 *
 * 示例 2：
 *
 * 输入：nums = [0,3,7,2,5,8,4,6,0,1]
 * 输出：9
 *
 *
 *
 *
 * 提示：
 *
 *
 * 0 <= nums.length <= 10^5
 * -10^9 <= nums[i] <= 10^9
 *
 *
 */
package lc128

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func longestConsecutive(nums []int) int {
	res := 0
	numsMap := make(map[int]bool, len(nums))
	for _, v := range nums {
		numsMap[v] = true
	}
	for v := range numsMap {
		if numsMap[v-1] {
			continue
		}
		curNum, curLen := v, 1
		for numsMap[curNum+1] {
			curNum++
			curLen++
		}
		res = max(res, curLen)
	}
	return res
}

// @lc code=end

/*
// @lcpr case=start
// [100,4,200,1,3,2]\n
// @lcpr case=end

// @lcpr case=start
// [0,3,7,2,5,8,4,6,0,1]\n
// @lcpr case=end

*/
