package main

import (
	"fmt"
	"sync"
)

/* Вычисление квадратов чисел из массива используя конкурентность */

func main() {
	wg := sync.WaitGroup{}
	wg.Add(5)

	res := 0
	for _, val := range []int{2, 4, 6, 8, 10} {
		go func(val int) {
			res += val * val
			wg.Done()
		}(val)
	}
	wg.Wait()
	fmt.Print(res)
}
