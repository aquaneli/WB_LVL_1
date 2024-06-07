package main

import "fmt"

func main() {
	arr := []int{1, 21, 3, 8, 2, 5, 1}
	quickSort(0, len(arr)-1, arr)
	fmt.Println(arr)

}

func quickSort(l, r int, arr []int) {
	if l >= r {
		return
	}

	s := partition(l, r, arr)
	quickSort(l, s-1, arr)
	quickSort(s, r, arr)
}

func partition(l, r int, arr []int) int {
	/* поработать над pivot */
	fmt.Println(l, r)
	tmp := arr[l:r]
	pivot := tmp[len(tmp)/2]

	for l < r {
		for arr[l] < pivot {
			l++
		}
		for arr[r] > pivot {
			r--
		}
		if l <= r {
			arr[l], arr[r] = arr[r], arr[l]
			l++
			r--
		}
	}
	return l
}
