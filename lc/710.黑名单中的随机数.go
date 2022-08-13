/*
 * @lc app=leetcode.cn id=710 lang=golang
 *
 * [710] 黑名单中的随机数
 *
 * https://leetcode.cn/problems/random-pick-with-blacklist/description/
 *
 * algorithms
 * Hard (44.18%)
 * Likes:    209
 * Dislikes: 0
 * Total Accepted:    25.4K
 * Total Submissions: 57.4K
 * Testcase Example:  '["Solution","pick","pick","pick","pick","pick","pick","pick"]\n' +
  '[[7,[2,3,5]],[],[],[],[],[],[],[]]'
 *
 * 给定一个整数 n 和一个 无重复 黑名单整数数组 blacklist 。设计一种算法，从 [0, n - 1] 范围内的任意整数中选取一个 未加入
 * 黑名单 blacklist 的整数。任何在上述范围内且不在黑名单 blacklist 中的整数都应该有 同等的可能性 被返回。
 *
 * 优化你的算法，使它最小化调用语言 内置 随机函数的次数。
 *
 * 实现 Solution 类:
 *
 *
 * Solution(int n, int[] blacklist) 初始化整数 n 和被加入黑名单 blacklist 的整数
 * int pick() 返回一个范围为 [0, n - 1] 且不在黑名单 blacklist 中的随机整数
 *
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入
 * ["Solution", "pick", "pick", "pick", "pick", "pick", "pick", "pick"]
 * [[7, [2, 3, 5]], [], [], [], [], [], [], []]
 * 输出
 * [null, 0, 4, 1, 6, 1, 0, 4]
 *
 * 解释
 * Solution solution = new Solution(7, [2, 3, 5]);
 * solution.pick(); // 返回0，任何[0,1,4,6]的整数都可以。注意，对于每一个pick的调用，
 * ⁠                // 0、1、4和6的返回概率必须相等(即概率为1/4)。
 * solution.pick(); // 返回 4
 * solution.pick(); // 返回 1
 * solution.pick(); // 返回 6
 * solution.pick(); // 返回 1
 * solution.pick(); // 返回 0
 * solution.pick(); // 返回 4
 *
 *
 *
 *
 * 提示:
 *
 *
 * 1 <= n <= 10^9
 * 0 <= blacklist.length <= min(10^5, n - 1)
 * 0 <= blacklist[i] < n
 * blacklist 中所有值都 不同
 * pick 最多被调用 2 * 10^4 次
 *
 *
*/
package lc

import "math/rand"

// @lc code=start
// type Solution struct {
type Solution710 struct {
    bound   int
    mapping map[int]int
}

// func Constructor(n int, blacklist []int) Solution {
func Constructor710(n int, blacklist []int) Solution710 {
    bound := n - len(blacklist)
    mapping := make(map[int]int)
    // map记录下blacklist的数字，方便查找
    for _, b := range blacklist {
        mapping[b] = -1
    }

    // 只用将bound之前的黑名单数字，与bound之后的白名单数字做一个映射
    i := bound
    for _, b := range blacklist {
        if b >= bound {
            continue
        }
        for mapping[i] == -1 {
            i++
        }
        mapping[b] = i
        i++
    }
    // return Solution{bound, mapping}
    return Solution710{bound, mapping}
}

// func (s *Solution) Pick() int {
func (s *Solution710) Pick710() int {
    index := rand.Intn(s.bound)
    if _, ok := s.mapping[index]; ok {
        return s.mapping[index]
    }
    return index
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(n, blacklist);
 * param_1 := obj.Pick();
 */
// @lc code=end
