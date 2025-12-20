/*
 * @lc app=leetcode.cn id=101 lang=golang
 *
 * [101] 对称二叉树
 *
 * https://leetcode.cn/problems/symmetric-tree/description/
 *
 * algorithms
 * Easy (57.90%)
 * Likes:    1985
 * Dislikes: 0
 * Total Accepted:    626.4K
 * Total Submissions: 1.1M
 * Testcase Example:  '[1,2,2,3,4,4,3]'
 *
 * 给你一个二叉树的根节点 root ， 检查它是否轴对称。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：root = [1,2,2,3,4,4,3]
 * 输出：true
 *
 *
 * 示例 2：
 *
 *
 * 输入：root = [1,2,2,null,3,null,3]
 * 输出：false
 *
 *
 *
 *
 * 提示：
 *
 *
 * 树中节点数目在范围 [1, 1000] 内
 * -100 <= Node.val <= 100
 *
 *
 *
 *
 * 进阶：你可以运用递归和迭代两种方法解决这个问题吗？
 *
 */
package lc101

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
func isSymmetric(root *TreeNode) bool {
	// // 递归
	// var isSameTree func(l, r *TreeNode) bool
	// isSameTree = func(l, r *TreeNode) bool {
	// 	if l == nil && r == nil {
	// 		return true
	// 	}
	// 	if l == nil || r == nil {
	// 		return false
	// 	}
	// 	if l.Val != r.Val {
	// 		return false
	// 	}
	// 	return isSameTree(l.Left, r.Right) && isSameTree(l.Right, r.Left)
	// }
	// return isSameTree(root.Left, root.Right)

	// 迭代
	q := []*TreeNode{root, root}
	for len(q) > 0 {
		l, r := q[0], q[1]
		q = q[2:]
		if l == nil && r == nil {
			continue
		}
		if l == nil || r == nil {
			return false
		}
		if l.Val != r.Val {
			return false
		}
		q = append(q, l.Left, r.Right, l.Right, r.Left)
	}
	return true
}

// @lc code=end
