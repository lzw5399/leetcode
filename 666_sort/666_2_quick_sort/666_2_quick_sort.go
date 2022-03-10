/**
 * @Author: lzw5399
 * @Date: 2021/6/4 19:59
 * @Desc: 快速排序
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
	var arr = []int{5, 2, 1, 4, 3}
	quicksort(arr)
	fmt.Println(arr)
}

func quicksort(array []int) {
	quickSortInner(array, 0, len(array)-1)
}

func quickSortInner(array []int, low, high int) {
	var pivotPos int //划分基准元素索引
	if low < high {
		pivotPos = partition(array, low, high)
		quickSortInner(array, low, pivotPos-1)
		quickSortInner(array, pivotPos+1, high)
	}
}

func partition(array []int, i int, j int) int {
	//第一次调用使用数组的第一个元素当作基准元素
	pivot := array[i]
	for i < j {
		for j > i && array[j] > pivot {
			j--
		}
		if j > i {
			array[i] = array[j]
			i++
		}
		for i < j && array[i] < pivot {
			i++
		}
		if i < j {
			array[j] = array[i]
			j--
		}
	}
	array[i] = pivot
	return i
}
