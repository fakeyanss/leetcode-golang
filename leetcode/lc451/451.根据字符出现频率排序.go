/*
 * @lc app=leetcode.cn id=451 lang=golang
 *
 * [451] 根据字符出现频率排序
 *
 * https://leetcode.cn/problems/sort-characters-by-frequency/description/
 *
 * algorithms
 * Medium (71.97%)
 * Likes:    445
 * Dislikes: 0
 * Total Accepted:    116.3K
 * Total Submissions: 161.6K
 * Testcase Example:  '"tree"'
 *
 * 给定一个字符串 s ，根据字符出现的 频率 对其进行 降序排序 。一个字符出现的 频率 是它出现在字符串中的次数。
 *
 * 返回 已排序的字符串 。如果有多个答案，返回其中任何一个。
 *
 *
 *
 * 示例 1:
 *
 *
 * 输入: s = "tree"
 * 输出: "eert"
 * 解释: 'e'出现两次，'r'和't'都只出现一次。
 * 因此'e'必须出现在'r'和't'之前。此外，"eetr"也是一个有效的答案。
 *
 *
 * 示例 2:
 *
 *
 * 输入: s = "cccaaa"
 * 输出: "cccaaa"
 * 解释: 'c'和'a'都出现三次。此外，"aaaccc"也是有效的答案。
 * 注意"cacaca"是不正确的，因为相同的字母必须放在一起。
 *
 *
 * 示例 3:
 *
 *
 * 输入: s = "Aabb"
 * 输出: "bbAa"
 * 解释: 此外，"bbaA"也是一个有效的答案，但"Aabb"是不正确的。
 * 注意'A'和'a'被认为是两种不同的字符。
 *
 *
 *
 *
 * 提示:
 *
 *
 * 1 <= s.length <= 5 * 10^5
 * s 由大小写英文字母和数字组成
 *
 *
 */
package lc451

import "container/heap"

// @lc code=start
func frequencySort(s string) string {
	// 典型的二叉堆问题
	occurrences := make(map[rune]int)
	for _, val := range s {
		occurrences[val]++
	}

	h := &EntryHeap{}
	heap.Init(h)
	for key, val := range occurrences {
		heap.Push(h, Entry{val: key, freq: val})
	}

	res := []rune{}
	for h.Len() > 0 {
		entry := heap.Pop(h).(Entry)
		for i := 0; i < entry.freq; i++ {
			res = append(res, entry.val)
		}
	}

	return string(res)
}

type Entry struct {
	val  rune
	freq int
}

type EntryHeap []Entry

func (h EntryHeap) Len() int {
	return len(h)
}

func (h EntryHeap) Less(i, j int) bool {
	return h[i].freq > h[j].freq
}

func (h EntryHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *EntryHeap) Push(x interface{}) {
	*h = append(*h, x.(Entry))
}

func (h *EntryHeap) Pop() interface{} {
	old := *h
	x := old[len(old)-1]
	*h = old[:len(old)-1]
	return x
}

// @lc code=end
