/*
 * @lc app=leetcode.cn id=114 lang=golang
 *
 * [114] 二叉树展开为链表
 *
 * https://leetcode.cn/problems/flatten-binary-tree-to-linked-list/description/
 *
 * algorithms
 * Medium (72.95%)
 * Likes:    1278
 * Dislikes: 0
 * Total Accepted:    294.3K
 * Total Submissions: 403.4K
 * Testcase Example:  '[1,2,5,3,4,null,6]'
 *
 * 给你二叉树的根结点 root ，请你将它展开为一个单链表：
 *
 *
 * 展开后的单链表应该同样使用 TreeNode ，其中 right 子指针指向链表中下一个结点，而左子指针始终为 null 。
 * 展开后的单链表应该与二叉树 先序遍历 顺序相同。
 *
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：root = [1,2,5,3,4,null,6]
 * 输出：[1,null,2,null,3,null,4,null,5,null,6]
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
 * 输入：root = [0]
 * 输出：[0]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 树中结点数在范围 [0, 2000] 内
 * -100
 *
 *
 *
 *
 * 进阶：你可以使用原地算法（O(1) 额外空间）展开这棵树吗？
 *
 */
package lc114

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
// 定义：输入节点 root，然后 root 为根的二叉树就会被拉平为一条链表
func flatten(root *TreeNode) {
	if root == nil {
		return
	}

	// 根据函数定义，已经拉平左右子树
	flatten(root.Left)
	flatten(root.Right)

	// 后序遍历位置
	l, r := root.Left, root.Right

	// 左子树换到右子树位置
	root.Left = nil
	root.Right = l

	// 左子树（当前的右子树）末端接上右子树（之前的右子树）
	p := root
	for p.Right != nil {
		p = p.Right
	}
	p.Right = r
}

// @lc code=end
