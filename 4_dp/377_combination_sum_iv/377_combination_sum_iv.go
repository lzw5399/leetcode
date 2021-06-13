/**
 * @Author: lzw5399
 * @Date: 2021/6/13 10:40
 * @Desc: 377. 组合总和 Ⅳ
 * @Desc: https://leetcode-cn.com/problems/combination-sum-iv/
 * @Desc: 给你一个由 不同 整数组成的数组 nums ，和一个目标整数 target 。
 * @Desc: 请你从 nums 中找出并返回总和为 target 的元素组合的个数。
 */
package main

import "fmt"

func main() {
	nums := []int{2}
	target := 3
	result := combinationSum4(nums, target)

	fmt.Println(result)
}

// nums种面额，凑齐target的总额
func combinationSum4(nums []int, target int) int {
	dp := make([]int, target+1)

	// 不选取任何元素时候，target才为0，即只有"不选取任何元素"一种方案
	dp[0] = 1

	for i := 1; i <= target; i++ {
		for _, num := range nums {
			if i-num >= 0 {
				dp[i] += dp[i-num]
			}
		}
	}

	if dp[target] == 0 {
		return -1
	}
	return dp[target]
}
