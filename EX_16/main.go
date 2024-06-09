package main

import "fmt"

/* Реализация быстрой сортировки */

func main() {
	arr := []int{1, 20, 3, 8, 2, 5, 1}
	quickSort(0, len(arr)-1, arr)
	fmt.Println(arr)
}

func quickSort(l, r int, arr []int) {
	if l >= r {
		return
	}

	s := partition(l, r, arr)
	/* Как только отсортировали главную часть массива мы сразу же вызываем
	рекурсивно функцию которая возьмет кусок от 0 элемента до элемента где пересеклись индексы массива*/
	quickSort(l, s-1, arr)
	/* Как только мы отсортировали подмассивы левой части мы сразу же начинаем сортировку рекурсивно правой части подмассивов*/
	quickSort(s, r, arr)
}

/*
Данная функция сортировка производит перестановку элементов если левый элемент больше
или равен опорному тогда мы ожидаем пока правый элемент не будет меньше или равен опорному,
чтобы произвести перестановку элементов.
*/
func partition(l, r int, arr []int) int {

	/* Находим опорный элемент */
	pivot := arr[(l+r)/2]
	/* В цикле меняем элементы местами сортируя их относитльно опорного элемента чтобы все элементы больше опорного были слева а меньше справа */
	for l <= r {

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