/*
 * @lc app=leetcode.cn id=104 lang=golang
 *
 * [104] 二叉树的最大深度
 *
 * https://leetcode.cn/problems/maximum-depth-of-binary-tree/description/
 *
 * algorithms
 * Easy (77.11%)
 * Likes:    1337
 * Dislikes: 0
 * Total Accepted:    816.3K
 * Total Submissions: 1.1M
 * Testcase Example:  '[3,9,20,null,null,15,7]'
 *
 * 给定一个二叉树，找出其最大深度。
 *
 * 二叉树的深度为根节点到最远叶子节点的最长路径上的节点数。
 *
 * 说明: 叶子节点是指没有子节点的节点。
 *
 * 示例：
 * 给定二叉树 [3,9,20,null,null,15,7]，
 *
 * ⁠   3
 * ⁠  / \
 * ⁠ 9  20
 * ⁠   /  \
 * ⁠  15   7
 *
 * 返回它的最大深度 3 。
 *
 */
package lc104

import "github.com/fakeYanss/leetcode-golang/helper"

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
func maxDepth(root *TreeNode) int {
	// var res, depth int
	// var traverse func(node *TreeNode)
	// traverse = func(node *TreeNode) {
	// 	if node == nil {
	// 		return
	// 	}
	// 	depth++
	// 	if node.Left == nil && node.Right == nil {
	// 		if res < depth {
	// 			res = depth
	// 		}
	// 	}
	// 	traverse(node.Left)
	// 	traverse(node.Right)
	// 	depth--
	// }
	// traverse(root)
	// return res

	if root == nil {
		return 0
	}
	leftMax := maxDepth(root.Left)
	rightMax := maxDepth(root.Right)
	if leftMax > rightMax {
		return leftMax + 1
	}
	return rightMax + 1
}

// @lc code=end
