/**
 * @Author: lzw5399
 * @Date: 2021/6/7 16:00
 * @Desc: 给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target  的那 两个 整数，并返回它们的数组下标
 * @Desc: https://leetcode-cn.com/problems/two-sum/
 * @Desc: 思路: 哈希法
 */
package main

import "fmt"

func main() {
	nums := []int{1, 6, 5, 4, 3, 2}

	result := twoSum(nums, 5)

	fmt.Println(result)
}

func twoSum(nums []int, target int) []int {
	// 值和index反转的map
	m := map[int]int{}
	for i, v := range nums {
		if p, exist := m[target-v]; exist {
			return []int{i, p}
		}
		m[v] = i
	}

	return nil
}