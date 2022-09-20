/*
 * @lc app=leetcode.cn id=78 lang=golang
 *
 * [78] 子集
 *
 * https://leetcode.cn/problems/subsets/description/
 *
 * algorithms
 * Medium (80.80%)
 * Likes:    1798
 * Dislikes: 0
 * Total Accepted:    524.5K
 * Total Submissions: 649.1K
 * Testcase Example:  '[1,2,3]'
 *
 * 给你一个整数数组 nums ，数组中的元素 互不相同 。返回该数组所有可能的子集（幂集）。
 *
 * 解集 不能 包含重复的子集。你可以按 任意顺序 返回解集。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [1,2,3]
 * 输出：[[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums = [0]
 * 输出：[[],[0]]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1
 * -10
 * nums 中的所有元素 互不相同
 *
 *
 */
package lc78

// @lc code=start
var res [][]int
var track []int

func subsets(nums []int) [][]int {
	res = make([][]int, 0)
	track = make([]int, 0)
	backtrack(nums, 0)
	return res
}

func backtrack(nums []int, start int) {
	// 前序位置
	t := make([]int, len(track))
	copy(t, track)
	res = append(res, t)

	// 回溯
	for i := start; i < len(nums); i++ {
		track = append(track, nums[i])
		backtrack(nums, i+1)
		track = track[:len(track)-1]
	}
}

// @lc code=end
