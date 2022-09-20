/*
 * @lc app=leetcode.cn id=46 lang=golang
 *
 * [46] 全排列
 *
 * https://leetcode.cn/problems/permutations/description/
 *
 * algorithms
 * Medium (78.75%)
 * Likes:    2220
 * Dislikes: 0
 * Total Accepted:    722.5K
 * Total Submissions: 917.4K
 * Testcase Example:  '[1,2,3]'
 *
 * 给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [1,2,3]
 * 输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums = [0,1]
 * 输出：[[0,1],[1,0]]
 *
 *
 * 示例 3：
 *
 *
 * 输入：nums = [1]
 * 输出：[[1]]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= nums.length <= 6
 * -10 <= nums[i] <= 10
 * nums 中的所有整数 互不相同
 *
 *
 */
package lc46

// @lc code=start
var (
	res   [][]int
	track []int
	used  []bool
)

func permute(nums []int) [][]int {
	res = make([][]int, 0)
	track = make([]int, 0)
	used = make([]bool, len(nums))
	backtrack(nums)
	return res
}

func backtrack(nums []int) {
	// base case
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
		track = append(track, nums[i])
		used[i] = true
		backtrack(nums)
		track = track[:len(track)-1]
		used[i] = false
	}
}

// @lc code=end
