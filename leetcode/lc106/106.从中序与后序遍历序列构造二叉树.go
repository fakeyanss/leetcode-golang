/*
 * @lc app=leetcode.cn id=106 lang=golang
 * @lcpr version=20004
 *
 * [106] 从中序与后序遍历序列构造二叉树
 *
 * https://leetcode.cn/problems/construct-binary-tree-from-inorder-and-postorder-traversal/description/
 *
 * algorithms
 * Medium (72.60%)
 * Likes:    1299
 * Dislikes: 0
 * Total Accepted:    440K
 * Total Submissions: 605.6K
 * Testcase Example:  '[9,3,15,20,7]\n[9,15,7,20,3]'
 *
 * 给定两个整数数组 inorder 和 postorder ，其中 inorder 是二叉树的中序遍历， postorder
 * 是同一棵树的后序遍历，请你构造并返回这颗 二叉树 。
 *
 *
 *
 * 示例 1:
 *
 * 输入：inorder = [9,3,15,20,7], postorder = [9,15,7,20,3]
 * 输出：[3,9,20,null,null,15,7]
 *
 *
 * 示例 2:
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
package lc106

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
func buildTree(inorder []int, postorder []int) *TreeNode {
	n := len(inorder)
	valToIndex := make(map[int]int, n)
	for i, val := range inorder {
		valToIndex[val] = i
	}
	var dfs func(postStart, postEnd, inStart, inEnd int) *TreeNode
	dfs = func(postStart, postEnd, inStart, inEnd int) *TreeNode {
		if postStart == postEnd {
			return nil
		}
		rootVal := postorder[postEnd-1]
		leftSize := valToIndex[rootVal] - inStart
		left := dfs(postStart, postStart+leftSize, inStart, inStart+leftSize)
		right := dfs(postStart+leftSize, postEnd-1, inStart+leftSize+1, inEnd)
		return &TreeNode{rootVal, left, right}
	}
	return dfs(0, n, 0, n) // 左闭右开区间
}

// @lc code=end

/*
// @lcpr case=start
// [9,3,15,20,7]\n[9,15,7,20,3]\n
// @lcpr case=end

// @lcpr case=start
// [-1]\n[-1]\n
// @lcpr case=end

*/
