/**
 * @Author: lzw5399
 * @Date: 2021/6/4 19:59
 * @Desc: 堆排序 https://blog.csdn.net/mofiu/article/details/83620743
 * @Desc: ----------------------------------------------------------
 * @Desc: 升序排序思路
 * @Desc: 1. 先把传入的slice构建成大顶堆
 * @Desc: 2. 然后从最后一个[分支root节点]，即: (len(arr)-1)/2
 * @Desc: 3. 判断是否存在[左右孩子]，如果存在则判断[左右孩子]是否比[根节点]的值大，大就替换。最大堆构建完成
 * @Desc: 4. 然后开始把堆顶跟堆最后一个替换，并逐个缩小堆的范围，完成排序
 */
package main

import "fmt"

func main() {
	nums := []int{2, 8, 5, 4, 9, 11}
	HeapSort2(nums)

	fmt.Println("排序完之后", nums)
}

func HeapSort2(nums []int) {
	// 1. 从最后一个root节点【(len(nums)-1)/2】构建最大堆
	endIndex := len(nums) - 1
	for i := endIndex / 2; i >= 0; i-- {
		buildMaxHeap(i, endIndex, nums)
	}

	fmt.Println("最大堆构建完毕", nums)

	// 2. 然后根据堆特性排序
	for i := endIndex; i >= 0; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		buildMaxHeap(0, i-1, nums) // 重新构建最大堆，并且i减少1
	}
}

func buildMaxHeap(currentRootIndex int, endIndex int, nums []int) {
	leftChildIndex, rightChildIndex, largestIndex := currentRootIndex*2+1, currentRootIndex*2+2, currentRootIndex

	// 判断有左节点 并且 左大于root节点
	if leftChildIndex <= endIndex && nums[leftChildIndex] > nums[largestIndex] {
		largestIndex = leftChildIndex
	}

	// 判断是否有右节点 并且 右大于root节点
	if rightChildIndex <= endIndex && nums[rightChildIndex] > nums[largestIndex] {
		largestIndex = rightChildIndex
	}

	// 不等于说明需要置换
	if largestIndex != currentRootIndex {
		nums[largestIndex], nums[currentRootIndex] = nums[currentRootIndex], nums[largestIndex]
		buildMaxHeap(largestIndex, endIndex, nums)
	}
}
