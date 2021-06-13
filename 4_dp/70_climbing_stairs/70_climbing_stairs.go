/**
 * @Author: lzw5399
 * @Date: 2021/6/13 11:39
 * @Desc: 70. 爬楼梯
 * @Desc: https://leetcode-cn.com/problems/climbing-stairs/
 * @Desc: 假设你正在爬楼梯。需要 n 阶你才能到达楼顶。
 * @Desc: 每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？
 */
package main

import "fmt"

func main() {
	r := climbStairs2(7)

	fmt.Println(r)
}

// 状态转移方程: f()
// f(i) = f(i-1) + f(i-2)
// 时间复杂度O(n)，空间复杂度O(n)
func climbStairs(n int) int {
	if n <= 2 {
		return n
	}
	dp := make([]int, n+1)
	dp[1] = 1
	dp[2] = 2

	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n]
}

// 状态转移方程: f()
// f(i) = f(i-1) + f(i-2)
// 时间复杂度O(n)，空间复杂度O(n)
func climbStairs2(n int) int {
	if n <= 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}

	prevTwo, prevOne, cur := 1, 2, 0
	for i := 3; i <= n; i++ {
		cur = prevTwo + prevOne
		prevTwo = prevOne
		prevOne = cur
	}

	return cur
}
