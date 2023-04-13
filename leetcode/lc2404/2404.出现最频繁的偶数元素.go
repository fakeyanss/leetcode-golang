/*
 * @lc app=leetcode.cn id=2404 lang=golang
 *
 * [2404] 出现最频繁的偶数元素
 *
 * https://leetcode.cn/problems/most-frequent-even-element/description/
 *
 * algorithms
 * Easy (58.45%)
 * Likes:    56
 * Dislikes: 0
 * Total Accepted:    32.9K
 * Total Submissions: 56.4K
 * Testcase Example:  '[0,1,2,2,4,4,1]'
 *
 * 给你一个整数数组 nums ，返回出现最频繁的偶数元素。
 *
 * 如果存在多个满足条件的元素，只需要返回 最小 的一个。如果不存在这样的元素，返回 -1 。
 *
 *
 *
 * 示例 1：
 *
 * 输入：nums = [0,1,2,2,4,4,1]
 * 输出：2
 * 解释：
 * 数组中的偶数元素为 0、2 和 4 ，在这些元素中，2 和 4 出现次数最多。
 * 返回最小的那个，即返回 2 。
 *
 * 示例 2：
 *
 * 输入：nums = [4,4,4,9,2,4]
 * 输出：4
 * 解释：4 是出现最频繁的偶数元素。
 *
 *
 * 示例 3：
 *
 * 输入：nums = [29,47,21,41,13,37,25,7]
 * 输出：-1
 * 解释：不存在偶数元素。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= nums.length <= 2000
 * 0 <= nums[i] <= 10^5
 *
 *
 */
package lc2404

// @lc code=start
// hash表计数
func mostFrequentEven(nums []int) int {
	freq := make(map[int]int)
	max, ans := 0, -1
	for _, v := range nums {
		if v&1 == 1 {
			continue
		}
		freq[v] += 1
		if max < freq[v] {
			max = freq[v]
			ans = v
		} else if max == freq[v] && ans > v {
			ans = v
		}
	}
	return ans
}

// @lc code=end
