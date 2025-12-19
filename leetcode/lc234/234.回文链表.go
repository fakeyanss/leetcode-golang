/*
 * @lc app=leetcode.cn id=234 lang=golang
 * @lcpr version=30305
 *
 * [234] 回文链表
 *
 * https://leetcode.cn/problems/palindrome-linked-list/description/
 *
 * algorithms
 * Easy (57.73%)
 * Likes:    2169
 * Dislikes: 0
 * Total Accepted:    1.1M
 * Total Submissions: 2M
 * Testcase Example:  '[1,2,2,1]\n[1,2]'
 *
 * 给你一个单链表的头节点 head ，请你判断该链表是否为回文链表。如果是，返回 true ；否则，返回 false 。
 *
 *
 *
 * 示例 1：
 *
 * 输入：head = [1,2,2,1]
 * 输出：true
 *
 *
 * 示例 2：
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

type ListNode struct {
	Val  int
	Next *ListNode
}

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
	// 找到中间点
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	newHead := slow
	if fast != nil {
		newHead = slow.Next
	}

	// 翻转后半段
	var pre *ListNode
	for newHead != nil {
		nxt := newHead.Next
		newHead.Next = pre
		pre = newHead
		newHead = nxt
	}
	newHead = pre

	// 挨个比较
	for newHead != nil { // 奇数个数时，head链表比newHead长度多1
		if head.Val != newHead.Val {
			return false
		}
		head, newHead = head.Next, newHead.Next
	}
	return true
}

// @lc code=end

/*
// @lcpr case=start
// [1,2,2,1]\n
// @lcpr case=end

// @lcpr case=start
// [1,2]\n
// @lcpr case=end

*/
