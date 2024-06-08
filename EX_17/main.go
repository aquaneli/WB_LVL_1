package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 5, 8, 20, 123}
	/* Бинарный поиск отсортированного массива с помощью рекурсии */
	fmt.Println(binarySearchRecurse(arr, 1))

	/* Бинарный поиск отсортированного массива с помощью цикла */
	fmt.Println(binarySearchCycle(arr, 1))
}

func binarySearchRecurse(arr []int, num int) int {
	return partition(0, len(arr)-1, arr, num)
}

/* Используя рекурсию мы делим массив на подмассив относительно середины где предположительно должен находиться нащ элемент*/
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

/* Так же находим середину массива и относительно опорного элемента делим на подмассивы и ищем в подмассивах искомы элемент */
func binarySearchCycle(arr []int, num int) int {
	l := 0
	r := len(arr) - 1

	for l <= r {
		step := (l + r) / 2
		if arr[step] == num {
			return step
		} else if num > arr[step] {
			l = step + 1
		} else if num < arr[step] {
			r = step - 1
		}
	}

	return -1
}
