/*
* @lc app=leetcode.cn id=437 lang=golang
* @lcpr version=30305
*
* [437] 路径总和 III
*
* https://leetcode.cn/problems/path-sum-iii/description/
*
  - algorithms
  - Medium (48.29%)
  - Likes:    2188
  - Dislikes: 0
  - Total Accepted:    540.3K
  - Total Submissions: 1.1M
  - Testcase Example:  '[10,5,-3,3,2,null,11,3,-2,null,1]\n' +
    '8\n' +
    '[5,4,8,11,null,13,4,7,2,null,null,5,1]\n' +
    '22'

*
* 给定一个二叉树的根节点 root ，和一个整数 targetSum ，求该二叉树里节点值之和等于 targetSum 的 路径 的数目。
*
* 路径 不需要从根节点开始，也不需要在叶子节点结束，但是路径方向必须是向下的（只能从父节点到子节点）。
*
*
*
* 示例 1：
*
*
*
* 输入：root = [10,5,-3,3,2,null,11,3,-2,null,1], targetSum = 8
* 输出：3
* 解释：和等于 8 的路径有 3 条，如图所示。
*
*
* 示例 2：
*
* 输入：root = [5,4,8,11,null,13,4,7,2,null,null,5,1], targetSum = 22
* 输出：3
*
*
*
*
* 提示:
*
*
* 二叉树的节点个数的范围是 [0,1000]
* -10^9 <= Node.val <= 10^9
* -1000 <= targetSum <= 1000
*
*
*/
package lc437

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
func pathSum(root *TreeNode, targetSum int) int {
    // key为和，value为和为key的路径个数
    presum := make(map[int]int)
    presum[0] = 1
    res := 0

    var dfs func(*TreeNode, int)
    dfs = func(node *TreeNode, sum int) {
        if node == nil {
            return
        }
        sum += node.Val
        res += presum[sum-targetSum]

        presum[sum]++
        dfs(node.Left, sum)
        dfs(node.Right, sum)
        presum[sum]--
    }
    dfs(root, 0)
    return res
}

// @lc code=end

/*
// @lcpr case=start
// [10,5,-3,3,2,null,11,3,-2,null,1]\n8\n
// @lcpr case=end

// @lcpr case=start
// [5,4,8,11,null,13,4,7,2,null,null,5,1]\n22\n
// @lcpr case=end

*/
