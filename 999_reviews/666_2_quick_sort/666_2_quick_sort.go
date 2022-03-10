package main

import "fmt"

func main() {
	var arr = []int{5, 2, 1, 4, 3, 0, 23, 11}
	quickSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}

func quickSort(arr []int, low int, high int) {
	if low < high {
		pivot := partition(arr, low, high)
		quickSort(arr, low, pivot-1)
		quickSort(arr, pivot+1, high)
	}
}

func partition(arr []int, low int, high int) int {
	pivot := arr[low]
	if low < high {
		// high
		for low < high && arr[high] > pivot {
			high--
		}
		arr[low] = arr[high]

		// low
		for low < high && arr[low] < pivot {
			low++
		}
		arr[high] = arr[low]
	}

	arr[low] = pivot
	return low
}
