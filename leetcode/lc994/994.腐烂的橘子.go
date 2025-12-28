/*
 * @lc app=leetcode.cn id=994 lang=golang
 * @lcpr version=30305
 *
 * [994] 腐烂的橘子
 *
 * https://leetcode.cn/problems/rotting-oranges/description/
 *
 * algorithms
 * Medium (55.14%)
 * Likes:    1104
 * Dislikes: 0
 * Total Accepted:    383.9K
 * Total Submissions: 695.7K
 * Testcase Example:  '[[2,1,1],[1,1,0],[0,1,1]]\n[[2,1,1],[0,1,1],[1,0,1]]\n[[0,2]]'
 *
 * 在给定的 m x n 网格 grid 中，每个单元格可以有以下三个值之一：
 *
 *
 * 值 0 代表空单元格；
 * 值 1 代表新鲜橘子；
 * 值 2 代表腐烂的橘子。
 *
 *
 * 每分钟，腐烂的橘子 周围 4 个方向上相邻 的新鲜橘子都会腐烂。
 *
 * 返回 直到单元格中没有新鲜橘子为止所必须经过的最小分钟数。如果不可能，返回 -1 。
 *
 *
 *
 * 示例 1：
 *
 *
 *
 * 输入：grid = [[2,1,1],[1,1,0],[0,1,1]]
 * 输出：4
 *
 *
 * 示例 2：
 *
 * 输入：grid = [[2,1,1],[0,1,1],[1,0,1]]
 * 输出：-1
 * 解释：左下角的橘子（第 2 行， 第 0 列）永远不会腐烂，因为腐烂只会发生在 4 个方向上。
 *
 *
 * 示例 3：
 *
 * 输入：grid = [[0,2]]
 * 输出：0
 * 解释：因为 0 分钟时已经没有新鲜橘子了，所以答案就是 0 。
 *
 *
 *
 *
 * 提示：
 *
 *
 * m == grid.length
 * n == grid[i].length
 * 1 <= m, n <= 10
 * grid[i][j] 仅为 0、1 或 2
 *
 *
 */
package lc994

// @lc code=start
type pair struct {
	x int
	y int
}

var directions = []pair{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

func orangesRotting(grid [][]int) int {
	var q []pair
	m, n := len(grid), len(grid[0])
	fresh := 0
	for i := range m {
		for j := range n {
			if grid[i][j] == 2 {
				q = append(q, pair{x: i, y: j})
			}
			if grid[i][j] == 1 {
				fresh++
			}
		}
	}

	res := 0
	for len(q) > 0 && fresh > 0 {
		res++
		sz := len(q)
		for range sz {
			rot := q[0]
			q = q[1:]
			for _, dir := range directions {
				x, y := rot.x+dir.x, rot.y+dir.y
				if x >= 0 && x < m && y >= 0 && y < n && grid[x][y] == 1 {
					grid[x][y] = 2
					fresh--
					q = append(q, pair{x, y})
				}
			}
		}

	}

	if fresh > 0 {
		return -1
	}
	return res
}

// @lc code=end

/*
// @lcpr case=start
// [[2,1,1],[1,1,0],[0,1,1]]\n
// @lcpr case=end

// @lcpr case=start
// [[2,1,1],[0,1,1],[1,0,1]]\n
// @lcpr case=end

// @lcpr case=start
// [[0,2]]\n
// @lcpr case=end

*/
