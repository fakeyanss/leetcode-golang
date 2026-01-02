/*
 * @lc app=leetcode.cn id=92 lang=golang
 * @lcpr version=20004
 *
 * [92] 反转链表 II
 *
 * https://leetcode.cn/problems/reverse-linked-list-ii/description/
 *
 * algorithms
 * Medium (56.95%)
 * Likes:    1902
 * Dislikes: 0
 * Total Accepted:    586K
 * Total Submissions: 1M
 * Testcase Example:  '[1,2,3,4,5]\n2\n4'
 *
 * 给你单链表的头指针 head 和两个整数 left 和 right ，其中 left <= right 。请你反转从位置 left 到位置 right
 * 的链表节点，返回 反转后的链表 。
 *
 *
 * 示例 1：
 *
 * 输入：head = [1,2,3,4,5], left = 2, right = 4
 * 输出：[1,4,3,2,5]
 *
 *
 * 示例 2：
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
 * 1 <= n <= 500
 * -500 <= Node.val <= 500
 * 1 <= left <= right <= n
 *
 *
 *
 *
 * 进阶： 你可以使用一趟扫描完成反转吗？
 *
 */
package lc92

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
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	dummy := &ListNode{Next: head}
	// 1. 将p移动到left的前一个节点
	p := dummy
	for range left - 1 {
		p = p.Next
	}
	// 1-2-3-4-5
	// |
	// p
	// 找到p的位置

	// 2. 反转[left,right]的节点
	cur := p.Next
	var pre *ListNode
	for range right - left + 1 {
		nxt := cur.Next
		cur.Next = pre
		pre = cur
		cur = nxt
	}
	// 1-2 4-3-2 5
	// |   |     |
	// p  pre    cur
	// 反转后节点的指向有变化

	// 3. 调整节点指向
	p.Next.Next = cur
	p.Next = pre

	return dummy.Next
}

// @lc code=end

/*
// @lcpr case=start
// [1,2,3,4,5]\n2\n4\n
// @lcpr case=end

// @lcpr case=start
// [5]\n1\n1\n
// @lcpr case=end

*/
