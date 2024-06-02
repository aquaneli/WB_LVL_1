package main

import (
	"fmt"
	"sync"
)

/*  Найти сумму квадратов последовательности чисел */

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
