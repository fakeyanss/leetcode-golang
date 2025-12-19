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
	dummy    *Node
	cache    map[int]*Node
}

type Node struct {
	key, val   int
	prev, next *Node
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		dummy:    func() *Node { d := &Node{}; d.prev, d.next = d, d; return d }(),
		cache:    make(map[int]*Node),
	}
}

func (c *LRUCache) Get(key int) int {
	if n := c.getNode(key); n != nil {
		return n.val
	} else {
		return -1
	}
}

func (c *LRUCache) Put(key int, value int) {
	if n := c.getNode(key); n != nil {
		n.val = value
	} else {
		newNode := &Node{key: key, val: value}
		for len(c.cache) >= c.capacity {
			last := c.dummy.prev
			c.removeLink(last)
			delete(c.cache, last.key)
		}
		c.pushToFront(newNode)
		c.cache[key] = newNode
	}
}

func (c *LRUCache) getNode(key int) *Node {
	if n, ok := c.cache[key]; ok {
		c.removeLink(n)
		c.pushToFront(n)
		return n
	} else {
		return nil
	}
}

func (c *LRUCache) pushToFront(n *Node) {
	n.prev = c.dummy
	n.next = c.dummy.next
	n.prev.next = n
	n.next.prev = n
}

func (c *LRUCache) removeLink(n *Node) {
	n.prev.next = n.next
	n.next.prev = n.prev
	n.prev, n.next = nil, nil
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
// @lc code=end
