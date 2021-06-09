package main

import (
	"fmt"
)

func main() {
	result := combine(4, 2)
	fmt.Println(result)
}

func combine(n int, k int) [][]int {
	var result [][]int
	path := make([]int, 0, k)

	// 回溯
	var dfs func(start int)
	dfs = func(start int) {
		if len(path) == k {
			temp := make([]int, k)
			copy(temp, path)
			result = append(result, temp)
			return
		}

		// 剪枝
		// 总共的 - (需求的 - 已经有的) +1
		for i := start; i <= n-(k-len(path))+1; i++ {
			path = append(path, i)
			dfs(i + 1)
			path = path[:len(path)-1]
		}
	}

	dfs(1)
	return result
}
