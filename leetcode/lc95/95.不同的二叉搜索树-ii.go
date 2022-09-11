/*
 * @lc app=leetcode.cn id=95 lang=golang
 *
 * [95] 不同的二叉搜索树 II
 *
 * https://leetcode.cn/problems/unique-binary-search-trees-ii/description/
 *
 * algorithms
 * Medium (72.67%)
 * Likes:    1295
 * Dislikes: 0
 * Total Accepted:    145.6K
 * Total Submissions: 200.4K
 * Testcase Example:  '3'
 *
 * 给你一个整数 n ，请你生成并返回所有由 n 个节点组成且节点值从 1 到 n 互不相同的不同 二叉搜索树 。可以按 任意顺序 返回答案。
 *
 *
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：n = 3
 * 输出：[[1,null,2,null,3],[1,null,3,2],[2,1,3],[3,1,null,null,2],[3,2,null,1]]
 *
 *
 * 示例 2：
 *
 *
 * 输入：n = 1
 * 输出：[[1]]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1
 *
 *
 *
 *
 */
package lc95

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
func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return []*TreeNode{}
	}
	return buildBST95(1, n)
}

func buildBST95(lo, hi int) []*TreeNode {
	if lo > hi {
		return []*TreeNode{nil}
	}

	res := make([]*TreeNode, 0)
	// 穷举 root 节点的所有可能
	for i := lo; i <= hi; i++ {
		// 递归构造出左右子树的所有合法 BST
		leftTree := buildBST95(lo, i-1)
		rightTree := buildBST95(i+1, hi)
		// 给 root 节点穷举所有左右子树的组合
		for _, l := range leftTree {
			for _, r := range rightTree {
				// i 作为根节点 root 的值
				res = append(res, &TreeNode{i, l, r})
			}
		}
	}
	return res
}

// @lc code=end
