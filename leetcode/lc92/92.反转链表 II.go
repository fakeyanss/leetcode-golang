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
	var reverseK func(head *ListNode, k int) *ListNode // 翻转前k个节点
	reverseK = func(head *ListNode, k int) *ListNode {
		if k == 1 {
			return head
		}
		newHead := reverseK(head.Next, k-1)
		// 此时 head.Next.Next 指向第k+1个节点, 应该将 head.Next 指向它
		// 1-2-3-4-5, k=4 --> 1-2, 4-3-2-5  --> 4-3-2-1-5
		head.Next.Next, head.Next = head, head.Next.Next
		return newHead
	}

	// base case
	if left == 1 {
		return reverseK(head, right)
	}
	// 前进到反转的起点触发base case
	head.Next = reverseBetween(head.Next, left-1, right-1)
	return head
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
