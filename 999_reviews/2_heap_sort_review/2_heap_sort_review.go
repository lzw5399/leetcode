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
	endIndex := len(nums) - 1
	for i := endIndex / 2; i >= 0; i-- {
		buildMaxHeap(i, endIndex, nums)
	}
	fmt.Println("最大堆构建完成", nums)

	for i := endIndex; i >= 0; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		buildMaxHeap(0, i-1, nums)
	}
}

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
