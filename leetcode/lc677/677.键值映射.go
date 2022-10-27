/*
 * @lc app=leetcode.cn id=677 lang=golang
 *
 * [677] 键值映射
 *
 * https://leetcode.cn/problems/map-sum-pairs/description/
 *
 * algorithms
 * Medium (65.97%)
 * Likes:    221
 * Dislikes: 0
 * Total Accepted:    44.9K
 * Total Submissions: 68.1K
 * Testcase Example:  '["MapSum","insert","sum","insert","sum"]\n' +
  '[[],["apple",3],["ap"],["app",2],["ap"]]'
 *
 * 设计一个 map ，满足以下几点:
 *
 *
 * 字符串表示键，整数表示值
 * 返回具有前缀等于给定字符串的键的值的总和
 *
 *
 * 实现一个 MapSum 类：
 *
 *
 * MapSum() 初始化 MapSum 对象
 * void insert(String key, int val) 插入 key-val 键值对，字符串表示键 key ，整数表示值 val 。如果键
 * key 已经存在，那么原来的键值对 key-value 将被替代成新的键值对。
 * int sum(string prefix) 返回所有以该前缀 prefix 开头的键 key 的值的总和。
 *
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：
 * ["MapSum", "insert", "sum", "insert", "sum"]
 * [[], ["apple", 3], ["ap"], ["app", 2], ["ap"]]
 * 输出：
 * [null, null, 3, null, 5]
 *
 * 解释：
 * MapSum mapSum = new MapSum();
 * mapSum.insert("apple", 3);
 * mapSum.sum("ap");           // 返回 3 (apple = 3)
 * mapSum.insert("app", 2);
 * mapSum.sum("ap");           // 返回 5 (apple + app = 3 + 2 = 5)
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= key.length, prefix.length <= 50
 * key 和 prefix 仅由小写英文字母组成
 * 1 <= val <= 1000
 * 最多调用 50 次 insert 和 sum
 *
 *
*/
package lc677

// @lc code=start
type MapSum struct {
    TrieMap
}

type Trie struct {
    Val      int
    Children [26]*Trie
}

type TrieMap struct {
    Size int
    Root *Trie
}

func Constructor() MapSum {
    return MapSum{}
}

func (t *MapSum) getNode(node *Trie, key string) *Trie {
    p := node
    for i := 0; i < len(key); i++ {
        if p == nil {
            return nil
        }
        p = p.Children[key[i]-'a']
    }
    return p
}

func (t *MapSum) put(node *Trie, key string, val int, i int) *Trie {
    if node == nil {
        node = new(Trie)
    }
    if i == len(key) {
        node.Val = val
        return node
    }
    c := key[i] - 'a'
    node.Children[c] = t.put(node.Children[c], key, val, i+1)
    return node
}

func (t *MapSum) get(key string) int {
    x := t.getNode(t.Root, key)
    if x == nil || x.Val == 0 {
        return 0
    }
    return x.Val
}

func (t *MapSum) Insert(key string, val int) {
    if t.get(key) == 0 {
        t.Size++
    }
    t.Root = t.put(t.Root, key, val, 0)
}

func (t *MapSum) Sum(prefix string) int {
    res := 0
    keys := t.keysWithPrefix(prefix)
    for _, key := range keys {
        res += t.get(key)
    }
    return res
}

func (t *MapSum) keysWithPrefix(prefix string) []string {
    res := []string{}
    // 找到TrieMap中匹配prefix的node
    x := t.getNode(t.Root, prefix)
    if x == nil {
        return res
    }
    // dfs遍历以x为跟的这个子树
    path := []byte(prefix)
    t.traverse(x, &path, &res)
    return res
}

func (t *MapSum) traverse(node *Trie, path *[]byte, keys *[]string) {
    if node == nil {
        return
    }
    if node.Val != 0 {
        *keys = append(*keys, string(*path))
    }
    // backtracking
    for i := 0; i < 26; i++ {
        *path = append(*path, byte(i)+'a')
        t.traverse(node.Children[i], path, keys)
        *path = (*path)[:len(*path)-1]
    }
}

/**
 * Your MapSum object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(key,val);
 * param_2 := obj.Sum(prefix);
 */
// @lc code=end
