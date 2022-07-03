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
func inorderTraversal(root *TreeNode) []int {
	// 递归
	// var inorder func(cur *TreeNode, arr *[]int)
	// inorder = func(cur *TreeNode, arr *[]int) {
	// 	if cur == nil {
	// 		return
	// 	}
	// 	inorder(cur.Left, arr)
	// 	*arr = append(*arr, cur.Val)
	// 	inorder(cur.Right, arr)
	// }

	// res := []int{}
	// inorder(root, &res)
	// return res

	// 迭代
	// stack := []*TreeNode{}
	// res := []int{}
	// for root != nil || len(stack) > 0 {
	// 	for root != nil {
	// 		stack = append(stack, root)
	// 		root = root.Left
	// 	}
	// 	root = stack[len(stack)-1]
	// 	stack = stack[:len(stack)-1]
	// 	res = append(res, root.Val)
	// 	root = root.Right
	// }
	// return res

	// 迭代&颜色标记

	WHITE, GRAY := 0, 1
	res := []int{}
	type Ele struct {
		color int
		node  *TreeNode
	}
	stack := []Ele{
		{WHITE, root},
	}
	for len(stack) > 0 {
		ele := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if ele.node == nil {
			continue
		}
		if ele.color == WHITE {
			stack = append(stack, Ele{WHITE, ele.node.Right})
			stack = append(stack, Ele{GRAY, ele.node})
			stack = append(stack, Ele{WHITE, ele.node.Left})
		} else {
			res = append(res, ele.node.Val)
		}
	}
	return res
}

// @lc code=end
