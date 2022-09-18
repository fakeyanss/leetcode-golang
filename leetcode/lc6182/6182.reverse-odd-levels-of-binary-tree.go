// https://leetcode.cn/problems/reverse-odd-levels-of-binary-tree/
package lc6182

import (
	"github.com/fakeYanss/leetcode-golang/helper"
)

type TreeNode = helper.TreeNode

func reverseOddLevels(root *TreeNode) *TreeNode {
	// dfs(root.Left, root.Right, true)
	bfs(root)
	return root
}

func dfs(node1, node2 *TreeNode, isOddLevel bool) {
	if node1 == nil { // 完美二叉树，node1为空时node2也为空
		return
	}
	if isOddLevel {
		node1.Val, node2.Val = node2.Val, node1.Val
	}
	dfs(node1.Left, node2.Right, !isOddLevel)
	dfs(node1.Right, node2.Left, !isOddLevel)
}

func bfs(root *TreeNode) {
	q := []*TreeNode{root}
	for level := 0; q[0].Left != nil; level ^= 1 {
		next := make([]*TreeNode, 0, len(q)*2)
		for _, node := range q {
			next = append(next, node.Left, node.Right)
		}
		q = next
		if level == 0 {
			for i, n := 0, len(q); i < n/2; i++ {
				x, y := q[i], q[n-1-i]
				x.Val, y.Val = y.Val, x.Val
			}
		}
	}
}
