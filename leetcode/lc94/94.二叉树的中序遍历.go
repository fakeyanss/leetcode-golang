/*
 * @lc app=leetcode.cn id=94 lang=golang
 *
 * [94] 二叉树的中序遍历
 *
 * https://leetcode.cn/problems/binary-tree-inorder-traversal/description/
 *
 * algorithms
 * Easy (75.97%)
 * Likes:    1476
 * Dislikes: 0
 * Total Accepted:    868.7K
 * Total Submissions: 1.1M
 * Testcase Example:  '[1,null,2,3]'
 *
 * 给定一个二叉树的根节点 root ，返回 它的 中序 遍历 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：root = [1,null,2,3]
 * 输出：[1,3,2]
 *
 *
 * 示例 2：
 *
 *
 * 输入：root = []
 * 输出：[]
 *
 *
 * 示例 3：
 *
 *
 * 输入：root = [1]
 * 输出：[1]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 树中节点数目在范围 [0, 100] 内
 * -100 <= Node.val <= 100
 *
 *
 *
 *
 * 进阶: 递归算法很简单，你可以通过迭代算法完成吗？
 *
 */
package lc94

import "github.com/fakeyanss/leetcode-golang/helper"

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
func inorderTraversal(root *TreeNode) []int {
	// 递归
	// var res []int
	// var traverse func(*TreeNode)
	// traverse = func(node *TreeNode) {
	// 	if node == nil {
	// 		return
	// 	}
	// 	traverse(node.Left)
	// 	res = append(res, node.Val)
	// 	traverse(node.Right)
	// }
	// traverse(root)
	// return res

	// 迭代
	var res []int
	cur := root
	stk := []*TreeNode{}
	for cur != nil || len(stk) > 0 {
		for cur != nil {
			stk = append(stk, cur)
			cur = cur.Left
		}
		cur = stk[len(stk)-1]
		stk = stk[:len(stk)-1]
		res = append(res, cur.Val)
		cur = cur.Right
	}
	return res
}

// @lc code=end
