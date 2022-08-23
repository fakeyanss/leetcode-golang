/*
 * @lc app=leetcode.cn id=106 lang=golang
 *
 * [106] 从中序与后序遍历序列构造二叉树
 *
 * https://leetcode.cn/problems/construct-binary-tree-from-inorder-and-postorder-traversal/description/
 *
 * algorithms
 * Medium (72.51%)
 * Likes:    821
 * Dislikes: 0
 * Total Accepted:    223.5K
 * Total Submissions: 308.3K
 * Testcase Example:  '[9,3,15,20,7]\n[9,15,7,20,3]'
 *
 * 给定两个整数数组 inorder 和 postorder ，其中 inorder 是二叉树的中序遍历， postorder
 * 是同一棵树的后序遍历，请你构造并返回这颗 二叉树 。
 *
 *
 *
 * 示例 1:
 *
 *
 * 输入：inorder = [9,3,15,20,7], postorder = [9,15,7,20,3]
 * 输出：[3,9,20,null,null,15,7]
 *
 *
 * 示例 2:
 *
 *
 * 输入：inorder = [-1], postorder = [-1]
 * 输出：[-1]
 *
 *
 *
 *
 * 提示:
 *
 *
 * 1 <= inorder.length <= 3000
 * postorder.length == inorder.length
 * -3000 <= inorder[i], postorder[i] <= 3000
 * inorder 和 postorder 都由 不同 的值组成
 * postorder 中每一个值都在 inorder 中
 * inorder 保证是树的中序遍历
 * postorder 保证是树的后序遍历
 *
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
var valToIndex106 map[int]int = make(map[int]int)

// func buildTree(inorder []int, postorder []int) *TreeNode {
func buildTree106_(inorder []int, postorder []int) *TreeNode {
	for i, val := range inorder {
		valToIndex106[val] = i
	}
	return buildTree106(inorder, 0, len(inorder)-1, postorder, 0, len(postorder)-1)
}

func buildTree106(inorder []int, inStart, inEnd int, postorder []int, postStart, postEnd int) *TreeNode {
	if postStart > postEnd {
		return nil
	}

	rootVal := postorder[postEnd]
	index := valToIndex106[rootVal]

	leftSize := index - inStart

	root := &TreeNode{rootVal,
		buildTree106(inorder, inStart, index-1, postorder, postStart, postStart+leftSize-1),
		buildTree106(inorder, index+1, inEnd, postorder, postStart+leftSize, postEnd-1)}
	return root
}

// @lc code=end
