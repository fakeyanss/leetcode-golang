package lc

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func NewListNode(val int, next *ListNode) *ListNode {
	return &ListNode{val, next}
}

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func NewTreeNode(val int, l *TreeNode, r *TreeNode) *TreeNode {
	return &TreeNode{val, l, r}
}

// Definition for a Node.
type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

// This is the interface that allows for creating nested lists.
// You should not implement it, or speculate about its implementation
type NestedInteger struct {
	Num int
	Ns  []*NestedInteger
}

// Return true if this NestedInteger holds a single integer, rather than a nested list.
func (ni NestedInteger) IsInteger() bool {
	return ni.Ns == nil
}

// Return the single integer that this NestedInteger holds, if it holds a single integer
// The result is undefined if this NestedInteger holds a nested list
// So before calling this method, you should have a check
func (ni NestedInteger) GetInteger() int {
	return ni.Num
}

// Set this NestedInteger to hold a single integer.
func (ni *NestedInteger) SetInteger(value int) {
	ni.Num = value
}

// Set this NestedInteger to hold a nested list and adds a nested integer to it.
func (ni *NestedInteger) Add(elem NestedInteger) {
	ni.Ns = append(ni.Ns, &elem)
}

// Return the nested list that this NestedInteger holds, if it holds a nested list
// The list length is zero if this NestedInteger holds a single integer
// You can access NestedInteger's List element directly if you want to modify it
func (ni NestedInteger) GetList() []*NestedInteger {
	return ni.Ns
}
