/*
 * @lc app=leetcode.cn id=25 lang=golang
 * @lcpr version=20004
 *
 * [25] K 个一组翻转链表
 *
 * https://leetcode.cn/problems/reverse-nodes-in-k-group/description/
 *
 * algorithms
 * Hard (68.75%)
 * Likes:    2440
 * Dislikes: 0
 * Total Accepted:    700K
 * Total Submissions: 1M
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
 * 输入：head = [1,2,3,4,5], k = 2
 * 输出：[2,1,4,3,5]
 *
 *
 * 示例 2：
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
package lc0025

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
func reverseKGroup(head *ListNode, k int) *ListNode {
	// // 递归思路
	// start, end := head, head
	// for i := 0; i < k; i++ {
	// 	if end == nil {
	// 		return start
	// 	}
	// 	end = end.Next
	// }

	// reverseBetween := func(s, e *ListNode) *ListNode {
	// 	var prev *ListNode
	// 	for s != e {
	// 		nxt := s.Next
	// 		s.Next = prev
	// 		prev = s
	// 		s = nxt
	// 	}
	// 	return prev
	// }

	// newHead := reverseBetween(start, end)
	// start.Next = reverseKGroup(end, k)
	// return newHead

	// 迭代思路
	n := 0
	for cur := head; cur != nil; cur = cur.Next {
		n++
	}

	dummy := &ListNode{Next: head}
	p0 := dummy
	var pre *ListNode
	cur := p0.Next

	// k个一组处理
	for ; n >= k; n -= k {
		for i := 0; i < k; i++ {
			nxt := cur.Next
			cur.Next = pre
			pre = cur
			cur = nxt
		}

		nxt := p0.Next
		p0.Next.Next = cur
		p0.Next = pre
		p0 = nxt
	}
	return dummy.Next
}

// @lc code=end

/*
// @lcpr case=start
// [1,2,3,4,5]\n2\n
// @lcpr case=end

// @lcpr case=start
// [1,2,3,4,5]\n3\n
// @lcpr case=end

*/
