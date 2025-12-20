/*
 * @lc app=leetcode.cn id=230 lang=golang
 *
 * [230] 二叉搜索树中第K小的元素
 *
 * https://leetcode.cn/problems/kth-smallest-element-in-a-bst/description/
 *
 * algorithms
 * Medium (75.68%)
 * Likes:    662
 * Dislikes: 0
 * Total Accepted:    232.8K
 * Total Submissions: 307.6K
 * Testcase Example:  '[3,1,4,null,2]\n1'
 *
 * 给定一个二叉搜索树的根节点 root ，和一个整数 k ，请你设计一个算法查找其中第 k 个最小元素（从 1 开始计数）。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：root = [3,1,4,null,2], k = 1
 * 输出：1
 *
 *
 * 示例 2：
 *
 *
 * 输入：root = [5,3,6,2,4,null,null,1], k = 3
 * 输出：3
 *
 *
 *
 *
 *
 *
 * 提示：
 *
 *
 * 树中的节点数为 n 。
 * 1
 * 0
 *
 *
 *
 *
 * 进阶：如果二叉搜索树经常被修改（插入/删除操作）并且你需要频繁地查找第 k 小的值，你将如何优化算法？
 *
 */
package lc230

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
func kthSmallest(root *TreeNode, k int) int {
	// // 中序遍历，递归
	// var traverse func(*TreeNode) int
	// traverse = func(node *TreeNode) int {
	//     if node == nil {
	//         return -1
	//     }
	//     leftRes := traverse(node.Left)
	//     if leftRes != -1 {
	//         return leftRes // 拿到结果提前返回
	//     }
	//     k--
	//     if k == 0 {
	//         return node.Val
	//     }
	//     return traverse(node.Right)
	// }
	// return traverse(root)
	// 迭代也可以提前结束
	stk := []*TreeNode{}
	cur := root
	for len(stk) > 0 || cur != nil {
		for cur != nil {
			stk = append(stk, cur)
			cur = cur.Left
		}
		cur = stk[len(stk)-1]
		stk = stk[:len(stk)-1]
		k--
		if k == 0 {
			return cur.Val
		}
		cur = cur.Right
	}
	return -1
}

// @lc code=end
