/**
 * leetcode solution
 */
package lc

func ArrayEqual(a, b []int) bool {
	if (a == nil) != (b == nil) {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for key, val := range a {
		if val != b[key] {
			return false
		}
	}
	return true
}

func PrintListNode(l *ListNode) []int {
	arr := []int{}
	for l != nil {
		arr = append(arr, l.Val)
		l = l.Next
	}
	// fmt.Printf("print all ListNode Val=%d\n", arr)
	return arr
}

func PrintTreeNode(n *TreeNode) []int {
	arr := make([]int, 0)
	var traverse func(n *TreeNode)
	traverse = func(n *TreeNode) {
		if n == nil {
			return
		}
		arr = append(arr, n.Val)
		traverse(n.Left)
		traverse(n.Right)
	}
	traverse(n)
	return arr
}
