// https://leetcode.cn/problems/sum-of-prefix-scores-of-strings/solution/by-lfool-w82u/

package lc6183

func sumPrefixScores(words []string) []int {
	// 构建前缀树
	root := &TrieNode{}
	for i, word := range words {
		cur := root
		for _, c := range word {
			c -= 'a'
			if cur.children[c] == nil {
				cur.children[c] = &TrieNode{}
			}
			cur = cur.children[c]
			cur.score++
		}
		cur.ids = append(cur.ids, i)
	}

	// 遍历查询
	ans := make([]int, len(words))
	var dfs func(*TrieNode, int)
	dfs = func(node *TrieNode, sum int) {
		sum += node.score
		for _, i := range node.ids {
			ans[i] += sum
		}
		for _, child := range node.children {
			if child != nil {
				dfs(child, sum)
			}
		}
	}
	dfs(root, 0)
	return ans
}

type TrieNode struct {
	children [26]*TrieNode
	ids      []int // 由于可能有相同字符串，每个字符串对应的节点需要用一个列表存储该字符串在 words 中的所有下标
	score    int   // 经过该节点的字符串个数，即以该节点为前缀的字符串的个数
}
