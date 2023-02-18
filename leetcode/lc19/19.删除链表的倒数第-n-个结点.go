/*
 * @lc app=leetcode.cn id=19 lang=golang
 *
 * [19] 删除链表的倒数第 N 个结点
 *
 * https://leetcode.cn/problems/remove-nth-node-from-end-of-list/description/
 *
 * algorithms
 * Medium (44.36%)
 * Likes:    2135
 * Dislikes: 0
 * Total Accepted:    856.3K
 * Total Submissions: 1.9M
 * Testcase Example:  '[1,2,3,4,5]\n2'
 *
 * 给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：head = [1,2,3,4,5], n = 2
 * 输出：[1,2,3,5]
 *
 *
 * 示例 2：
 *
 *
 * 输入：head = [1], n = 1
 * 输出：[]
 *
 *
 * 示例 3：
 *
 *
 * 输入：head = [1,2], n = 1
 * 输出：[1]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 链表中结点的数目为 sz
 * 1 <= sz <= 30
 * 0 <= Node.val <= 100
 * 1 <= n <= sz
 *
 *
 *
 *
 * 进阶：你能尝试使用一趟扫描实现吗？
 *
 */
package lc0019

import "github.com/fakeyanss/leetcode-golang/helper"

type ListNode = helper.ListNode

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{0, head}
	// 双指针，找正数第len+1-n个
	p, q := head, dummy
	for i := 0; i < n; i++ {
		p = p.Next
	}
	for p != nil {
		p = p.Next
		q = q.Next
	}
	// q 现在指向第 len-n 个节点，即倒数第 n-1 个节点, 删除 q.next
	q.Next = q.Next.Next
	return dummy.Next
}

// @lc code=end
