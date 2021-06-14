/**
 * @Author: lzw5399
 * @Date: 2021/6/9 13:52
 * @Desc: 216. 组合总和 III https://leetcode-cn.com/problems/combination-sum-iii/
 * @Desc: 找出所有相加之和为 n 的 k 个数的组合。组合中只允许含有 1 - 100 的正整数，并且每种组合中不存在重复的数字
 */
package main

import "fmt"

func main() {
	result := combinationSum(2, 18)
	fmt.Println(result)
	fmt.Println(len(result))
}

func combinationSum(k int, n int) [][]int {
	var result [][]int

	if k <= 0 || n <= k {
		return result
	}

	path := make([]int, 0, k)

	var dfs func(start int, rest int)
	dfs = func(start int, rest int) {
		if rest < 0 {
			return
		}

		if len(path) == k && rest == 0 {
			temp := make([]int, k)
			copy(temp, path)
			result = append(result, temp)
			return
		}

		for i := start; i <= 9-(k-len(path))+1; i++ {
			path = append(path, i)
			dfs(i+1, rest-i)
			path = path[:len(path)-1]
		}
	}

	dfs(1, n)
	return result
}
