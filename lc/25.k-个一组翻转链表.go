/*
 * @lc app=leetcode.cn id=25 lang=golang
 *
 * [25] K 个一组翻转链表
 *
 * https://leetcode.cn/problems/reverse-nodes-in-k-group/description/
 *
 * algorithms
 * Hard (67.58%)
 * Likes:    1729
 * Dislikes: 0
 * Total Accepted:    372.5K
 * Total Submissions: 551.1K
 * Testcase Example:  '[1,2,3,4,5]\n2'
 *
 * 给你链表的头节点 head ，每 k 个节点一组进行翻转，请你返回修改后的链表。
 *
 * k 是一个正整数，它的值小于或等于链表的长度。如果节点总数不是 k 的整数倍，那么请将最后剩余的节点保持原有顺序。
 *
 * 你不能只是单纯的改变节点内部的值，而是需要实际进行节点交换。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：head = [1,2,3,4,5], k = 2
 * 输出：[2,1,4,3,5]
 *
 *
 * 示例 2：
 *
 *
 *
 *
 * 输入：head = [1,2,3,4,5], k = 3
 * 输出：[3,2,1,4,5]
 *
 *
 *
 * 提示：
 *
 *
 * 链表中的节点数目为 n
 * 1 <= k <= n <= 5000
 * 0 <= Node.val <= 1000
 *
 *
 *
 *
 * 进阶：你可以设计一个只用 O(1) 额外内存空间的算法解决此问题吗？
 *
 *
 *
 *
 */
package lc

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
	start, end := head, head
	for cnt := 0; cnt < k; cnt++ {
		// 不足 k 个，不需要反转，base case
		if end == nil {
			return head
		}
		end = end.Next
	}
	// 反转前 k 个元素
	newHead := reverseBetweenListNode(start, end)
	// 递归反转后续链表并连接起来
	start.Next = reverseKGroup(end, k)
	return newHead
}

// 反转区间 [start, end) 的元素，注意是左闭右开
func reverseBetweenListNode(start *ListNode, end *ListNode) *ListNode {
	var pre *ListNode
	cur, next := start, start
	for cur != end {
		next = cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}

// @lc code=end
