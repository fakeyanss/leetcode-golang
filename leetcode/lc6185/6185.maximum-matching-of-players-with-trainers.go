// https://ghproxy.com/https://raw.githubusercontent.com/fakeYanss/imgplace/master//2022/2022-09-18_20220918150655.png

// https://leetcode.cn/problems/maximum-matching-of-players-with-trainers/
package lc6185

import "sort"

// 考虑能力值最小的运动员 player[i]，他应当匹配训练能力值大于等于 player[i] 且最接近 player[i] 的训练师（如果选了一个训练能力值更大的，可能会导致能力值更大的运动员无法匹配）。
// 那么对 players 和 trainers 从小到大排序，然后双指针模拟即可。
func matchPlayersAndTrainers(players []int, trainers []int) int {
	sort.Ints(players)
	sort.Ints(trainers)
	j, n := 0, len(trainers)
	for i, p := range players {
		for j < n && trainers[j] < p {
			j++
		}
		if j == n {
			return i // 未找到匹配
		}
		j++ // 匹配一位
	}
	return len(players) // 全匹配
}
