/*
 * @lc app=leetcode.cn id=77 lang=golang
 *
 * [77] 组合
 *
 * https://leetcode.cn/problems/combinations/description/
 *
 * algorithms
 * Medium (77.22%)
 * Likes:    1141
 * Dislikes: 0
 * Total Accepted:    427.7K
 * Total Submissions: 553.9K
 * Testcase Example:  '4\n2'
 *
 * 给定两个整数 n 和 k，返回范围 [1, n] 中所有可能的 k 个数的组合。
 *
 * 你可以按 任何顺序 返回答案。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：n = 4, k = 2
 * 输出：
 * [
 * ⁠ [2,4],
 * ⁠ [3,4],
 * ⁠ [2,3],
 * ⁠ [1,2],
 * ⁠ [1,3],
 * ⁠ [1,4],
 * ]
 *
 * 示例 2：
 *
 *
 * 输入：n = 1, k = 1
 * 输出：[[1]]
 *
 *
 *
 * 提示：
 *
 *
 * 1
 * 1
 *
 *
 */
package lc77

// @lc code=start
var (
	res   [][]int
	track []int
)

func combine(n int, k int) [][]int {
	res = make([][]int, 0)
	track = make([]int, 0)
	backtrack(n, k, 1)
	return res
}

func backtrack(n, k, start int) {
	if len(track) == k {
		t := make([]int, len(track))
		copy(t, track)
		res = append(res, t)
		return
	}

	for i := start; i <= n; i++ {
		track = append(track, i)
		backtrack(n, k, i+1)
		track = track[:len(track)-1]
	}
}

// @lc code=end
