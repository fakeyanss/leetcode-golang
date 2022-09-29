/*
 * @lc app=leetcode.cn id=337 lang=golang
 *
 * [337] 打家劫舍 III
 *
 * https://leetcode.cn/problems/house-robber-iii/description/
 *
 * algorithms
 * Medium (60.94%)
 * Likes:    1447
 * Dislikes: 0
 * Total Accepted:    226.4K
 * Total Submissions: 371.5K
 * Testcase Example:  '[3,2,3,null,3,null,1]'
 *
 * 小偷又发现了一个新的可行窃的地区。这个地区只有一个入口，我们称之为 root 。
 *
 * 除了 root 之外，每栋房子有且只有一个“父“房子与之相连。一番侦察之后，聪明的小偷意识到“这个地方的所有房屋的排列类似于一棵二叉树”。 如果
 * 两个直接相连的房子在同一天晚上被打劫 ，房屋将自动报警。
 *
 * 给定二叉树的 root 。返回 在不触动警报的情况下 ，小偷能够盗取的最高金额 。
 *
 *
 *
 * 示例 1:
 *
 *
 *
 *
 * 输入: root = [3,2,3,null,3,null,1]
 * 输出: 7
 * 解释: 小偷一晚能够盗取的最高金额 3 + 3 + 1 = 7
 *
 * 示例 2:
 *
 *
 *
 *
 * 输入: root = [3,4,5,1,3,null,1]
 * 输出: 9
 * 解释: 小偷一晚能够盗取的最高金额 4 + 5 = 9
 *
 *
 *
 *
 * 提示：
 *
 *
 *
 *
 * 树的节点数在 [1, 10^4] 范围内
 * 0 <= Node.val <= 10^4
 *
 *
 */
package lc337

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

func rob(root *TreeNode) int {

	// case 执行超时
	// memo := make(map[*TreeNode]int)

	// var dfs func(*TreeNode) int
	// dfs = func(node *TreeNode) int {
	// 	// base case
	// 	if node == nil {
	// 		return 0
	// 	}

	// 	if val, ok := memo[node]; ok {
	// 		return val
	// 	}

	// 	doRob := node.Val
	// 	if node.Left != nil {
	// 		doRob += rob(node.Left.Left) + rob(node.Left.Right)
	// 	}
	// 	if node.Right != nil {
	// 		doRob += rob(node.Right.Left) + rob(node.Right.Right)
	// 	}

	// 	notDoRob := rob(node.Left) + rob(node.Right)

	// 	res := maxInt(doRob, notDoRob)
	// 	memo[node] = res
	// 	return res
	// }

	// return dfs(root)

	var postOrder func(node *TreeNode) []int
	postOrder = func(node *TreeNode) []int {
		if node == nil {
			return []int{0, 0}
		}
		// 递归左右节点，得到左右节点偷与不偷的金钱
		left := postOrder(node.Left)
		right := postOrder(node.Right)
		// 偷当前节点与不偷左右子节点的值
		val1 := node.Val + left[1] + right[1]
		// 不偷当前节点，偷左右子节点的/不偷的大值
		val2 := maxInt(left[0], left[1]) + maxInt(right[0], right[1])
		// dp数组（dp table）以及下标的含义：下标为1记录不偷该节点所得到的的最大金钱，下标为0记录偷该节点所得到的的最大金钱
		return []int{val1, val2}
	}
	res := postOrder(root)
	return maxInt(res[0], res[1])
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
