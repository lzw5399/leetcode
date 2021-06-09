/**
 * @Author: lzw5399
 * @Date: 2021/6/9 0:22
 * @Desc:
 */
package main

import "fmt"

func main() {
	re := combination(4, 2)
	fmt.Println(re)
}

// 回溯+剪枝
func combination(n int, k int) [][]int {
	var result [][]int
	if n <= 0 || k <= 0 || k > n {
		return result
	}
	path := make([]int, 0, k)

	// 深度优先遍历进行回溯
	var dfs func(start int)
	dfs = func(start int) {
		if len(path) == k {
			temp := make([]int, k)
			copy(temp, path)
			result = append(result, temp)
			return
		}

		// 默认只是 i <= n
		// 剪枝情况下 i <= n - ()
		for i := start; i <= n-(k-len(path))+1; i++ {
			path = append(path, i)

			dfs(i + 1)

			// 移除最后一个
			path = path[:len(path)-1]
		}
	}

	dfs(1)

	return result
}
