/*
 * @lc app=leetcode.cn id=669 lang=golang
 *
 * [669] 修剪二叉搜索树
 *
 * https://leetcode.cn/problems/trim-a-binary-search-tree/description/
 *
 * algorithms
 * Medium (66.79%)
 * Likes:    588
 * Dislikes: 0
 * Total Accepted:    73K
 * Total Submissions: 109.3K
 * Testcase Example:  '[1,0,2]\n1\n2'
 *
 * 给你二叉搜索树的根节点 root ，同时给定最小边界low 和最大边界 high。通过修剪二叉搜索树，使得所有节点的值在[low, high]中。修剪树
 * 不应该 改变保留在树中的元素的相对结构 (即，如果没有被移除，原有的父代子代关系都应当保留)。 可以证明，存在 唯一的答案 。
 *
 * 所以结果应当返回修剪好的二叉搜索树的新的根节点。注意，根节点可能会根据给定的边界发生改变。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：root = [1,0,2], low = 1, high = 2
 * 输出：[1,null,2]
 *
 *
 * 示例 2：
 *
 *
 * 输入：root = [3,0,4,null,2,null,null,1], low = 1, high = 3
 * 输出：[3,2,null,1]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 树中节点数在范围 [1, 10^4] 内
 * 0 <= Node.val <= 10^4
 * 树中每个节点的值都是 唯一 的
 * 题目数据保证输入是一棵有效的二叉搜索树
 * 0 <= low <= high <= 10^4
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
func trimBST(root *TreeNode, low int, high int) *TreeNode {
	if root == nil {
		return nil
	}

	// 1、root.val < lo，这种情况下 root 节点本身和 root 的左子树全都是小于 lo 的，都需要被剪掉。
	// 2、root.val > hi，这种情况下 root 节点本身和 root 的右子树全都是大于 hi 的，都需要被剪掉。
	if root.Val < low {
		// 直接返回 root.right, 等于删除 root 以及 root 的左子树
		return trimBST(root.Right, low, high)
	}

	if root.Val > high {
		// 直接返回 root.left, 等于删除 root 以及 root的右子树
		return trimBST(root.Left, low, high)
	}

	// 闭区间 [lo, hi] 内的节点什么都不做
	root.Left = trimBST(root.Left, low, high)
	root.Right = trimBST(root.Right, low, high)
	return root
}

// @lc code=end
