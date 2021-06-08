/**
 * @Author: lzw5399
 * @Date: 2021/6/7 21:42
 * @Desc: 77. 组合 https://leetcode-cn.com/problems/combinations/
 * @Desc: 给定两个整数 n 和 k，返回 1 ... n 中所有可能的 k 个数的组合。
 */
package main

import "fmt"

func main() {
	ans := combineJava(10, 3)
	fmt.Println(ans)
}

func combineJava(n int, k int) (ans [][]int) {
	if k < 0 || n < k {
		return
	}
	var dfs func(start int)
	path := make([]int, 0, k)
	dfs = func(start int) {
		if len(path) == k {
			temp := make([]int, k)
			copy(temp, path)
			ans = append(ans, temp)
			return
		}

		for i := start; i <= n; i++ {
			path = append(path, i)

			fmt.Println("递归之前", path)

			dfs(i + 1)

			fmt.Println("递归之后", path)
			// 移除最后一个元素
			path = path[:len(path)-1]

		}
	}

	dfs(1)
	return
}

//func combine(n int, k int) (ans [][]int) {
//	var temp []int
//	var dfs func(int)
//	dfs = func(cur int) {
//		// 剪枝：temp 长度加上区间 [cur, n] 的长度小于 k，不可能构造出长度为 k 的 temp
//		if len(temp)+(n-cur+1) < k {
//			return
//		}
//		// 记录合法的答案
//		if len(temp) == k {
//			comb := make([]int, k)
//			copy(comb, temp)
//			ans = append(ans, comb)
//			return
//		}
//		// 考虑选择当前位置
//		temp = append(temp, cur)
//
//		dfs(cur + 1)
//		temp = temp[:len(temp)-1]
//
//		// 考虑不选择当前位置
//		dfs(cur + 1)
//	}
//	dfs(1)
//	return
//}
