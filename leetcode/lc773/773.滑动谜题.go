/*
 * @lc app=leetcode.cn id=773 lang=golang
 *
 * [773] 滑动谜题
 *
 * https://leetcode.cn/problems/sliding-puzzle/description/
 *
 * algorithms
 * Hard (70.28%)
 * Likes:    278
 * Dislikes: 0
 * Total Accepted:    31K
 * Total Submissions: 44.1K
 * Testcase Example:  '[[1,2,3],[4,0,5]]'
 *
 * 在一个 2 x 3 的板上（board）有 5 块砖瓦，用数字 1~5 来表示, 以及一块空缺用 0 来表示。一次 移动 定义为选择 0
 * 与一个相邻的数字（上下左右）进行交换.
 *
 * 最终当板 board 的结果是 [[1,2,3],[4,5,0]] 谜板被解开。
 *
 * 给出一个谜板的初始状态 board ，返回最少可以通过多少次移动解开谜板，如果不能解开谜板，则返回 -1 。
 *
 *
 *
 * 示例 1：
 *
 *
 *
 *
 * 输入：board = [[1,2,3],[4,0,5]]
 * 输出：1
 * 解释：交换 0 和 5 ，1 步完成
 *
 *
 * 示例 2:
 *
 *
 *
 *
 * 输入：board = [[1,2,3],[5,4,0]]
 * 输出：-1
 * 解释：没有办法完成谜板
 *
 *
 * 示例 3:
 *
 *
 *
 *
 * 输入：board = [[4,1,2],[5,0,3]]
 * 输出：5
 * 解释：
 * 最少完成谜板的最少移动次数是 5 ，
 * 一种移动路径:
 * 尚未移动: [[4,1,2],[5,0,3]]
 * 移动 1 次: [[4,1,2],[0,5,3]]
 * 移动 2 次: [[0,1,2],[4,5,3]]
 * 移动 3 次: [[1,0,2],[4,5,3]]
 * 移动 4 次: [[1,2,0],[4,5,3]]
 * 移动 5 次: [[1,2,3],[4,5,0]]
 *
 *
 *
 *
 * 提示：
 *
 *
 * board.length == 2
 * board[i].length == 3
 * 0 <= board[i][j] <= 5
 * board[i][j] 中每个值都 不同
 *
 *
 */
package lc773

import "strings"

// @lc code=start

// board转为一维字符串start后，start[i]的相邻位置为neighbors[i]
var neighbors = [6][]int{{1, 3}, {0, 2, 4}, {1, 5}, {0, 4}, {1, 3, 5}, {2, 4}}

const target = "123450"

func slidingPuzzle(board [][]int) int {
	s := make([]byte, 0, 6)
	for _, i := range board {
		for _, j := range i {
			s = append(s, '0'+byte(j))
		}
	}
	start := string(s)

	// 开始bfs遍历
	q := make([]string, 0)
	visited := make(map[string]bool)

	q = append(q, start)
	visited[start] = true
	step := 0

	for len(q) > 0 {
		n := len(q)
		for i := 0; i < n; i++ {
			cur := q[0]
			q = q[1:]

			if target == cur {
				return step
			}

			// 找到数字0的位置
			idx := strings.Index(cur, "0")
			// idx := 0
			// for ; cur[idx] != '0'; idx++ {
			// }
			s := []byte(cur)
			for _, nei := range neighbors[idx] {
				// 遍历0的相邻节点，交换位置添加到队列
				s[idx], s[nei] = s[nei], s[idx]
				next := string(s)
				if !visited[next] {
					q = append(q, next)
					visited[next] = true
				}
				s[idx], s[nei] = s[nei], s[idx]
			}
		}

		step++
	}

	return -1
}

// @lc code=end
