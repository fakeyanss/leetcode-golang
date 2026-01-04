/*
 * @lc app=leetcode.cn id=530 lang=golang
 * @lcpr version=30305
 *
 * [530] 二叉搜索树的最小绝对差
 *
 * https://leetcode.cn/problems/minimum-absolute-difference-in-bst/description/
 *
 * algorithms
 * Easy (63.05%)
 * Likes:    649
 * Dislikes: 0
 * Total Accepted:    358.1K
 * Total Submissions: 567.9K
 * Testcase Example:  '[4,2,6,1,3]\n[1,0,48,null,null,12,49]'
 *
 * 给你一个二叉搜索树的根节点 root ，返回 树中任意两不同节点值之间的最小差值 。
 *
 * 差值是一个正数，其数值等于两值之差的绝对值。
 *
 *
 *
 * 示例 1：
 *
 * 输入：root = [4,2,6,1,3]
 * 输出：1
 *
 *
 * 示例 2：
 *
 * 输入：root = [1,0,48,null,null,12,49]
 * 输出：1
 *
 *
 *
 *
 * 提示：
 *
 *
 * 树中节点的数目范围是 [2, 10^4]
 * 0 <= Node.val <= 10^5
 *
 *
 *
 *
 * 注意：本题与 783 https://leetcode.cn/problems/minimum-distance-between-bst-nodes/
 * 相同
 *
 */
package lc530

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
func getMinimumDifference(root *TreeNode) int {
	// 二叉搜索树，中序遍历即升序排列，求相邻两个元素的差的最小值
	res := 100001
	pre := -100001 // 起始值为最小值
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		res = min(res, node.Val-pre)
		pre = node.Val
		dfs(node.Right)
	}
	dfs(root)
	return res
}

// @lc code=end

/*
// @lcpr case=start
// [4,2,6,1,3]\n
// @lcpr case=end

// @lcpr case=start
// [1,0,48,null,null,12,49]\n
// @lcpr case=end

*/
