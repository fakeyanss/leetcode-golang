/*
 * @lc app=leetcode.cn id=77 lang=golang
 * @lcpr version=20004
 *
 * [77] 组合
 *
 * https://leetcode.cn/problems/combinations/description/
 *
 * algorithms
 * Medium (77.44%)
 * Likes:    1714
 * Dislikes: 0
 * Total Accepted:    817.4K
 * Total Submissions: 1.1M
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
 * 输入：n = 1, k = 1
 * 输出：[[1]]
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= n <= 20
 * 1 <= k <= n
 *
 *
 */
package lc77

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func combine(n int, k int) [][]int {
	var res [][]int
	var track []int
	var backtrack func(start int)
	backtrack = func(start int) {
		if len(track) == k {
			tmp := append([]int{}, track...)
			res = append(res, tmp)
			return
		}
		for i := start; i <= n; i++ {
			track = append(track, i)
			backtrack(i + 1)
			track = track[:len(track)-1]
		}
	}
	backtrack(1)
	return res
}

// @lc code=end

/*
// @lcpr case=start
// 4\n2\n
// @lcpr case=end

// @lcpr case=start
// 1\n1\n
// @lcpr case=end

*/
