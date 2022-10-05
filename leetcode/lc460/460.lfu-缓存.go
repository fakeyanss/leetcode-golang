/*
 * @lc app=leetcode.cn id=460 lang=golang
 *
 * [460] LFU 缓存
 *
 * https://leetcode.cn/problems/lfu-cache/description/
 *
 * algorithms
 * Hard (44.38%)
 * Likes:    605
 * Dislikes: 0
 * Total Accepted:    56.6K
 * Total Submissions: 127.5K
 * Testcase Example:  '["LFUCache","put","put","get","put","get","get","put","get","get","get"]\n' +
  '[[2],[1,1],[2,2],[1],[3,3],[2],[3],[4,4],[1],[3],[4]]'
 *
 * 请你为 最不经常使用（LFU）缓存算法设计并实现数据结构。
 *
 * 实现 LFUCache 类：
 *
 *
 * LFUCache(int capacity) - 用数据结构的容量 capacity 初始化对象
 * int get(int key) - 如果键 key 存在于缓存中，则获取键的值，否则返回 -1 。
 * void put(int key, int value) - 如果键 key 已存在，则变更其值；如果键不存在，请插入键值对。当缓存达到其容量
 * capacity 时，则应该在插入新项之前，移除最不经常使用的项。在此问题中，当存在平局（即两个或更多个键具有相同使用频率）时，应该去除 最近最久未使用
 * 的键。
 *
 *
 * 为了确定最不常使用的键，可以为缓存中的每个键维护一个 使用计数器 。使用计数最小的键是最久未使用的键。
 *
 * 当一个键首次插入到缓存中时，它的使用计数器被设置为 1 (由于 put 操作)。对缓存中的键执行 get 或 put 操作，使用计数器的值将会递增。
 *
 * 函数 get 和 put 必须以 O(1) 的平均时间复杂度运行。
 *
 *
 *
 * 示例：
 *
 *
 * 输入：
 * ["LFUCache", "put", "put", "get", "put", "get", "get", "put", "get", "get",
 * "get"]
 * [[2], [1, 1], [2, 2], [1], [3, 3], [2], [3], [4, 4], [1], [3], [4]]
 * 输出：
 * [null, null, null, 1, null, -1, 3, null, -1, 3, 4]
 *
 * 解释：
 * // cnt(x) = 键 x 的使用计数
 * // cache=[] 将显示最后一次使用的顺序（最左边的元素是最近的）
 * LFUCache lfu = new LFUCache(2);
 * lfu.put(1, 1);   // cache=[1,_], cnt(1)=1
 * lfu.put(2, 2);   // cache=[2,1], cnt(2)=1, cnt(1)=1
 * lfu.get(1);      // 返回 1
 * ⁠                // cache=[1,2], cnt(2)=1, cnt(1)=2
 * lfu.put(3, 3);   // 去除键 2 ，因为 cnt(2)=1 ，使用计数最小
 * ⁠                // cache=[3,1], cnt(3)=1, cnt(1)=2
 * lfu.get(2);      // 返回 -1（未找到）
 * lfu.get(3);      // 返回 3
 * ⁠                // cache=[3,1], cnt(3)=2, cnt(1)=2
 * lfu.put(4, 4);   // 去除键 1 ，1 和 3 的 cnt 相同，但 1 最久未使用
 * ⁠                // cache=[4,3], cnt(4)=1, cnt(3)=2
 * lfu.get(1);      // 返回 -1（未找到）
 * lfu.get(3);      // 返回 3
 * ⁠                // cache=[3,4], cnt(4)=1, cnt(3)=3
 * lfu.get(4);      // 返回 4
 * ⁠                // cache=[3,4], cnt(4)=2, cnt(3)=3
 *
 *
 *
 * 提示：
 *
 *
 * 0 <= capacity <= 10^4
 * 0 <= key <= 10^5
 * 0 <= value <= 10^9
 * 最多调用 2 * 10^5 次 get 和 put 方法
 *
 *
*/
package lc460

// @lc code=start
type CacheNode struct {
    k, v       int
    prev, next *CacheNode
    parent     *FreqNode
}

type FreqNode struct {
    freq       int
    head, tail *CacheNode
    prev, next *FreqNode
}

type LFUCache struct {
    cache   map[int]*CacheNode
    lfuHead *FreqNode
    cap     int
    size    int
}

func Constructor(capacity int) LFUCache {
    zeroFreqNode := &FreqNode{
        freq: 0, // head freq is true
    }
    return LFUCache{
        cache:   make(map[int]*CacheNode, capacity),
        lfuHead: zeroFreqNode,
        cap:     capacity,
    }
}

func (lfu *LFUCache) Get(key int) int {
    if found, ok := lfu.cache[key]; ok {
        newCacheNode := found.moveAddOneFreq(key, found.v)
        lfu.cache[key] = newCacheNode
        return found.v
    } else {
        return -1
    }
}

func (lfu *LFUCache) Put(key int, value int) {
    if lfu.cap == 0 {
        return
    }
    if found, ok := lfu.cache[key]; ok {
        newCacheNode := found.moveAddOneFreq(key, value)
        lfu.cache[key] = newCacheNode
    } else {
        if lfu.size >= lfu.cap {
            lfuNode := lfu.lfuHead.next.tail
            delete(lfu.cache, lfuNode.k)
            lfuNode.evictNode()
        } else {
            lfu.size++
        }

        var oneFreqNode *FreqNode
        if lfu.lfuHead.next != nil && lfu.lfuHead.next.freq == 1 {
            oneFreqNode = lfu.lfuHead.next
        } else {
            oneFreqNode = NewFreqNode(1, lfu.lfuHead, lfu.lfuHead.next)
        }

        newCacheNode := &CacheNode{
            k:      key,
            v:      value,
            prev:   nil,
            next:   oneFreqNode.head,
            parent: oneFreqNode,
        }
        lfu.cache[key] = newCacheNode
        if oneFreqNode.head == nil {
            oneFreqNode.tail = newCacheNode
        } else {
            oneFreqNode.head.prev = newCacheNode
        }
        oneFreqNode.head = newCacheNode
    }
}

func NewFreqNode(freq int, prev *FreqNode, next *FreqNode) *FreqNode {
    nn := &FreqNode{
        freq: freq,
        prev: prev,
        next: next,
    }
    if prev != nil {
        prev.next = nn
    }
    if next != nil {
        next.prev = nn
    }
    return nn
}

func (n *CacheNode) moveAddOneFreq(key, val int) *CacheNode {
    var newFreqNode *FreqNode
    curFreqNode := n.parent

    if curFreqNode.next != nil && curFreqNode.next.freq == (curFreqNode.freq+1) {
        // 下一个freqNode正好是freq+1
        newFreqNode = curFreqNode.next
    } else {
        // freq+1正好位于当前的freqNode和下一个freqNode之间
        newFreqNode = NewFreqNode(curFreqNode.freq+1, curFreqNode, curFreqNode.next)
    }

    newCacheNode := &CacheNode{
        k:      key,
        v:      val,
        next:   newFreqNode.head,
        parent: newFreqNode,
    }
    // 将cacheNode记录到freqNode
    if newFreqNode.head == nil {
        newFreqNode.tail = newCacheNode
    } else {
        newFreqNode.head.prev = newCacheNode
    }

    newFreqNode.head = newCacheNode

    n.evictNode()
    return newCacheNode
}

func (node *CacheNode) evictNode() {
    // remove the cache node from the linkedList
    // remove the freqNode(parent) if it is empty
    // do nothing to the map
    curFreqNode := node.parent
    if node.prev != nil {
        node.prev.next = node.next
    } else {
        curFreqNode.head = node.next
    }

    if node.next != nil {
        node.next.prev = node.prev
    } else {
        curFreqNode.tail = node.prev
    }

    if curFreqNode.head == nil {
        curFreqNode.removeFreqNode()
    }
}

func (node *FreqNode) removeFreqNode() {
    if node.freq == 0 {
        panic("should not remove the head")
    }
    node.prev.next = node.next
    if node.next != nil {
        node.next.prev = node.prev
    }
}

/**
 * Your LFUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
// @lc code=end
