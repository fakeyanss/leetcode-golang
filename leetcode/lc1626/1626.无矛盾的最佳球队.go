/*
 * @lc app=leetcode.cn id=1626 lang=golang
 *
 * [1626] 无矛盾的最佳球队
 *
 * https://leetcode.cn/problems/best-team-with-no-conflicts/description/
 *
 * algorithms
 * Medium (52.81%)
 * Likes:    164
 * Dislikes: 0
 * Total Accepted:    21.2K
 * Total Submissions: 40.1K
 * Testcase Example:  '[1,3,5,10,15]\n[1,2,3,4,5]'
 *
 * 假设你是球队的经理。对于即将到来的锦标赛，你想组合一支总体得分最高的球队。球队的得分是球队中所有球员的分数 总和 。
 *
 * 然而，球队中的矛盾会限制球员的发挥，所以必须选出一支 没有矛盾 的球队。如果一名年龄较小球员的分数 严格大于
 * 一名年龄较大的球员，则存在矛盾。同龄球员之间不会发生矛盾。
 *
 * 给你两个列表 scores 和 ages，其中每组 scores[i] 和 ages[i] 表示第 i 名球员的分数和年龄。请你返回
 * 所有可能的无矛盾球队中得分最高那支的分数 。
 *
 *
 *
 * 示例 1：
 *
 * 输入：scores = [1,3,5,10,15], ages = [1,2,3,4,5]
 * 输出：34
 * 解释：你可以选中所有球员。
 *
 * 示例 2：
 *
 * 输入：scores = [4,5,6,5], ages = [2,1,2,1]
 * 输出：16
 * 解释：最佳的选择是后 3 名球员。注意，你可以选中多个同龄球员。
 *
 *
 * 示例 3：
 *
 * 输入：scores = [1,2,3,5], ages = [8,9,10,1]
 * 输出：6
 * 解释：最佳的选择是前 3 名球员。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= scores.length, ages.length <= 1000
 * scores.length == ages.length
 * 1 <= scores[i] <= 10^6
 * 1 <= ages[i] <= 1000
 *
 *
 */
package lc1626

import "sort"

// @lc code=start
func bestTeamScore(scores []int, ages []int) int {
	type pair struct {
		score int
		age   int
	}
	n := len(scores)
	p := make([]pair, n)
	for i, v := range scores {
		p[i] = pair{
			score: v,
			age:   ages[i],
		}
	}
	// 按分数排序，其次按年龄排序
	sort.Slice(p, func(i, j int) bool {
		return p[i].score < p[j].score || p[i].score == p[j].score && p[i].age < p[j].age
	})
	// 定义dp[i]为以p[i]结尾的球队的最大得分
	// 状态转移方程：dp[i] = max(dp[j] + p[i].score), j<i且p[j].age<p[i].age
	// 初始dp[0] = 0
	dp := make([]int, n)
	ans := 0
	for i, m := range p {
		for j, n := range p[:i] {
			if n.age <= m.age {
				dp[i] = max(dp[i], dp[j])
			}
		}
		dp[i] += m.score
		ans = max(ans, dp[i])
	}
	return ans
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

// @lc code=end
