/*
 * @lc app=leetcode.cn id=23 lang=golang
 * @lcpr version=20004
 *
 * [23] 合并 K 个升序链表
 *
 * https://leetcode.cn/problems/merge-k-sorted-lists/description/
 *
 * algorithms
 * Hard (60.76%)
 * Likes:    2918
 * Dislikes: 0
 * Total Accepted:    901.5K
 * Total Submissions: 1.5M
 * Testcase Example:  '[[1,4,5],[1,3,4],[2,6]]'
 *
 * 给你一个链表数组，每个链表都已经按升序排列。
 *
 * 请你将所有链表合并到一个升序链表中，返回合并后的链表。
 *
 *
 *
 * 示例 1：
 *
 * 输入：lists = [[1,4,5],[1,3,4],[2,6]]
 * 输出：[1,1,2,3,4,4,5,6]
 * 解释：链表数组如下：
 * [
 * ⁠ 1->4->5,
 * ⁠ 1->3->4,
 * ⁠ 2->6
 * ]
 * 将它们合并到一个有序链表中得到。
 * 1->1->2->3->4->4->5->6
 *
 *
 * 示例 2：
 *
 * 输入：lists = []
 * 输出：[]
 *
 *
 * 示例 3：
 *
 * 输入：lists = [[]]
 * 输出：[]
 *
 *
 *
 *
 * 提示：
 *
 *
 * k == lists.length
 * 0 <= k <= 10^4
 * 0 <= lists[i].length <= 500
 * -10^4 <= lists[i][j] <= 10^4
 * lists[i] 按 升序 排列
 * lists[i].length 的总和不超过 10^4
 *
 *
 */
package lc23

import "container/heap"

// @lcpr-template-start
type ListNode struct {
	Val  int
	Next *ListNode
}

// @lcpr-template-end
// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

//  分治
// func mergeKLists(lists []*ListNode) *ListNode {
// 	n := len(lists)
// 	if n == 0 {
// 		return nil
// 	}
// 	if n == 1 {
// 		return lists[0]
// 	}
// 	left := mergeKLists(lists[:n/2])
// 	right := mergeKLists(lists[n/2:])
// 	return mergeTwoLists(left, right)
// }

// func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
// 	dummy := &ListNode{}
// 	cur := dummy
// 	for list1 != nil && list2 != nil {
// 		if list1.Val < list2.Val {
// 			cur.Next = list1
// 			list1 = list1.Next
// 		} else {
// 			cur.Next = list2
// 			list2 = list2.Next
// 		}
// 		cur = cur.Next
// 	}
// 	if list1 == nil {
// 		cur.Next = list2
// 	}
// 	if list2 == nil {
// 		cur.Next = list1
// 	}
// 	return dummy.Next
// }

// 最小堆
type hp []*ListNode

func (h hp) Len() int           { return len(h) }
func (h hp) Less(i, j int) bool { return h[i].Val < h[j].Val }
func (h hp) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(x any)        { *h = append(*h, x.(*ListNode)) }
func (h *hp) Pop() any          { a := *h; x := a[len(a)-1]; *h = a[:len(a)-1]; return x }

func mergeKLists(lists []*ListNode) *ListNode {
	h := hp{}
	for _, head := range lists {
		if head != nil {
			h = append(h, head)
		}
	}
	heap.Init(&h)

	dummy := &ListNode{}
	cur := dummy
	for h.Len() > 0 {
		n := heap.Pop(&h).(*ListNode)
		if n.Next != nil {
			heap.Push(&h, n.Next)
		}
		cur.Next = n
		cur = cur.Next
	}
	return dummy.Next
}

// @lc code=end

/*
// @lcpr case=start
// [[1,4,5],[1,3,4],[2,6]]\n
// @lcpr case=end

// @lcpr case=start
// []\n
// @lcpr case=end

// @lcpr case=start
// [[]]\n
// @lcpr case=end

*/
