package main

import (
	"fmt"
	"sync"
)

/* Вычисление квадратов чисел из массива используя конкурентность */

func main() {
	var arr = [...]int{2, 4, 6, 8, 10}
	wg := sync.WaitGroup{}
	wg.Add(1)

	ExampleOne(arr, &wg)

	wg.Wait()
}

/* 1. Простой вариант вывода квадрата чисел */

func ExampleOne(arr [5]int, wg *sync.WaitGroup) {
	wgOne := sync.WaitGroup{}
	wgOne.Add(5)
	for i := range arr {
		go func(val int) {
			fmt.Println(val * val)
			wgOne.Done()
		}(arr[i])
	}
	wgOne.Wait()
	wg.Done()
}
