/**
 * @Author: lzw5399
 * @Date: 2021/6/3 20:59
 * @Desc: 求无序数组中第k大的数 https://leetcode-cn.com/problems/kth-largest-element-in-an-array/
 * @Desc: 思路: 堆排序
 */
package main

import "fmt"

func main() {
	nums := []int{1, 22, 53, 15, 6, 11, 13}
	r := findKthLargest(nums, 3)
	fmt.Println(r)
}

func findKthLargest(nums []int, k int) int {
}