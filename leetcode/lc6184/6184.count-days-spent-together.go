// https://ghproxy.com/https://raw.githubusercontent.com/fakeYanss/imgplace/master//2022/2022-09-18_20220918143300.png

// https://leetcode.cn/problems/count-days-spent-together/
package lc6184

var days = []int{0, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

func init() {
	for i := 1; i <= 12; i++ {
		days[i] = days[i-1] + days[i]
	}
}

func countDaysTogether(arriveAlice, leaveAlice, arriveBob, leaveBob string) int {
	leave := calcDays(min(leaveAlice, leaveBob))
	arrive := calcDays(max(arriveAlice, arriveBob))
	ans := leave - arrive + 1
	if ans < 0 {
		ans = 0
	}
	return ans
}

func calcDays(s string) (day int) {
	end := s[0]&15*10 + s[1]&15 - 1 // '1'&15=1
	day = days[end]
	return day + int(s[3]&15*10+s[4]&15)
}

func min(a, b string) string {
	if b < a {
		return b
	}
	return a
}

func max(a, b string) string {
	if b > a {
		return b
	}
	return a
}
