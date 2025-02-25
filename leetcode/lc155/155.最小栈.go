/*
* @lc app=leetcode.cn id=155 lang=golang
* @lcpr version=20004
*
* [155] 最小栈
*
* https://leetcode.cn/problems/min-stack/description/
*
  - algorithms
  - Medium (60.67%)
  - Likes:    1870
  - Dislikes: 0
  - Total Accepted:    680.4K
  - Total Submissions: 1.1M
  - Testcase Example:  '["MinStack","push","push","push","getMin","pop","top","getMin"]\n' +
    '[[],[-2],[0],[-3],[],[],[],[]]'

*
* 设计一个支持 push ，pop ，top 操作，并能在常数时间内检索到最小元素的栈。
*
* 实现 MinStack 类:
*
*
* MinStack() 初始化堆栈对象。
* void push(int val) 将元素val推入堆栈。
* void pop() 删除堆栈顶部的元素。
* int top() 获取堆栈顶部的元素。
* int getMin() 获取堆栈中的最小元素。
*
*
*
*
* 示例 1:
*
* 输入：
* ["MinStack","push","push","push","getMin","pop","top","getMin"]
* [[],[-2],[0],[-3],[],[],[],[]]
*
* 输出：
* [null,null,null,null,-3,null,0,-2]
*
* 解释：
* MinStack minStack = new MinStack();
* minStack.push(-2);
* minStack.push(0);
* minStack.push(-3);
* minStack.getMin();   --> 返回 -3.
* minStack.pop();
* minStack.top();      --> 返回 0.
* minStack.getMin();   --> 返回 -2.
*
*
*
*
* 提示：
*
*
* -2^31 <= val <= 2^31 - 1
* pop、top 和 getMin 操作总是在 非空栈 上调用
* push, pop, top, and getMin最多被调用 3 * 10^4 次
*
*
*/
package lc155

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
type MinStack struct {
	data   []int
	minVal []int
}

func Constructor() MinStack {
	return MinStack{}
}

func (ms *MinStack) Push(val int) {
	ms.data = append(ms.data, val)
	if len(ms.minVal) == 0 || val < ms.minVal[len(ms.minVal)-1] {
		ms.minVal = append(ms.minVal, val)
	} else {
		ms.minVal = append(ms.minVal, ms.minVal[len(ms.minVal)-1])
	}
}

func (ms *MinStack) Pop() {
	ms.data = ms.data[:len(ms.data)-1]
	ms.minVal = ms.minVal[:len(ms.minVal)-1]
}

func (ms *MinStack) Top() int {
	return ms.data[len(ms.data)-1]
}

func (ms *MinStack) GetMin() int {
	return ms.minVal[len(ms.minVal)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
// @lc code=end

/*
// @lcpr case=start
// ["MinStack","push","push","push","getMin","pop","top","getMin"][[],[-2],[0],[-3],[],[],[],[]]\n
// @lcpr case=end

*/
