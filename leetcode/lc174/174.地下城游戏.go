/*
 * @lc app=leetcode.cn id=174 lang=golang
 *
 * [174] 地下城游戏
 *
 * https://leetcode.cn/problems/dungeon-game/description/
 *
 * algorithms
 * Hard (48.73%)
 * Likes:    669
 * Dislikes: 0
 * Total Accepted:    55K
 * Total Submissions: 112.9K
 * Testcase Example:  '[[-2,-3,3],[-5,-10,1],[10,30,-5]]'
 *
 *
 * table.dungeon, .dungeon th, .dungeon td {
 * ⁠ border:3px solid black;
 * }
 *
 * ⁠.dungeon th, .dungeon td {
 * ⁠   text-align: center;
 * ⁠   height: 70px;
 * ⁠   width: 70px;
 * }
 *
 *
 * 一些恶魔抓住了公主（P）并将她关在了地下城的右下角。地下城是由 M x N
 * 个房间组成的二维网格。我们英勇的骑士（K）最初被安置在左上角的房间里，他必须穿过地下城并通过对抗恶魔来拯救公主。
 *
 * 骑士的初始健康点数为一个正整数。如果他的健康点数在某一时刻降至 0 或以下，他会立即死亡。
 *
 * 有些房间由恶魔守卫，因此骑士在进入这些房间时会失去健康点数（若房间里的值为负整数，则表示骑士将损失健康点数）；其他房间要么是空的（房间里的值为
 * 0），要么包含增加骑士健康点数的魔法球（若房间里的值为正整数，则表示骑士将增加健康点数）。
 *
 * 为了尽快到达公主，骑士决定每次只向右或向下移动一步。
 *
 *
 *
 * 编写一个函数来计算确保骑士能够拯救到公主所需的最低初始健康点数。
 *
 * 例如，考虑到如下布局的地下城，如果骑士遵循最佳路径 右 -> 右 -> 下 -> 下，则骑士的初始健康点数至少为 7。
 *
 *
 * ⁠
 * -2 (K)
 * -3
 * 3
 * ⁠
 * ⁠
 * -5
 * -10
 * 1
 * ⁠
 * ⁠
 * 10
 * 30
 * -5 (P)
 * ⁠
 *
 *
 *
 *
 *
 * 说明:
 *
 *
 *
 * 骑士的健康点数没有上限。
 *
 * 任何房间都可能对骑士的健康点数造成威胁，也可能增加骑士的健康点数，包括骑士进入的左上角房间以及公主被监禁的右下角房间。
 *
 */
package lc174

import "math"

// @lc code=start
var memo [][]int

func calculateMinimumHP(dungeon [][]int) int {
	m, n := len(dungeon), len(dungeon[0])
	memo = make([][]int, m)
	for i := range memo {
		memo[i] = make([]int, n)
	}
	return dp(dungeon, 0, 0)
}

// dp(i, j) = x, 表示从 grid[i][j] 到达终点（右下角）所需的最少生命值是 x
// res = min(dp(i + 1, j),dp(i, j + 1)) - grid[i][j];
// dp(i, j) = res <= 0 ? 1 : res;
func dp(grid [][]int, i, j int) int {
	m, n := len(grid), len(grid[0])
	if i == m-1 && j == n-1 {
		if grid[i][j] >= 0 {
			return 1
		} else {
			return -grid[i][j] + 1
		}
	}

	if i == m || j == n {
		return math.MaxInt
	}

	if memo[i][j] != 0 {
		return memo[i][j]
	}

	res := minInt(dp(grid, i+1, j), dp(grid, i, j+1)) - grid[i][j]
	if res <= 0 {
		res = 1
	}

	memo[i][j] = res
	return memo[i][j]
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end
