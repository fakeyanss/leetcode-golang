/*
 * @lc app=leetcode.cn id=86 lang=golang
 *
 * [86] 分隔链表
 *
 * https://leetcode.cn/problems/partition-list/description/
 *
 * algorithms
 * Medium (63.46%)
 * Likes:    593
 * Dislikes: 0
 * Total Accepted:    162.7K
 * Total Submissions: 256.4K
 * Testcase Example:  '[1,4,3,2,5,2]\n3'
 *
 * 给你一个链表的头节点 head 和一个特定值 x ，请你对链表进行分隔，使得所有 小于 x 的节点都出现在 大于或等于 x 的节点之前。
 *
 * 你应当 保留 两个分区中每个节点的初始相对位置。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：head = [1,4,3,2,5,2], x = 3
 * 输出：[1,2,2,4,3,5]
 *
 *
 * 示例 2：
 *
 *
 * 输入：head = [2,1], x = 2
 * 输出：[1,2]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 链表中节点的数目在范围 [0, 200] 内
 * -100
 * -200
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
func partition(head *ListNode, x int) *ListNode {
	l1, l2 := &ListNode{}, &ListNode{}
	dummy := l1
	tmp := l2
	for head != nil {
		if head.Val < x {
			l1.Next = head
			l1 = l1.Next
		} else {
			l2.Next = head
			l2 = l2.Next
		}
		head, head.Next = head.Next, nil

	}
	l1.Next = tmp.Next
	return dummy.Next
}

// @lc code=end
