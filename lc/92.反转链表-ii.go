/*
 * @lc app=leetcode.cn id=92 lang=golang
 *
 * [92] 反转链表 II
 *
 * https://leetcode.cn/problems/reverse-linked-list-ii/description/
 *
 * algorithms
 * Medium (55.47%)
 * Likes:    1352
 * Dislikes: 0
 * Total Accepted:    319.5K
 * Total Submissions: 576K
 * Testcase Example:  '[1,2,3,4,5]\n2\n4'
 *
 * 给你单链表的头指针 head 和两个整数 left 和 right ，其中 left  。请你反转从位置 left 到位置 right 的链表节点，返回
 * 反转后的链表 。
 *
 *
 * 示例 1：
 *
 *
 * 输入：head = [1,2,3,4,5], left = 2, right = 4
 * 输出：[1,4,3,2,5]
 *
 *
 * 示例 2：
 *
 *
 * 输入：head = [5], left = 1, right = 1
 * 输出：[5]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 链表中节点数目为 n
 * 1
 * -500
 * 1
 *
 *
 *
 *
 * 进阶： 你可以使用一趟扫描完成反转吗？
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
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	// base case
	if left == 1 {
		return reverseK(head, right)
	}
	// 前进到反转的起点触发 base case
	head.Next = reverseBetween(head.Next, left-1, right-1)
	return head
}

// 反转前k个节点
func reverseK(head *ListNode, k int) *ListNode {
	// base case
	if k == 1 {
		return head
	}
	// 类似反转整个链表
	last := reverseK(head.Next, k-1)
	// 此时 head.Next.Next 指向第k+1个节点, 应该将 head.Next 指向它
	// successor := head.Next.Next
	// head.Next.Next = head
	// head.Next = successor
	head.Next.Next, head.Next = head, head.Next.Next
	return last
}

// @lc code=end
