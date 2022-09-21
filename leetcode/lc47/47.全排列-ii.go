/*
 * @lc app=leetcode.cn id=47 lang=golang
 *
 * [47] 全排列 II
 *
 * https://leetcode.cn/problems/permutations-ii/description/
 *
 * algorithms
 * Medium (65.20%)
 * Likes:    1196
 * Dislikes: 0
 * Total Accepted:    379.9K
 * Total Submissions: 582.7K
 * Testcase Example:  '[1,1,2]'
 *
 * 给定一个可包含重复数字的序列 nums ，按任意顺序 返回所有不重复的全排列。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [1,1,2]
 * 输出：
 * [[1,1,2],
 * ⁠[1,2,1],
 * ⁠[2,1,1]]
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums = [1,2,3]
 * 输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= nums.length <= 8
 * -10 <= nums[i] <= 10
 *
 *
 */
package lc47

import "sort"

// @lc code=start
var (
	res   [][]int
	track []int
	used  []bool
)

func permuteUnique(nums []int) [][]int {
	res = make([][]int, 0)
	track = make([]int, 0)
	used = make([]bool, len(nums))
	sort.Ints(nums)
	backtrack(nums)
	return res
}

func backtrack(nums []int) {
	if len(track) == len(nums) {
		t := make([]int, len(track))
		copy(t, track)
		res = append(res, t)
		return
	}

	for i := 0; i < len(nums); i++ {
		if used[i] {
			continue
		}
		// 剪枝, 相邻相同的元素，!used[i-1] 表示只有在前一个元素被使用时，才继续穷举
		// 实际上使用 used[i-1] 也可以，但是剪枝效率低一些
		// if i > 0 && nums[i] == nums[i-1] && used[i-1] {
		if i > 0 && nums[i] == nums[i-1] && !used[i-1] {
			continue
		}

		track = append(track, nums[i])
		used[i] = true
		backtrack(nums)
		track = track[:len(track)-1]
		used[i] = false
	}
}

// @lc code=end
