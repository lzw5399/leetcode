/**
 * @Author: lzw5399
 * @Date: 2021/6/13 17:00
 * @Desc: 322. 零钱兑换
 * @Desc: https://leetcode-cn.com/problems/coin-change/
 * @Desc: 给定不同面额的硬币 coins 和一个总金额 amount
 * @Desc: 编写一个函数来计算可以凑成总金额所需的最少的硬币个数。
 * @Desc: 如果没有任何一种硬币组合能组成总金额返回-1
 */
package main

import "fmt"

func main() {
	//nums := []int{1, 2, 5}
	//target := 11
	nums := []int{2}
	target := 3
	result := coinChange(nums, target)

	fmt.Println(result)
}

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	// 初始化给一个不可能存在的值，用于后续判断凑不齐的情况
	for i, _ := range dp {
		dp[i] = amount + 1
	}
	dp[0] = 0

	for i := 1; i <= amount; i++ {
		for _, coin := range coins {
			if i-coin >= 0 {
				dp[i] = min(dp[i], dp[i-coin]+1)
			}
		}
	}

	if dp[amount] == amount+1 {
		return -1
	}
	return dp[amount]
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
