/*
 * @lc app=leetcode.cn id=486 lang=golang
 *
 * [486] 预测赢家
 *
 * https://leetcode.cn/problems/predict-the-winner/description/
 *
 * algorithms
 * Medium (59.28%)
 * Likes:    580
 * Dislikes: 0
 * Total Accepted:    60.7K
 * Total Submissions: 102.4K
 * Testcase Example:  '[1,5,2]'
 *
 * 给你一个整数数组 nums 。玩家 1 和玩家 2 基于这个数组设计了一个游戏。
 *
 * 玩家 1 和玩家 2 轮流进行自己的回合，玩家 1 先手。开始时，两个玩家的初始分值都是 0
 * 。每一回合，玩家从数组的任意一端取一个数字（即，nums[0] 或 nums[nums.length - 1]），取到的数字将会从数组中移除（数组长度减
 * 1 ）。玩家选中的数字将会加到他的得分上。当数组中没有剩余数字可取时，游戏结束。
 *
 * 如果玩家 1 能成为赢家，返回 true 。如果两个玩家得分相等，同样认为玩家 1 是游戏的赢家，也返回 true
 * 。你可以假设每个玩家的玩法都会使他的分数最大化。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [1,5,2]
 * 输出：false
 * 解释：一开始，玩家 1 可以从 1 和 2 中进行选择。
 * 如果他选择 2（或者 1 ），那么玩家 2 可以从 1（或者 2 ）和 5 中进行选择。如果玩家 2 选择了 5 ，那么玩家 1 则只剩下 1（或者 2
 * ）可选。
 * 所以，玩家 1 的最终分数为 1 + 2 = 3，而玩家 2 为 5 。
 * 因此，玩家 1 永远不会成为赢家，返回 false 。
 *
 * 示例 2：
 *
 *
 * 输入：nums = [1,5,233,7]
 * 输出：true
 * 解释：玩家 1 一开始选择 1 。然后玩家 2 必须从 5 和 7 中进行选择。无论玩家 2 选择了哪个，玩家 1 都可以选择 233 。
 * 最终，玩家 1（234 分）比玩家 2（12 分）获得更多的分数，所以返回 true，表示玩家 1 可以成为赢家。
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= nums.length <= 20
 * 0 <= nums[i] <= 10^7
 *
 *
 */
package lc486

// @lc code=start
type pair struct {
	fir int
	sec int
}

func PredictTheWinner(nums []int) bool {
	n := len(nums)
	dp := make([][]pair, n)
	for i := range dp {
		dp[i] = make([]pair, n)
		for j := range dp[i] {
			dp[i][j] = pair{0, 0}
		}
	}

	// base case
	for i := 0; i < n; i++ {
		dp[i][i].fir = nums[i]
		dp[i][i].sec = 0
	}

	// 根据dp状态转移方程，二维矩阵只有右上部分有值，且对角线作为base case
	for i := n - 2; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			left := nums[i] + dp[i+1][j].sec
			right := nums[j] + dp[i][j-1].sec
			// 先手肯定会选择更大的结果，后手的选择随之改变
			if left > right {
				dp[i][j].fir = left
				dp[i][j].sec = dp[i+1][j].fir
			} else {
				dp[i][j].fir = right
				dp[i][j].sec = dp[i][j-1].fir
			}
		}
	}
	return dp[0][n-1].fir >= dp[0][n-1].sec
}

// dp 定义：
// dp[i][j].fir = x 表示，对于 nums[i...j] 这部分石头堆，先手能获得的最高分数为 x。
// dp[i][j].sec = y 表示，对于 nums[i...j] 这部分石头堆，后手能获得的最高分数为 y。

// 状态转移：
// dp[i][j].fir = max(nums[i] + dp[i+1][j].sec, nums[j] + dp[i][j-1].sec)
// dp[i][j].fir = max(     选择最左边的石头堆     ,     选择最右边的石头堆      )
// 解释：我作为先手，面对 nums[i...j] 时，有两种选择：
// 要么我选择最左边的那一堆石头，然后面对 nums[i+1...j]
// 但是此时轮到对方，相当于我变成了后手；
// 要么我选择最右边的那一堆石头，然后面对 nums[i...j-1]
// 但是此时轮到对方，相当于我变成了后手。

// if 先手选择左边:
//     dp[i][j].sec = dp[i+1][j].fir
// if 先手选择右边:
//     dp[i][j].sec = dp[i][j-1].fir
// 解释：我作为后手，要等先手先选择，有两种情况：
// 如果先手选择了最左边那堆，给我剩下了 nums[i+1...j]
// 此时轮到我，我变成了先手；
// 如果先手选择了最右边那堆，给我剩下了 nums[i...j-1]
// 此时轮到我，我变成了先手。

// @lc code=end
