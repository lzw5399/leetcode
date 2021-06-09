/**
 * @Author: lzw5399
 * @Date: 2021/6/7 21:42
 * @Desc: 77. 组合 https://leetcode-cn.com/problems/combinations/
 * @Desc: 给定两个整数 n 和 k，返回 1 ... n 中所有可能的 k 个数的组合。
 */
package main

import "fmt"

func main() {
	ans := combine(4, 2)
	fmt.Println(ans)
}

func combine(n int, k int) [][]int {
	var ans [][]int

	if k < 0 || n < k {
		return ans
	}

	path := make([]int, 0, k)

	var dfs func(start int)
	dfs = func(start int) {
		if len(path) == k {
			temp := make([]int, k)
			copy(temp, path)
			ans = append(ans, temp)
			return
		}

		// start <= 总共的(n) - 还缺多少个( 需要多少个 - 有了多少个 )
		// 5 = 7 - (3 - 0 + 1)
		// 搜索起点的上界 + 接下来要选择的元素个数 - 1 = n
		// 即: 搜索起点的上界 = n - (k -len(path)) + 1
		for i := start; i <= n-(k-len(path))+1; i++ {
			path = append(path, i)
			dfs(i + 1)
			// 移除最后一个元素
			path = path[:len(path)-1]

		}
	}

	dfs(1)
	return ans
}
