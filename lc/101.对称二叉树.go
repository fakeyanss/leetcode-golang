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
package lc

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
	// 递归
	// if root == nil {
	// 	return true
	// }
	// return checkRecursive(root.Left, root.Right)

	// 迭代
	q := []*TreeNode{}
	q = append(q, root)
	q = append(q, root)
	for len(q) > 0 {
		a, b := q[0], q[1]
		q = q[2:]
		if a == nil && b == nil {
			continue
		}
		if a == nil || b == nil {
			return false
		}
		if a.Val != b.Val {
			return false
		}
		q = append(q, a.Left)
		q = append(q, b.Right)
		q = append(q, a.Right)
		q = append(q, b.Left)
	}
	return true
}

func checkRecursive(left *TreeNode, right *TreeNode) bool {
	if left == nil || right == nil {
		return left == right
	}
	if left.Val != right.Val {
		return false
	}
	return checkRecursive(left.Right, right.Left) && checkRecursive(left.Left, right.Right)
}

// @lc code=end
