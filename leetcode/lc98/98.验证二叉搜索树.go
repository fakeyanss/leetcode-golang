/*
 * @lc app=leetcode.cn id=98 lang=golang
 * @lcpr version=20004
 *
 * [98] 验证二叉搜索树
 *
 * https://leetcode.cn/problems/validate-binary-search-tree/description/
 *
 * algorithms
 * Medium (38.69%)
 * Likes:    2481
 * Dislikes: 0
 * Total Accepted:    1.1M
 * Total Submissions: 2.7M
 * Testcase Example:  '[2,1,3]'
 *
 * 给你一个二叉树的根节点 root ，判断其是否是一个有效的二叉搜索树。
 *
 * 有效 二叉搜索树定义如下：
 *
 *
 * 节点的左子树只包含 小于 当前节点的数。
 * 节点的右子树只包含 大于 当前节点的数。
 * 所有左子树和右子树自身必须也是二叉搜索树。
 *
 *
 *
 *
 * 示例 1：
 *
 * 输入：root = [2,1,3]
 * 输出：true
 *
 *
 * 示例 2：
 *
 * 输入：root = [5,1,4,null,null,3,6]
 * 输出：false
 * 解释：根节点的值是 5 ，但是右子节点的值是 4 。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 树中节点数目范围在[1, 10^4] 内
 * -2^31 <= Node.val <= 2^31 - 1
 *
 *
 */
package lc98

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
func isValidBST(root *TreeNode) bool {
	var isValid func(node *TreeNode, minV, maxV int) bool
	isValid = func(node *TreeNode, minV, maxV int) bool {
		if node == nil {
			return true
		}
		if node.Val <= minV || node.Val >= maxV {
			return false
		}
		return isValid(node.Left, minV, node.Val) && isValid(node.Right, node.Val, maxV)
	}
	return isValid(root, math.MinInt, math.MaxInt)
}

// @lc code=end

/*
// @lcpr case=start
// [2,1,3]\n
// @lcpr case=end

// @lcpr case=start
// [5,1,4,null,null,3,6]\n
// @lcpr case=end

*/
