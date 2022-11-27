/*
 * @lc app=leetcode.cn id=1104 lang=golang
 *
 * [1104] 二叉树寻路
 *
 * https://leetcode.cn/problems/path-in-zigzag-labelled-binary-tree/description/
 *
 * algorithms
 * Medium (75.98%)
 * Likes:    188
 * Dislikes: 0
 * Total Accepted:    31.8K
 * Total Submissions: 41.8K
 * Testcase Example:  '14'
 *
 * 在一棵无限的二叉树上，每个节点都有两个子节点，树中的节点 逐行 依次按 “之” 字形进行标记。
 *
 * 如下图所示，在奇数行（即，第一行、第三行、第五行……）中，按从左到右的顺序进行标记；
 *
 * 而偶数行（即，第二行、第四行、第六行……）中，按从右到左的顺序进行标记。
 *
 *
 *
 * 给你树上某一个节点的标号 label，请你返回从根节点到该标号为 label 节点的路径，该路径是由途经的节点标号所组成的。
 *
 *
 *
 * 示例 1：
 *
 * 输入：label = 14
 * 输出：[1,3,4,14]
 *
 *
 * 示例 2：
 *
 * 输入：label = 26
 * 输出：[1,2,6,10,26]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= label <= 10^6
 *
 *
 */
package lc1104

import (
	"math"
	"sort"
)

// @lc code=start
func pathInZigZagTree(label int) []int {
	// 二叉堆变化。
	// 二叉堆用数组存储，
	// 一个节点下标x，
	// 父节点下标为x/2，
	// 左孩子为x*2，
	// 右孩子为x*2+1

	path := []int{}
	for label >= 1 {
		path = append(path, label)
		label = label / 2

		depth := log(label)
		rang := getLevelRange(depth)
		// 找到父节点
		label = rang[0] + rang[1] - label
	}
	// 反转顺序。正确做法应该是下边反转，但是在这个二叉堆中，每一层的节点只出现一个，所以整体是有序的，只需要重排序即可
	sort.Ints(path)
	return path
}

// 根据下表求当前层级深度
func log(x int) int {
	return int(math.Log(float64(x)) / math.Ln2)
}

// 获取这一层的下标最小值和最大值
func getLevelRange(n int) []int {
	p := int(math.Pow(2, float64(n)))
	return []int{p, 2*p - 1}
}

// @lc code=end
