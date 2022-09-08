/*
 * @lc app=leetcode.cn id=797 lang=golang
 *
 * [797] 所有可能的路径
 *
 * https://leetcode.cn/problems/all-paths-from-source-to-target/description/
 *
 * algorithms
 * Medium (79.06%)
 * Likes:    331
 * Dislikes: 0
 * Total Accepted:    81.7K
 * Total Submissions: 103.4K
 * Testcase Example:  '[[1,2],[3],[3],[]]'
 *
 * 给你一个有 n 个节点的 有向无环图（DAG），请你找出所有从节点 0 到节点 n-1 的路径并输出（不要求按特定顺序）
 *
 * graph[i] 是一个从节点 i 可以访问的所有节点的列表（即从节点 i 到节点 graph[i][j]存在一条有向边）。
 *
 *
 *
 * 示例 1：
 *
 *
 *
 *
 * 输入：graph = [[1,2],[3],[3],[]]
 * 输出：[[0,1,3],[0,2,3]]
 * 解释：有两条路径 0 -> 1 -> 3 和 0 -> 2 -> 3
 *
 *
 * 示例 2：
 *
 *
 *
 *
 * 输入：graph = [[4,3,1],[3,2,4],[3],[4],[]]
 * 输出：[[0,4],[0,3,4],[0,1,3,4],[0,1,2,3,4],[0,1,4]]
 *
 *
 *
 *
 * 提示：
 *
 *
 * n == graph.length
 * 2 <= n <= 15
 * 0 <= graph[i][j] < n
 * graph[i][j] != i（即不存在自环）
 * graph[i] 中的所有元素 互不相同
 * 保证输入为 有向无环图（DAG）
 *
 *
 *
 *
 */
package lc

// @lc code=start
func allPathsSourceTarget(graph [][]int) [][]int {
	res := make([][]int, 0)
	path := make([]int, 0)

	var traverse func(graph [][]int, s int, path []int)
	traverse = func(graph [][]int, s int, path []int) {
		// 添加当前节点到路径
		path = append(path, s)

		// 到达终点
		if s == len(graph)-1 {
			p := make([]int, len(path))
			copy(p, path)
			res = append(res, p)
			path = path[:len(path)-1]
			return
		}

		// 递归每个相邻节点
		for _, val := range graph[s] {
			traverse(graph, val, path)
		}

		// 从路径移除当前节点
		path = path[:len(path)-1]
	}
	traverse(graph, 0, path)
	return res
}

// @lc code=end
