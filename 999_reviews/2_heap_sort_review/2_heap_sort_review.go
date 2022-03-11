package main

import (
	"fmt"
)

func main() {
	nums := []int{7, 5, 1, 8, 9, 31, 12}
	heapSort(nums)
	fmt.Println(nums)
}

func heapSort(nums []int) {
	// 构建大顶堆
	endIdx := len(nums) - 1
	for i := endIdx / 2; i >= 0; i-- {
		buildMaxHeap(i, endIdx, nums)
	}

	// 堆顶移到末尾，并重新构建
	for i := endIdx; i >= 0; i-- {
		nums[i], nums[0] = nums[0], nums[i]
		buildMaxHeap(0, i-1, nums)
	}
}

func buildMaxHeap(currentRootIdx, endIdx int, nums []int) {
	leftChildIdx, rightChildIdx, largestIdx := currentRootIdx*2+1, currentRootIdx*2+2, currentRootIdx

	if leftChildIdx <= endIdx && nums[leftChildIdx] > nums[largestIdx] {
		largestIdx = leftChildIdx
	}

	if rightChildIdx <= endIdx && nums[rightChildIdx] > nums[largestIdx] {
		largestIdx = rightChildIdx
	}

	if largestIdx != currentRootIdx {
		nums[largestIdx], nums[currentRootIdx] = nums[currentRootIdx], nums[largestIdx]
		buildMaxHeap(largestIdx, endIdx, nums)
	}
}
