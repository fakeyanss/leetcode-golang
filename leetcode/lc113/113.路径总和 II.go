/*
 * @lc app=leetcode.cn id=113 lang=golang
 * @lcpr version=30305
 *
 * [113] 路径总和 II
 *
 * https://leetcode.cn/problems/path-sum-ii/description/
 *
 * algorithms
 * Medium (64.13%)
 * Likes:    1218
 * Dislikes: 0
 * Total Accepted:    498.7K
 * Total Submissions: 777.3K
 * Testcase Example:  '[5,4,8,11,null,13,4,7,2,null,null,5,1]\n22\n[1,2,3]\n5\n[1,2]\n0'
 *
 * 给你二叉树的根节点 root 和一个整数目标和 targetSum ，找出所有 从根节点到叶子节点 路径总和等于给定目标和的路径。
 *
 * 叶子节点 是指没有子节点的节点。
 *
 *
 *
 *
 *
 * 示例 1：
 *
 * 输入：root = [5,4,8,11,null,13,4,7,2,null,null,5,1], targetSum = 22
 * 输出：[[5,4,11,2],[5,8,4,5]]
 *
 *
 * 示例 2：
 *
 * 输入：root = [1,2,3], targetSum = 5
 * 输出：[]
 *
 *
 * 示例 3：
 *
 * 输入：root = [1,2], targetSum = 0
 * 输出：[]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 树中节点总数在范围 [0, 5000] 内
 * -1000 <= Node.val <= 1000
 * -1000 <= targetSum <= 1000
 *
 *
 *
 *
 */
package lc113

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, targetSum int) [][]int {
	var res [][]int
	var track []int
	var dfs func(*TreeNode, int)
	dfs = func(node *TreeNode, target int) {
		if node == nil {
			return
		}

		track = append(track, node.Val)
		defer func() { track = track[:len(track)-1] }()

		if node.Left == nil && node.Right == nil && target == node.Val {
			res = append(res, append([]int{}, track...))
			return
		}

		dfs(node.Left, target-node.Val)
		dfs(node.Right, target-node.Val)
	}
	dfs(root, targetSum)
	return res
}

// @lc code=end

/*
// @lcpr case=start
// [5,4,8,11,null,13,4,7,2,null,null,5,1]\n22\n
// @lcpr case=end

// @lcpr case=start
// [1,2,3]\n5\n
// @lcpr case=end

// @lcpr case=start
// [1,2]\n0\n
// @lcpr case=end

*/
