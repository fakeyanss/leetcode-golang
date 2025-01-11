/*
 * @lc app=leetcode.cn id=39 lang=golang
 * @lcpr version=20004
 *
 * [39] 组合总和
 *
 * https://leetcode.cn/problems/combination-sum/description/
 *
 * algorithms
 * Medium (73.28%)
 * Likes:    2938
 * Dislikes: 0
 * Total Accepted:    1.1M
 * Total Submissions: 1.5M
 * Testcase Example:  '[2,3,6,7]\n7'
 *
 * 给你一个 无重复元素 的整数数组 candidates 和一个目标整数 target ，找出 candidates 中可以使数字和为目标数 target
 * 的 所有 不同组合 ，并以列表形式返回。你可以按 任意顺序 返回这些组合。
 *
 * candidates 中的 同一个 数字可以 无限制重复被选取 。如果至少一个数字的被选数量不同，则两种组合是不同的。
 *
 * 对于给定的输入，保证和为 target 的不同组合数少于 150 个。
 *
 *
 *
 * 示例 1：
 *
 * 输入：candidates = [2,3,6,7], target = 7
 * 输出：[[2,2,3],[7]]
 * 解释：
 * 2 和 3 可以形成一组候选，2 + 2 + 3 = 7 。注意 2 可以使用多次。
 * 7 也是一个候选， 7 = 7 。
 * 仅有这两种组合。
 *
 * 示例 2：
 *
 * 输入: candidates = [2,3,5], target = 8
 * 输出: [[2,2,2,2],[2,3,3],[3,5]]
 *
 * 示例 3：
 *
 * 输入: candidates = [2], target = 1
 * 输出: []
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= candidates.length <= 30
 * 2 <= candidates[i] <= 40
 * candidates 的所有元素 互不相同
 * 1 <= target <= 40
 *
 *
 */
package lc39

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func combinationSum(candidates []int, target int) [][]int {
	var res [][]int
	var track []int
	var trackSum int
	var backtrack func(start int)
	backtrack = func(start int) {
		if trackSum > target {
			return
		}
		if trackSum == target {
			tmp := append([]int{}, track...)
			res = append(res, tmp)
		}
		for i := start; i < len(candidates); i++ {
			trackSum += candidates[i]
			track = append(track, candidates[i])
			backtrack(i) // 元素可重复选，回溯下一个仍然是当前下标
			track = track[:len(track)-1]
			trackSum -= candidates[i]
		}
	}

	backtrack(0)

	return res

	// 完全背包，dp[i]的含义从容量为i的背包有几种装法变成具体装什么
	// dp := make([][][]int, target+1)
	// dp[0] = [][]int{{}}
	// for i := 0; i < len(candidates); i++ {
	// 	for j := candidates[i]; j < len(dp); j++ {
	// 		for _, group := range dp[j-candidates[i]] {
	// 			newGroup := append([]int{}, group...)
	// 			newGroup = append(newGroup, candidates[i])
	// 			dp[j] = append(dp[j], newGroup)
	// 		}
	// 	}
	// }
	// return dp[target]
}

// @lc code=end

/*
// @lcpr case=start
// [2,3,6,7]\n7\n
// @lcpr case=end

// @lcpr case=start
// [2,3,5]\n8\n
// @lcpr case=end

// @lcpr case=start
// [2]\n1\n
// @lcpr case=end

*/
