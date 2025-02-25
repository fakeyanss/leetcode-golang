/*
 * @lc app=leetcode.cn id=133 lang=golang
 * @lcpr version=20004
 *
 * [133] 克隆图
 *
 * https://leetcode.cn/problems/clone-graph/description/
 *
 * algorithms
 * Medium (72.16%)
 * Likes:    759
 * Dislikes: 0
 * Total Accepted:    169.9K
 * Total Submissions: 234.9K
 * Testcase Example:  '[[2,4],[1,3],[2,4],[1,3]]'
 *
 * 给你无向 连通 图中一个节点的引用，请你返回该图的 深拷贝（克隆）。
 *
 * 图中的每个节点都包含它的值 val（int） 和其邻居的列表（list[Node]）。
 *
 * class Node {
 * ⁠   public int val;
 * ⁠   public List<Node> neighbors;
 * }
 *
 *
 *
 * 测试用例格式：
 *
 * 简单起见，每个节点的值都和它的索引相同。例如，第一个节点值为 1（val = 1），第二个节点值为 2（val =
 * 2），以此类推。该图在测试用例中使用邻接列表表示。
 *
 * 邻接列表 是用于表示有限图的无序列表的集合。每个列表都描述了图中节点的邻居集。
 *
 * 给定节点将始终是图中的第一个节点（值为 1）。你必须将 给定节点的拷贝 作为对克隆图的引用返回。
 *
 *
 *
 * 示例 1：
 *
 *
 *
 * 输入：adjList = [[2,4],[1,3],[2,4],[1,3]]
 * 输出：[[2,4],[1,3],[2,4],[1,3]]
 * 解释：
 * 图中有 4 个节点。
 * 节点 1 的值是 1，它有两个邻居：节点 2 和 4 。
 * 节点 2 的值是 2，它有两个邻居：节点 1 和 3 。
 * 节点 3 的值是 3，它有两个邻居：节点 2 和 4 。
 * 节点 4 的值是 4，它有两个邻居：节点 1 和 3 。
 *
 *
 * 示例 2：
 *
 *
 *
 * 输入：adjList = [[]]
 * 输出：[[]]
 * 解释：输入包含一个空列表。该图仅仅只有一个值为 1 的节点，它没有任何邻居。
 *
 *
 * 示例 3：
 *
 * 输入：adjList = []
 * 输出：[]
 * 解释：这个图是空的，它不含任何节点。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 这张图中的节点数在 [0, 100] 之间。
 * 1 <= Node.val <= 100
 * 每个节点值 Node.val 都是唯一的，
 * 图中没有重复的边，也没有自环。
 * 图是连通图，你可以从给定节点访问到所有节点。
 *
 *
 */
package lc133

// @lcpr-template-start

type Node struct {
	Val       int
	Neighbors []*Node
}

// @lcpr-template-end
// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

var visited = map[*Node]*Node{}

func cloneGraph(node *Node) *Node {
	// return dfs(node)
	return bfs(node)
}

func dfs(node *Node) *Node {
	if node == nil {
		return nil
	}
	if val, ok := visited[node]; ok { // 如果该节点已经被访问过了，则直接从哈希表中取出对应的克隆节点返回
		return val
	}

	cloneNode := &Node{node.Val, []*Node{}} // 克隆节点，注意到为了深拷贝我们不会克隆它的邻居的列表
	visited[node] = cloneNode

	// 遍历该节点的邻居并更新克隆节点的邻居列表
	for _, n := range node.Neighbors {
		cloneNode.Neighbors = append(cloneNode.Neighbors, dfs(n))
	}
	return cloneNode
}

func bfs(node *Node) *Node {
	if node == nil {
		return nil
	}
	q := []*Node{node}
	visited[node] = &Node{node.Val, []*Node{}}
	for len(q) > 0 {
		cur := q[0]
		q = q[1:]
		// 遍历队头节点的邻居
		for _, n := range cur.Neighbors {
			if _, ok := visited[n]; !ok {
				// 如果没有被访问过，就克隆并存储在哈希表中，并加入队列
				visited[n] = &Node{n.Val, []*Node{}}
				q = append(q, n)
			}
			// 更新当前节点的邻居列表
			visited[cur].Neighbors = append(visited[cur].Neighbors, visited[n])
		}
	}
	return visited[node]
}

// @lc code=end

/*
// @lcpr case=start
// [[2,4],[1,3],[2,4],[1,3]]\n
// @lcpr case=end

// @lcpr case=start
// [[]]\n
// @lcpr case=end

// @lcpr case=start
// []\n
// @lcpr case=end

*/
