package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 5, 8, 20}
	fmt.Println(binarySearch(arr, 20))
}

func binarySearch(arr []int, num int) int {
	return partition(0, len(arr)-1, arr, num)
}

func partition(l, r int, arr []int, num int) int {
	step := (l + r) / 2

	if arr[step] == num {
		return step
	}

	if num > arr[step] && l != r {
		return partition(step+1, r, arr, num)
	} else if num < arr[step] && l != r {
		return partition(l, step-1, arr, num)
	}

	return -1
}
