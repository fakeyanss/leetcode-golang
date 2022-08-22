package lc

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func NewListNode(val int, next *ListNode) *ListNode {
	return &ListNode{val, next}
}

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func NewTreeNode(val int, l *TreeNode, r *TreeNode) *TreeNode {
	return &TreeNode{val, l, r}
}

/**
 * Definition for a Node.
 */
type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}
