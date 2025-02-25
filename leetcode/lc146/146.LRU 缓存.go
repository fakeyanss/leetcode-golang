/*
* @lc app=leetcode.cn id=146 lang=golang
* @lcpr version=20004
*
* [146] LRU 缓存
*
* https://leetcode.cn/problems/lru-cache/description/
*
  - algorithms
  - Medium (54.07%)
  - Likes:    3360
  - Dislikes: 0
  - Total Accepted:    755.3K
  - Total Submissions: 1.4M
  - Testcase Example:  '["LRUCache","put","put","get","put","get","put","get","get","get"]\n' +
    '[[2],[1,1],[2,2],[1],[3,3],[2],[4,4],[1],[3],[4]]'

*
* 请你设计并实现一个满足  LRU (最近最少使用) 缓存 约束的数据结构。
*
* 实现 LRUCache 类：
*
*
*
*
* LRUCache(int capacity) 以 正整数 作为容量 capacity 初始化 LRU 缓存
* int get(int key) 如果关键字 key 存在于缓存中，则返回关键字的值，否则返回 -1 。
* void put(int key, int value) 如果关键字 key 已经存在，则变更其数据值 value ；如果不存在，则向缓存中插入该组
* key-value 。如果插入操作导致关键字数量超过 capacity ，则应该 逐出 最久未使用的关键字。
*
*
* 函数 get 和 put 必须以 O(1) 的平均时间复杂度运行。
*
*
*
*
*
* 示例：
*
* 输入
* ["LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"]
* [[2], [1, 1], [2, 2], [1], [3, 3], [2], [4, 4], [1], [3], [4]]
* 输出
* [null, null, null, 1, null, -1, null, -1, 3, 4]
*
* 解释
* LRUCache lRUCache = new LRUCache(2);
* lRUCache.put(1, 1); // 缓存是 {1=1}
* lRUCache.put(2, 2); // 缓存是 {1=1, 2=2}
* lRUCache.get(1);    // 返回 1
* lRUCache.put(3, 3); // 该操作会使得关键字 2 作废，缓存是 {1=1, 3=3}
* lRUCache.get(2);    // 返回 -1 (未找到)
* lRUCache.put(4, 4); // 该操作会使得关键字 1 作废，缓存是 {4=4, 3=3}
* lRUCache.get(1);    // 返回 -1 (未找到)
* lRUCache.get(3);    // 返回 3
* lRUCache.get(4);    // 返回 4
*
*
*
*
* 提示：
*
*
* 1 <= capacity <= 3000
* 0 <= key <= 10000
* 0 <= value <= 10^5
* 最多调用 2 * 10^5 次 get 和 put
*
*
*/
package lc146

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start

type LRUCache struct {
	capacity int
	dummy    *Node // 哨兵节点
	cache    map[int]*Node
}

type Node struct {
	prev, next *Node
	key, val   int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		dummy: func() *Node {
			dummy := &Node{}
			dummy.prev, dummy.next = dummy, dummy
			return dummy
		}(),
		cache: make(map[int]*Node, capacity),
	}
}

func (c *LRUCache) Get(key int) int {
	if node := c.getNode(key); node == nil {
		return -1
	} else {
		return node.val
	}
}

func (c *LRUCache) Put(key int, value int) {
	node := c.getNode(key)
	if node != nil {
		node.val = value // 更新
		return
	}
	node = &Node{key: key, val: value} // 插入
	c.cache[key] = node
	c.pushFront(node)
	if len(c.cache) > c.capacity {
		tailNode := c.dummy.prev
		delete(c.cache, tailNode.key)
		c.remove(tailNode)
	}

}

func (c *LRUCache) remove(node *Node) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (c *LRUCache) pushFront(node *Node) {
	node.prev = c.dummy
	node.next = c.dummy.next
	node.prev.next = node
	node.next.prev = node
}

func (c *LRUCache) getNode(key int) *Node {
	if node, ok := c.cache[key]; !ok {
		return nil
	} else {
		// 先删除再移动到链表头
		c.remove(node)
		c.pushFront(node)
		return node
	}
}

// type LRUCache struct {
// 	capacity  int
// 	list      *list.List // 双向链表
// 	keyToNode map[int]*list.Element
// }

// type entry struct {
// 	key, value int
// }

// func Constructor(capacity int) LRUCache {
// 	return LRUCache{capacity, list.New(), map[int]*list.Element{}}
// }

// func (c *LRUCache) Get(key int) int {
// 	node := c.keyToNode[key]
// 	if node == nil {
// 		return -1
// 	}
// 	c.list.MoveToFront(node) // 移动到双链表头部
// 	return node.Value.(entry).value
// }

// func (c *LRUCache) Put(key int, value int) {
// 	if node := c.keyToNode[key]; node != nil {
// 		node.Value = entry{key, value} // 更新
// 		c.list.MoveToFront(node)       // 移动到双链表头部
// 		return
// 	}
// 	c.keyToNode[key] = c.list.PushFront(entry{key, value})
// 	if len(c.keyToNode) > c.capacity {
// 		delete(c.keyToNode, c.list.Remove(c.list.Back()).(entry).key) // 去掉双调表尾部的节点
// 	}
// }

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
// @lc code=end
