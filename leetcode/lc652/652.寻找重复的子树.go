/*
 * @lc app=leetcode.cn id=652 lang=golang
 *
 * [652] 寻找重复的子树
 *
 * https://leetcode.cn/problems/find-duplicate-subtrees/description/
 *
 * algorithms
 * Medium (58.76%)
 * Likes:    460
 * Dislikes: 0
 * Total Accepted:    58.9K
 * Total Submissions: 100.2K
 * Testcase Example:  '[1,2,3,4,null,2,4,null,null,4]'
 *
 * 给定一棵二叉树 root，返回所有重复的子树。
 *
 * 对于同一类的重复子树，你只需要返回其中任意一棵的根结点即可。
 *
 * 如果两棵树具有相同的结构和相同的结点值，则它们是重复的。
 *
 *
 *
 * 示例 1：
 *
 *
 *
 *
 * 输入：root = [1,2,3,4,null,2,4,null,null,4]
 * 输出：[[2,4],[4]]
 *
 * 示例 2：
 *
 *
 *
 *
 * 输入：root = [2,1,1]
 * 输出：[[1]]
 *
 * 示例 3：
 *
 *
 *
 *
 * 输入：root = [2,2,2,3,null,3,null]
 * 输出：[[2,3],[3]]
 *
 *
 *
 * 提示：
 *
 *
 * 树中的结点数在[1,10^4]范围内。
 * -200 <= Node.val <= 200
 *
 *
 */
package lc652

import (
	"strconv"

	"github.com/fakeyanss/leetcode-golang/helper"
)

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
func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	memo := make(map[string]int)
	res := make([]*TreeNode, 0)

	var traverse func(node *TreeNode) string
	traverse = func(node *TreeNode) string {
		if node == nil {
			return "#"
		}

		left := traverse(node.Left)
		right := traverse(node.Right)

		// 后序遍历，将所有结果存入map
		subTree := left + "," + right + "," + strconv.Itoa(node.Val)

		cnt := memo[subTree]
		// 多次重复也只会计入结果一次
		if cnt == 1 {
			res = append(res, node)
		}
		// 给子树出现次数加1
		memo[subTree] = cnt + 1
		return subTree
	}

	traverse(root)

	return res
}

// @lc code=end
