/*
 * @lc app=leetcode.cn id=234 lang=golang
 *
 * [234] 回文链表
 *
 * https://leetcode.cn/problems/palindrome-linked-list/description/
 *
 * algorithms
 * Easy (52.13%)
 * Likes:    1452
 * Dislikes: 0
 * Total Accepted:    472.5K
 * Total Submissions: 906.4K
 * Testcase Example:  '[1,2,2,1]'
 *
 * 给你一个单链表的头节点 head ，请你判断该链表是否为回文链表。如果是，返回 true ；否则，返回 false 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：head = [1,2,2,1]
 * 输出：true
 *
 *
 * 示例 2：
 *
 *
 * 输入：head = [1,2]
 * 输出：false
 *
 *
 *
 *
 * 提示：
 *
 *
 * 链表中节点数目在范围[1, 10^5] 内
 * 0 <= Node.val <= 9
 *
 *
 *
 *
 * 进阶：你能否用 O(n) 时间复杂度和 O(1) 空间复杂度解决此题？
 *
 */
package lc234

import "github.com/fakeYanss/leetcode-golang/helper"

type ListNode = helper.ListNode

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome234(head *ListNode) bool {
	// func isPalindrome(head *ListNode) bool {
	// 先找倒中间节点
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// 如果是长度是奇数，将slow后移一位
	if fast != nil {
		slow = slow.Next
	}

	// 反转后半段
	tail := reverseLinkedList(slow)
	for tail != nil {
		if tail.Val != head.Val {
			return false
		}
		head = head.Next
		tail = tail.Next
	}
	return true
}

func reverseLinkedList(node *ListNode) *ListNode {
	var pre *ListNode
	cur := node
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}

// @lc code=end
