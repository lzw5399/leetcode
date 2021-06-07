/**
 * @Author: lzw5399
 * @Date: 2021/6/6 20:35
 * @Desc: 升序排列
 */
package main

import "fmt"

func main() {
	nums := []int{2, 8, 5, 3, 9, 1}
	heapSort(nums)
	fmt.Println(nums)
}

// 1. 建立最大堆
//   1.1
// 2. 将堆顶移到最后，并缩小堆的大小，并重新上浮
func heapSort(nums []int) {
	endIndex := len(nums) - 1
	for i := endIndex / 2; i >= 0; i-- {
		buildMaxHeap(i, endIndex, nums)
	}
	fmt.Println("最大堆构建完毕", nums)

	for i := endIndex; i >= 0; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		buildMaxHeap(0, i-1, nums)
	}
}

func buildMaxHeap(currentRootIndex int, endIndex int, nums []int) {
	leftChildIndex, rightChildIndex, largestIndex := currentRootIndex*2+1, currentRootIndex*2+2, currentRootIndex

	if leftChildIndex <= endIndex && nums[leftChildIndex] > nums[largestIndex] {
		largestIndex = leftChildIndex
	}

	if rightChildIndex <= endIndex && nums[rightChildIndex] > nums[largestIndex] {
		largestIndex = rightChildIndex
	}

	if largestIndex != currentRootIndex {
		nums[currentRootIndex], nums[largestIndex] = nums[largestIndex], nums[currentRootIndex]
		buildMaxHeap(largestIndex, endIndex, nums)
	}
}
