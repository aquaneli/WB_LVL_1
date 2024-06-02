package main

import (
	"fmt"
	"sync"
)

/* Вычисление квадратов чисел из массива используя конкурентность */

func main() {
	var arr = [...]int{2, 4, 6, 8, 10}
	wg := sync.WaitGroup{}
	wg.Add(len(arr))

	/* в цикле создаем горутины для конкурентных вычислений */
	for i := range arr {
		go func(val int) {
			fmt.Println(val * val)
			wg.Done()
		}(arr[i])
	}
	wg.Wait()
}
