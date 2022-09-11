/*
 * @lc app=leetcode.cn id=543 lang=golang
 *
 * [543] 二叉树的直径
 *
 * https://leetcode.cn/problems/diameter-of-binary-tree/description/
 *
 * algorithms
 * Easy (57.59%)
 * Likes:    1133
 * Dislikes: 0
 * Total Accepted:    253.9K
 * Total Submissions: 440.9K
 * Testcase Example:  '[1,2,3,4,5]'
 *
 * 给定一棵二叉树，你需要计算它的直径长度。一棵二叉树的直径长度是任意两个结点路径长度中的最大值。这条路径可能穿过也可能不穿过根结点。
 *
 *
 *
 * 示例 :
 * 给定二叉树
 *
 * ⁠         1
 * ⁠        / \
 * ⁠       2   3
 * ⁠      / \
 * ⁠     4   5
 *
 *
 * 返回 3, 它的长度是路径 [4,2,1,3] 或者 [5,2,1,3]。
 *
 *
 *
 * 注意：两结点之间的路径长度是以它们之间边的数目表示。
 *
 */
package lc543

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
func diameterOfBinaryTree(root *TreeNode) int {
	res := 0
	var postOrder func(root *TreeNode) int
	postOrder = func(root *TreeNode) int {
		if root == nil {
			return 0
		}

		left := postOrder(root.Left)
		right := postOrder(root.Right)

		if res < left+right {
			res = left + right
		}

		if left > right {
			return left + 1
		} else {
			return right + 1
		}
	}

	postOrder(root)

	return res
}

// @lc code=end
