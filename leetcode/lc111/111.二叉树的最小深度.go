/*
 * @lc app=leetcode.cn id=111 lang=golang
 * @lcpr version=20004
 *
 * [111] 二叉树的最小深度
 *
 * https://leetcode.cn/problems/minimum-depth-of-binary-tree/description/
 *
 * algorithms
 * Easy (55.22%)
 * Likes:    1251
 * Dislikes: 0
 * Total Accepted:    767K
 * Total Submissions: 1.4M
 * Testcase Example:  '[3,9,20,null,null,15,7]'
 *
 * 给定一个二叉树，找出其最小深度。
 *
 * 最小深度是从根节点到最近叶子节点的最短路径上的节点数量。
 *
 * 说明：叶子节点是指没有子节点的节点。
 *
 *
 *
 * 示例 1：
 *
 * 输入：root = [3,9,20,null,null,15,7]
 * 输出：2
 *
 *
 * 示例 2：
 *
 * 输入：root = [2,null,3,null,4,null,5,null,6]
 * 输出：5
 *
 *
 *
 *
 * 提示：
 *
 *
 * 树中节点数的范围在 [0, 10^5] 内
 * -1000 <= Node.val <= 1000
 *
 *
 */
package lc111

// @lcpr-template-start

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// @lcpr-template-end
// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minDepth(root *TreeNode) int {
	// if root == nil {
	// 	return 0
	// }
	// leftDepth := minDepth(root.Left)
	// rightDepth := minDepth(root.Right)
	// if leftDepth == 0 {
	// 	return rightDepth + 1
	// }
	// if rightDepth == 0 {
	// 	return leftDepth + 1
	// }
	// return min(leftDepth, rightDepth) + 1

	if root == nil {
		return 0
	}
	queue := []*TreeNode{root}
	depth := 1

	// 这一层向下遍历
	for len(queue) > 0 {
		n := len(queue)
		// 这一层从左向右遍历
		for i := 0; i < n; i++ {
			cur := queue[0]
			queue = queue[1:]

			if cur.Left == nil && cur.Right == nil {
				return depth
			}

			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}
			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}
		}

		depth++
	}

	return depth
}

// @lc code=end

/*
// @lcpr case=start
// [3,9,20,null,null,15,7]\n
// @lcpr case=end

// @lcpr case=start
// [2,null,3,null,4,null,5,null,6]\n
// @lcpr case=end

*/
