/*
 * @lc app=leetcode.cn id=105 lang=golang
 * @lcpr version=20004
 *
 * [105] 从前序与中序遍历序列构造二叉树
 *
 * https://leetcode.cn/problems/construct-binary-tree-from-preorder-and-inorder-traversal/description/
 *
 * algorithms
 * Medium (72.26%)
 * Likes:    2447
 * Dislikes: 0
 * Total Accepted:    758.7K
 * Total Submissions: 1M
 * Testcase Example:  '[3,9,20,15,7]\n[9,3,15,20,7]'
 *
 * 给定两个整数数组 preorder 和 inorder ，其中 preorder 是二叉树的先序遍历， inorder
 * 是同一棵树的中序遍历，请构造二叉树并返回其根节点。
 *
 *
 *
 * 示例 1:
 *
 * 输入: preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]
 * 输出: [3,9,20,null,null,15,7]
 *
 *
 * 示例 2:
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
func buildTree(preorder []int, inorder []int) *TreeNode {
	// prestart                                      preEnd
	//┌───────┬──────────────────────────┬─────────────────┐    preorder
	//└───────┴──────────────────────────┼─────────────────┘
	// index       preStart+1+leftSize ◄─┘
	//
	// inStart                                        inEnd
	//┌──────────────────────────┬───────┬─────────────────┐    inorder
	//└──────────────────────────┴───────┴─────────────────┘
	//├─────►   leftSize  ◄──────┤ index

	// // 递归解法 O(n^2)，每次需要找preorder[0]在inorder中的位置
	// n := len(preorder)
	// if n == 0 { // 空节点
	// 	return nil
	// }
	// leftSize := slices.Index(inorder, preorder[0]) // 左子树的大小
	// left := buildTree(preorder[1:1+leftSize], inorder[:leftSize])
	// right := buildTree(preorder[1+leftSize:], inorder[1+leftSize:])
	// return &TreeNode{preorder[0], left, right}

	// 优化 O(n)，先预存inorder中每个值的下标位置
	n := len(preorder)
	valToIndex := make(map[int]int, n)
	for i, val := range inorder {
		valToIndex[val] = i
	}

	var dfs func(preStart, preEnd, inStart, inEnd int) *TreeNode
	dfs = func(preStart, preEnd, inStart, inEnd int) *TreeNode {
		if preStart == preEnd {
			return nil
		}
		rootVal := preorder[preStart]
		leftSize := valToIndex[rootVal] - inStart
		left := dfs(preStart+1, preStart+1+leftSize, inStart, inStart+leftSize)
		right := dfs(preStart+1+leftSize, preEnd, inStart+leftSize+1, inEnd)
		return &TreeNode{rootVal, left, right}
	}
	return dfs(0, n, 0, n) // 左闭右开
}

// @lc code=end

/*
// @lcpr case=start
// [3,9,20,15,7]\n[9,3,15,20,7]\n
// @lcpr case=end

// @lcpr case=start
// [-1]\n[-1]\n
// @lcpr case=end

*/
