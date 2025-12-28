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
func subsets(nums []int) [][]int {
	// 对每个nums[i]，都考虑选或不选
	n := len(nums)
	var res [][]int
	track := make([]int, 0, n)

	var dfs func(int)
	dfs = func(i int) {
		if i == n {
			res = append(res, append([]int{}, track...))
			return
		}

		dfs(i + 1) // 不选nums[i]

		track = append(track, nums[i]) // 选nums[i]
		dfs(i + 1)
		track = track[:len(track)-1]
	}
	dfs(0)
	return res
}

// @lc code=end
