/*
 * @lc app=leetcode.cn id=787 lang=golang
 *
 * [787] K 站中转内最便宜的航班
 *
 * https://leetcode.cn/problems/cheapest-flights-within-k-stops/description/
 *
 * algorithms
 * Medium (39.26%)
 * Likes:    527
 * Dislikes: 0
 * Total Accepted:    58.7K
 * Total Submissions: 149.5K
 * Testcase Example:  '4\n[[0,1,100],[1,2,100],[2,0,100],[1,3,600],[2,3,200]]\n0\n3\n1'
 *
 * 有 n 个城市通过一些航班连接。给你一个数组 flights ，其中 flights[i] = [fromi, toi, pricei]
 * ，表示该航班都从城市 fromi 开始，以价格 pricei 抵达 toi。
 *
 * 现在给定所有的城市和航班，以及出发城市 src 和目的地 dst，你的任务是找到出一条最多经过 k 站中转的路线，使得从 src 到 dst 的
 * 价格最便宜 ，并返回该价格。 如果不存在这样的路线，则输出 -1。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入:
 * n = 3, edges = [[0,1,100],[1,2,100],[0,2,500]]
 * src = 0, dst = 2, k = 1
 * 输出: 200
 * 解释:
 * 城市航班图如下
 *
 *
 * 从城市 0 到城市 2 在 1 站中转以内的最便宜价格是 200，如图中红色所示。
 *
 * 示例 2：
 *
 *
 * 输入:
 * n = 3, edges = [[0,1,100],[1,2,100],[0,2,500]]
 * src = 0, dst = 2, k = 0
 * 输出: 500
 * 解释:
 * 城市航班图如下
 *
 *
 * 从城市 0 到城市 2 在 0 站中转以内的最便宜价格是 500，如图中蓝色所示。
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= n <= 100
 * 0 <= flights.length <= (n * (n - 1) / 2)
 * flights[i].length == 3
 * 0 <= fromi, toi < n
 * fromi != toi
 * 1 <= pricei <= 10^4
 * 航班没有重复，且不存在自环
 * 0 <= src, dst, k < n
 * src != dst
 *
 *
 */
package lc787

import "math"

// @lc code=start
var (
	indegree map[int][]node
	memo     [][]int
	source   int
)

type node struct {
	from  int
	price int
}

func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	// k个中转站，即k+1条边
	k++
	source = src
	memo = make([][]int, n)
	for i := range memo {
		memo[i] = make([]int, k+1)
	}

	indegree = make(map[int][]node)
	for _, f := range flights {
		from, to, price := f[0], f[1], f[2]
		if _, ok := indegree[to]; !ok {
			indegree[to] = make([]node, 0)
		}
		indegree[to] = append(indegree[to], node{from, price})
	}

	return dp(dst, k)
}

// dp(dst, k) = min(dp(s1, k-1) + w1, dp(s2, k-1) + w2, ...), w1,w2为指向dst的节点
func dp(dst, k int) int {
	// base case
	if source == dst {
		return 0
	}
	if k == 0 {
		return -1
	}

	if memo[dst][k] != 0 {
		return memo[dst][k]
	}

	res := math.MaxInt
	if nodes, ok := indegree[dst]; ok {
		// 当dst有多个入度节点时，分解为子问题
		for _, node := range nodes {
			// 从src到相邻的入度节点需要的最短路径权重
			subProblem := dp(node.from, k-1)
			// 跳过无解的结果
			if subProblem != -1 {
				res = minInt(res, subProblem+node.price)
			}
		}
	}
	if res == math.MaxInt {
		memo[dst][k] = -1
	} else {
		memo[dst][k] = res
	}
	return memo[dst][k]
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end
