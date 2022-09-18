// https://leetcode.cn/problems/smallest-even-multiple/
package lc6180

func smallestEvenMultiple(n int) int {
	return lcm(n, 2)
}

// 最小公倍数 = 两数乘积 / 两数的最大公约数(gcd)
func lcm(x, y int) int {
	return x * y / gcd(x, y)
}

func gcd(x, y int) int {
	if y == 0 {
		return x
	}
	return gcd(y, x%y)
}
