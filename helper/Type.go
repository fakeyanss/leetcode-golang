package helper

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

// Definition for a Trie Tree.
type Trie struct {
	Val      interface{} // 存储节点的值，可记录字符串key对应的value，实现TrieSet时无用
	Children []*Trie     // 存储指向子节点的指针
}

type TrieMap struct {
	Size int   // map中的trie键值对个数
	Root *Trie // trie树的根节点
}

// 在TrieMap中添加<k,v>
func (t *TrieMap) Put(key string, val interface{}) {
	if !t.ContainsKey(key) {
		t.Size++
	}
	t.Root = t.put(t.Root, key, val, 0)
}

// 额外的辅助函数，递归添加节点
func (t *TrieMap) put(node *Trie, key string, val interface{}, i int) *Trie {
	if node == nil {
		// 如果子节点不存在，新建
		node = new(Trie)
	}
	if i == len(key) {
		// key的路径已经插入完成，将val存入节点
		node.Val = val
		return node
	}
	c := key[i]
	// 递归插入子节点
	node.Children[c] = t.put(node.Children[c], key, val, i+1)
	return node
}

// 删除TrieMap中的<k,v>
func (t *TrieMap) Remove(key string) {
	if !t.ContainsKey(key) {
		return
	}
	t.Root = t.remove(t.Root, key, 0)
	t.Size--
}

// 额外的辅助函数，递归删除节点
func (t *TrieMap) remove(node *Trie, key string, i int) *Trie {
	if node == nil {
		return nil
	}
	if i == len(key) {
		// 找到了key对应的Trie，删除val
		node.Val = nil
	} else {
		c := key[i]
		// 递归去子树删除
		node.Children[c] = t.remove(node.Children[c], key, i+1)
	}

	// 举个例子：
	// 一个TrieMap添加了以下节点：
	// [<them:1>, <zip:2>, <the:3>, <the:4>, <app:5>, <that:6>]
	// 现在删除them，那么应该只删除them的叶子节点。
	// 因为the、th都有key经过，不能删除

	// 后序位置，递归路径上的节点需要被清理
	if node.Val != nil {
		// 如果该TrieNode还存储着Val，说明有一个key对应这个节点
		return node
	}
	// 检查该TrieNode是否还有后缀
	for c := 0; c < 256; c++ {
		if node.Children[c] != nil {
			// 只要存在一个子节点（后缀树枝），就不需要被清理
			return node
		}
	}
	// 既没有存储 val，也没有后缀树枝，则该节点需要被清理
	return nil
}

// 搜索TrieMap中k对应的v，不存在则返回nil
func (t *TrieMap) Get(key string) interface{} {
	// 从root开始搜索key
	x := t.getNode(t.Root, key)
	if x == nil || x.Val == nil {
		// x为空或x.Val为空，都说明key没有对应值
		return nil
	}
	return x.Val
}

// 额外的辅助函数，从节点node开始搜索key，不存在则返回nil
func (t *TrieMap) getNode(node *Trie, key string) *Trie {
	p := node
	// 从node开始搜索key
	for i := 0; i < len(key); i++ {
		if p == nil {
			// 无法向下搜索
			return nil
		}
		// 向下搜索
		c := key[i]
		p = p.Children[c]
	}
	return p
}

// 判断k是否存在与TrieMap中
func (t *TrieMap) ContainsKey(key string) bool {
	return t.Get(key) != nil
}

// 判断TrieMap中是否存在前缀为prefix的k
func (t *TrieMap) HasKeyWithPrefix(prefix string) bool {
	return t.getNode(t.Root, prefix) != nil
}

// 在TrieMap中搜索query的最短前缀
func (t *TrieMap) ShortestPrefixOf(query string) string {
	p := t.Root
	// 从节点node开始搜索key
	for i := 0; i < len(query); i++ {
		if p == nil {
			// 无法向下搜索
			return ""
		}
		if p.Val != nil {
			// 找到第一个k为query的前缀，直接返回
			return query[:i]
		}
		// 向下搜索
		c := query[i]
		p = p.Children[c]
	}

	if p != nil && p.Val != nil {
		// query本身是一个k
		return query
	}
	return ""
}

// 在TrieMap中搜索query的最长前缀
func (t *TrieMap) LongestPrefixOf(query string) string {
	p := t.Root
	maxLen := 0
	for i := 0; i < len(query); i++ {
		if p == nil {
			// 无法向下搜索
			break
		}
		if p.Val != nil {
			// 找到一个k为query的前缀，更新最大长度
			maxLen = i
		}
		// 向下搜索
		c := query[i]
		p = p.Children[c]
	}

	if p != nil && p.Val != nil {
		// query本身是一个k
		return query
	}
	return query[:maxLen]
}

// 在TrieMap中搜索搜索前缀为prefix的k列表
func (t *TrieMap) KeysWithPrefix(prefix string) []string {
	res := []string{}
	// 找到TrieMap中匹配prefix的node
	x := t.getNode(t.Root, prefix)
	if x == nil {
		return res
	}
	// dfs遍历以x为跟的这个子树
	t.traverse(x, &[]byte{}, &res)
	return res
}

// 遍历以node为根的Trie树，找到所有键
func (t *TrieMap) traverse(node *Trie, path *[]byte, res *[]string) {
	if node == nil {
		// 到达叶子结点
		return
	}

	if node.Val != nil {
		// 找到一个k，添加到结果列表
		*res = append(*res, string(*path))
	}
	// 回溯法遍历
	for c := 0; c < 256; c++ {
		*path = append(*path, byte(c))
		t.traverse(node.Children[c], path, res)
		*path = (*path)[:len(*path)-1]
	}
}

// 在TrieMap中搜索前缀为pattern的k列表，支持 . 匹配所有字符
func (t *TrieMap) KeysWithPattern(pattern string) []string {
	return []string{}
}

// 判断TrieMap中是否存在前缀为pattern的k，支持 . 匹配所有字符
func (t *TrieMap) HasKeyWithPattern(pattern string) bool { return false }
