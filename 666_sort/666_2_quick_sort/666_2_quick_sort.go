/**
 * @Author: lzw5399
 * @Date: 2021/6/4 19:59
 * @Desc: 快速排序 https://blog.csdn.net/nrsc272420199/article/details/82587933
 */
package main

import "fmt"

func main() {
	var arr = []int{5, 2, 1, 4, 3, 0, 23, 11}
	quicksort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}

func quicksort(array []int, low int, high int) {
	//划分基准元素索引
	if low < high {
		pivotPos := partition(array, low, high)
		quicksort(array, low, pivotPos-1)
		quicksort(array, pivotPos+1, high)
	}
}

func partition(array []int, low int, high int) int {
	//第一次调用使用数组的第一个元素当作基准元素
	pivot := array[low]
	for low < high {
		// 当队尾元素大于等于基准元素时，向前挪动
		for high > low && array[high] > pivot {
			high--
		}
		// 队尾元素小于pivot，赋值给low
		array[low] = array[high]

		// 队首元素小于等于temp, 向后挪动
		for low < high && array[low] < pivot {
			low++
		}
		// 队首元素小于pivot，赋值给low
		array[high] = array[low]
	}

	array[low] = pivot
	return low
}
