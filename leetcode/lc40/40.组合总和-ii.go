/*
 * @lc app=leetcode.cn id=40 lang=golang
 *
 * [40] 组合总和 II
 *
 * https://leetcode.cn/problems/combination-sum-ii/description/
 *
 * algorithms
 * Medium (60.42%)
 * Likes:    1110
 * Dislikes: 0
 * Total Accepted:    351.4K
 * Total Submissions: 581.6K
 * Testcase Example:  '[10,1,2,7,6,1,5]\n8'
 *
 * 给定一个候选人编号的集合 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。
 *
 * candidates 中的每个数字在每个组合中只能使用 一次 。
 *
 * 注意：解集不能包含重复的组合。
 *
 *
 *
 * 示例 1:
 *
 *
 * 输入: candidates = [10,1,2,7,6,1,5], target = 8,
 * 输出:
 * [
 * [1,1,6],
 * [1,2,5],
 * [1,7],
 * [2,6]
 * ]
 *
 * 示例 2:
 *
 *
 * 输入: candidates = [2,5,2,1,2], target = 5,
 * 输出:
 * [
 * [1,2,2],
 * [5]
 * ]
 *
 *
 *
 * 提示:
 *
 *
 * 1 <= candidates.length <= 100
 * 1 <= candidates[i] <= 50
 * 1 <= target <= 30
 *
 *
 */
package lc40

import "sort"

// @lc code=start

func combinationSum2(candidates []int, target int) [][]int {
	var res [][]int
	var track []int
	var trackSum int
	var backtrack func(start int)
	sort.Ints(candidates)
	backtrack = func(start int) {
		if trackSum > target {
			return
		}
		if trackSum == target {
			tmp := append([]int{}, track...)
			res = append(res, tmp)
		}
		for i := start; i < len(candidates); i++ {
			if i != start && candidates[i] == candidates[i-1] {
				continue
			}
			trackSum += candidates[i]
			track = append(track, candidates[i])
			backtrack(i + 1) // 元素不可重复选，回溯下一个仍然是当前下标
			track = track[:len(track)-1]
			trackSum -= candidates[i]
		}
	}

	backtrack(0)

	return res
}

// @lc code=end
