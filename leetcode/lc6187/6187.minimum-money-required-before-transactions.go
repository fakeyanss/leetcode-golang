// https://leetcode.cn/problems/minimum-money-required-before-transactions/
package lc6187

func minimumMoney(transactions [][]int) int64 {
	// totalLose, mx := 0, 0
	// for _, t := range transactions {
	// 	totalLose += max(t[0]-t[1], 0)
	// 	mx = max(mx, min(t[0], t[1]))
	// }
	// return int64(totalLose + mx)

	// 首先把所有 亏损 的项目求个总亏损额 S，
	// 然后逐个考虑交易，
	// 对于交易 A = [cost, back]，
	// 如果它亏损，需要在总和 S 中减去 A 的亏损额，得到其他项目的总亏损额 S1，此时初始钱数至少为 cost + S1；
	// 如果它不亏损，那么无需在总和 S 中减去，初始钱数至少为 cost + S。答案就是对每个交易求得的最大值
	totalLost := 0
	for _, t := range transactions {
		totalLost += max(t[0]-t[1], 0)
	}
	res := 0
	for _, t := range transactions {
		tmp := totalLost
		if t[0] > t[1] {
			tmp -= t[0] - t[1]
		}
		res = max(res, tmp+t[0])
	}
	return int64(res)
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}
func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}
