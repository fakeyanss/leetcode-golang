/*
 * @lc app=leetcode.cn id=124 lang=golang
 *
 * [124] 二叉树中的最大路径和
 *
 * https://leetcode.cn/problems/binary-tree-maximum-path-sum/description/
 *
 * algorithms
 * Hard (45.13%)
 * Likes:    1653
 * Dislikes: 0
 * Total Accepted:    250K
 * Total Submissions: 553.9K
 * Testcase Example:  '[1,2,3]'
 *
 * 路径 被定义为一条从树中任意节点出发，沿父节点-子节点连接，达到任意节点的序列。同一个节点在一条路径序列中 至多出现一次 。该路径 至少包含一个
 * 节点，且不一定经过根节点。
 *
 * 路径和 是路径中各节点值的总和。
 *
 * 给你一个二叉树的根节点 root ，返回其 最大路径和 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：root = [1,2,3]
 * 输出：6
 * 解释：最优路径是 2 -> 1 -> 3 ，路径和为 2 + 1 + 3 = 6
 *
 * 示例 2：
 *
 *
 * 输入：root = [-10,9,20,null,null,15,7]
 * 输出：42
 * 解释：最优路径是 15 -> 20 -> 7 ，路径和为 15 + 20 + 7 = 42
 *
 *
 *
 *
 * 提示：
 *
 *
 * 树中节点数目范围是 [1, 3 * 10^4]
 * -1000
 *
 *
 */
package lc124

import (
	"math"

	"github.com/fakeYanss/leetcode-golang/helper"
)

type TreeNode = helper.TreeNode

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxPathSum(root *TreeNode) int {
	// dp[root] = max(max(dp[root.left], 0), max(dp[root.right], 0)) + root.val
	// 后序遍历
	maxSum := math.MinInt
	var oneSideMax func(node *TreeNode) int
	oneSideMax = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		leftMax := MaxInt(oneSideMax(node.Left), 0)
		rightMax := MaxInt(oneSideMax(node.Right), 0)
		maxSum = MaxInt(maxSum, node.Val+leftMax+rightMax)

		// 这里的重点是，在遍历每一个顶点的过程中，对比对到其中的路径最大值
		return node.Val + MaxInt(leftMax, rightMax)
	}

	oneSideMax(root)
	return maxSum
}

func MaxInt(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// @lc code=end
