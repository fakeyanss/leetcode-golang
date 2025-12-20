/*
 * @lc app=leetcode.cn id=124 lang=golang
 * @lcpr version=20004
 *
 * [124] 二叉树中的最大路径和
 *
 * https://leetcode.cn/problems/binary-tree-maximum-path-sum/description/
 *
 * algorithms
 * Hard (46.16%)
 * Likes:    2337
 * Dislikes: 0
 * Total Accepted:    489.6K
 * Total Submissions: 1.1M
 * Testcase Example:  '[1,2,3]'
 *
 * 二叉树中的 路径 被定义为一条节点序列，序列中每对相邻节点之间都存在一条边。同一个节点在一条路径序列中 至多出现一次 。该路径 至少包含一个
 * 节点，且不一定经过根节点。
 *
 * 路径和 是路径中各节点值的总和。
 *
 * 给你一个二叉树的根节点 root ，返回其 最大路径和 。
 *
 *
 *
 * 示例 1：
 *
 * 输入：root = [1,2,3]
 * 输出：6
 * 解释：最优路径是 2 -> 1 -> 3 ，路径和为 2 + 1 + 3 = 6
 *
 * 示例 2：
 *
 * 输入：root = [-10,9,20,null,null,15,7]
 * 输出：42
 * 解释：最优路径是 15 -> 20 -> 7 ，路径和为 15 + 20 + 7 = 42
 *
 *
 *
 *
 * 提示：
 *
 *
 * 树中节点数目范围是 [1, 3 * 10^4]
 * -1000 <= Node.val <= 1000
 *
 *
 */
package lc124

import "math"

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
func maxPathSum(root *TreeNode) int {
	maxsum := math.MinInt

	// dfs获取以当前节点的起点的最大路径和
	// maxsum可以在遍历过程更新
	var dfs func(*TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		// 获取左右子树最大路径和，丢弃负数
		lMax := max(dfs(node.Left), 0)
		rMax := max(dfs(node.Right), 0)
		maxsum = max(maxsum, lMax+rMax+node.Val)
		return node.Val + max(lMax, rMax)
	}
	dfs(root)
	return maxsum
}

// @lc code=end

/*
// @lcpr case=start
// [1,2,3]\n
// @lcpr case=end

// @lcpr case=start
// [-10,9,20,null,null,15,7]\n
// @lcpr case=end

*/
