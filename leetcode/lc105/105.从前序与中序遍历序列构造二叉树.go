/*
 * @lc app=leetcode.cn id=105 lang=golang
 *
 * [105] 从前序与中序遍历序列构造二叉树
 *
 * https://leetcode.cn/problems/construct-binary-tree-from-preorder-and-inorder-traversal/description/
 *
 * algorithms
 * Medium (71.34%)
 * Likes:    1701
 * Dislikes: 0
 * Total Accepted:    407K
 * Total Submissions: 570.4K
 * Testcase Example:  '[3,9,20,15,7]\n[9,3,15,20,7]'
 *
 * 给定两个整数数组 preorder 和 inorder ，其中 preorder 是二叉树的先序遍历， inorder
 * 是同一棵树的中序遍历，请构造二叉树并返回其根节点。
 *
 *
 *
 * 示例 1:
 *
 *
 * 输入: preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]
 * 输出: [3,9,20,null,null,15,7]
 *
 *
 * 示例 2:
 *
 *
 * 输入: preorder = [-1], inorder = [-1]
 * 输出: [-1]
 *
 *
 *
 *
 * 提示:
 *
 *
 * 1 <= preorder.length <= 3000
 * inorder.length == preorder.length
 * -3000 <= preorder[i], inorder[i] <= 3000
 * preorder 和 inorder 均 无重复 元素
 * inorder 均出现在 preorder
 * preorder 保证 为二叉树的前序遍历序列
 * inorder 保证 为二叉树的中序遍历序列
 *
 *
 */
package lc105

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
var valToIndex map[int]int = make(map[int]int)

func buildTree(preorder []int, inorder []int) *TreeNode {
	for i, val := range inorder {
		valToIndex[val] = i
	}
	return buildTree105(preorder, 0, len(preorder)-1, inorder, 0, len(inorder)-1)
}

func buildTree105(preorder []int, preStart, preEnd int, inorder []int, inStart, inEnd int) *TreeNode {
	if preStart > preEnd {
		return nil
	}

	// root是preorder第一个元素
	rootVal := preorder[preStart]
	// root.val在inorder的index
	index := valToIndex[rootVal]

	// 左子树节点个数
	leftSize := index - inStart

	// prestart                                      preEnd
	//┌───────┬──────────────────────────┬─────────────────┐    preorder
	//└───────┴──────────────────────────┼─────────────────┘
	//               preStart+leftSize ◄─┘
	//
	// inStart                                        inEnd
	//┌──────────────────────────┬───────┬─────────────────┐    inorder
	//└──────────────────────────┴───────┴─────────────────┘
	//├─────►   leftSize  ◄──────┤ index

	root := &TreeNode{rootVal, nil, nil}
	root.Left = buildTree105(preorder, preStart+1, preStart+leftSize, inorder, inStart, index-1)
	root.Right = buildTree105(preorder, preStart+leftSize+1, preEnd, inorder, index+1, inEnd)
	return root
}

// @lc code=end
