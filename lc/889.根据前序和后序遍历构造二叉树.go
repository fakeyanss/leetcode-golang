/*
 * @lc app=leetcode.cn id=889 lang=golang
 *
 * [889] 根据前序和后序遍历构造二叉树
 *
 * https://leetcode.cn/problems/construct-binary-tree-from-preorder-and-postorder-traversal/description/
 *
 * algorithms
 * Medium (67.83%)
 * Likes:    269
 * Dislikes: 0
 * Total Accepted:    29.5K
 * Total Submissions: 43.5K
 * Testcase Example:  '[1,2,4,5,3,6,7]\n[4,5,2,6,7,3,1]'
 *
 * 给定两个整数数组，preorder 和 postorder ，其中 preorder 是一个具有 无重复 值的二叉树的前序遍历，postorder
 * 是同一棵树的后序遍历，重构并返回二叉树。
 *
 * 如果存在多个答案，您可以返回其中 任何 一个。
 *
 *
 *
 * 示例 1：
 *
 *
 *
 *
 * 输入：preorder = [1,2,4,5,3,6,7], postorder = [4,5,2,6,7,3,1]
 * 输出：[1,2,3,4,5,6,7]
 *
 *
 * 示例 2:
 *
 *
 * 输入: preorder = [1], postorder = [1]
 * 输出: [1]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= preorder.length <= 30
 * 1 <= preorder[i] <= preorder.length
 * preorder 中所有值都 不同
 * postorder.length == preorder.length
 * 1 <= postorder[i] <= postorder.length
 * postorder 中所有值都 不同
 * 保证 preorder 和 postorder 是同一棵二叉树的前序遍历和后序遍历
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

var valToIndex889 map[int]int = make(map[int]int)

func constructFromPrePost(preorder []int, postorder []int) *TreeNode {
	for i, val := range postorder {
		valToIndex889[val] = i
	}
	return buildTree889(preorder, 0, len(preorder)-1, postorder, 0, len(postorder)-1)
}

func buildTree889(pre []int, preStart, preEnd int, post []int, postStart, postEnd int) *TreeNode {
	if preStart > preEnd {
		return nil
	}
	if preStart == preEnd {
		return &TreeNode{pre[preStart], nil, nil}
	}

	rootVal := pre[preStart]
	leftRootVal := pre[preStart+1]
	index := valToIndex889[leftRootVal]
	leftSize := index - postStart + 1

	return &TreeNode{rootVal,
		buildTree889(pre, preStart+1, preStart+leftSize, post, postStart, index),
		buildTree889(pre, preStart+leftSize+1, preEnd, post, index+1, postEnd-1)}
}

// @lc code=end
