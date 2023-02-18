/*
 * @lc app=leetcode.cn id=21 lang=golang
 *
 * [21] 合并两个有序链表
 *
 * https://leetcode.cn/problems/merge-two-sorted-lists/description/
 *
 * algorithms
 * Easy (66.72%)
 * Likes:    2499
 * Dislikes: 0
 * Total Accepted:    1.1M
 * Total Submissions: 1.6M
 * Testcase Example:  '[1,2,4]\n[1,3,4]'
 *
 * 将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：l1 = [1,2,4], l2 = [1,3,4]
 * 输出：[1,1,2,3,4,4]
 *
 *
 * 示例 2：
 *
 *
 * 输入：l1 = [], l2 = []
 * 输出：[]
 *
 *
 * 示例 3：
 *
 *
 * 输入：l1 = [], l2 = [0]
 * 输出：[0]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 两个链表的节点数目范围是 [0, 50]
 * -100
 * l1 和 l2 均按 非递减顺序 排列
 *
 *
 */
package lc0021

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
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	// 递归
	// if list1 == nil {
	// 	return list2
	// }
	// if list2 == nil {
	// 	return list1
	// }
	// if list1.Val < list2.Val {
	// 	list1.Next = mergeTwoLists(list1.Next, list2)
	// 	return list1
	// }
	// list2.Next = mergeTwoLists(list1, list2.Next)
	// return list2

	// 迭代
	res := &ListNode{}
	cursor := res
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			cursor.Next = list1
			cursor = list1
			list1 = list1.Next
		} else {
			cursor.Next = list2
			cursor = list2
			list2 = list2.Next
		}
	}
	if list1 == nil {
		cursor.Next = list2
	} else {
		cursor.Next = list1
	}
	return res.Next
}

// @lc code=end
