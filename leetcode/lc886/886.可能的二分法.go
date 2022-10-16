/*
 * @lc app=leetcode.cn id=886 lang=golang
 *
 * [886] 可能的二分法
 *
 * https://leetcode.cn/problems/possible-bipartition/description/
 *
 * algorithms
 * Medium (49.87%)
 * Likes:    262
 * Dislikes: 0
 * Total Accepted:    28.4K
 * Total Submissions: 56.7K
 * Testcase Example:  '4\n[[1,2],[1,3],[2,4]]'
 *
 * 给定一组 n 人（编号为 1, 2, ..., n）， 我们想把每个人分进任意大小的两组。每个人都可能不喜欢其他人，那么他们不应该属于同一组。
 *
 * 给定整数 n 和数组 dislikes ，其中 dislikes[i] = [ai, bi] ，表示不允许将编号为 ai 和
 * bi的人归入同一组。当可以用这种方法将所有人分进两组时，返回 true；否则返回 false。
 *
 *
 *
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：n = 4, dislikes = [[1,2],[1,3],[2,4]]
 * 输出：true
 * 解释：group1 [1,4], group2 [2,3]
 *
 *
 * 示例 2：
 *
 *
 * 输入：n = 3, dislikes = [[1,2],[1,3],[2,3]]
 * 输出：false
 *
 *
 * 示例 3：
 *
 *
 * 输入：n = 5, dislikes = [[1,2],[2,3],[3,4],[4,5],[1,5]]
 * 输出：false
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= n <= 2000
 * 0 <= dislikes.length <= 10^4
 * dislikes[i].length == 2
 * 1 <= dislikes[i][j] <= n
 * ai < bi
 * dislikes 中每一组都 不同
 *
 *
 *
 *
 */
package lc886

// @lc code=start
func possibleBipartition(n int, dislikes [][]int) bool {
	return isBipartiteGraph(n, dislikes)
}

// 二分图解法
var (
	ok      bool
	color   []bool
	visited []bool
)

func isBipartiteGraph(n int, dislikes [][]int) bool {
	// 图编号从1开始
	ok = true
	color = make([]bool, n+1)
	visited = make([]bool, n+1)
	// 转化为邻接表
	graph := buildGraph(n, dislikes)

	for v := 1; v <= n; v++ {
		if !visited[v] {
			traverse(graph, v)
		}
	}
	return ok
}

func buildGraph(n int, dislikes [][]int) [][]int {
	// 图节点编号 1..n
	graph := make([][]int, n+1)
	for i := 1; i <= n; i++ {
		graph[i] = make([]int, 0)
	}
	for _, edge := range dislikes {
		a, b := edge[0], edge[1]
		// 无向图相当于双向图
		graph[a] = append(graph[a], b)
		graph[b] = append(graph[b], a)
	}
	return graph
}

func traverse(graph [][]int, v int) {
	if !ok {
		return
	}
	visited[v] = true
	for _, w := range graph[v] {
		if !visited[w] {
			// 相邻节点 w 没有被访问过
			// 那么应该给节点 w 涂上和节点 v 不同的颜色
			color[w] = !color[v]
			traverse(graph, w)
		} else {
			// 相邻节点 w 已经被访问过
			// 根据 v 和 w 的颜色判断是否是二分图
			if color[w] == color[v] {
				ok = false
			}
		}
	}
}

// @lc code=end
