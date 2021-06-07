/**
 * @Author: lzw5399
 * @Date: 2021/6/3 20:59
 * @Desc: 求无序数组中第k大的数 https://leetcode-cn.com/problems/kth-largest-element-in-an-array/
 * @Desc: 思路: 堆排序
 */
package main

import "fmt"

func main() {
	nums := []int{1, 22, 53, 7, 6, 11, 13}
	r := findKthLargest(nums, 3)
	fmt.Println(r)

	fmt.Println(nums)
}

func findKthLargest(nums []int, k int) int {
	endIndex := len(nums) - 1
	// 从最后一个分支节点开始构建最小堆
	for i:=endIndex / 2; i>=0; i-- {
		buildMaxHeap(i, endIndex, nums)
	}
	fmt.Println("最小堆构建完成", nums)

	// 移动k次堆顶到数组末尾
	for i:= endIndex; i>=k ; i-- {
		nums[i], nums[0] = nums[0], nums[i]
		buildMaxHeap(0, i-1, nums)
	}

	// 返回倒着数第k个数
	return nums[len(nums) - k]
}

// 构建最小堆
func buildMaxHeap(currentRoot int, endIndex int, nums []int) {
	leftChild, rightChild, largest := currentRoot*2+1, currentRoot*2+2, currentRoot
	if leftChild <= endIndex && nums[leftChild] > nums[largest] {
		largest = leftChild
	}
	if rightChild <= endIndex && nums[rightChild] > nums[largest] {
		largest = rightChild
	}
	if largest != currentRoot {
		nums[largest], nums[currentRoot] = nums[currentRoot], nums[largest]
		buildMaxHeap(largest, endIndex, nums)
	}
}

