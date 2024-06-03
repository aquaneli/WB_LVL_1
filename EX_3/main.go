package main

import (
	"fmt"
	"sync"
)

/*  Найти сумму квадратов последовательности чисел */

func main() {
	ExampleOne()
	ExampleTwo()
	ExampleThree()
}

/* 1. Вывод суммы квадратов чисел используя mutex */
func ExampleOne() {
	wg := sync.WaitGroup{}
	var mutex sync.Mutex
	res := 0
	wg.Add(5)

	for _, val := range []int{2, 4, 6, 8, 10} {
		go func(val int) {
			/*  Блокируем доступ с помощью mutex */
			mutex.Lock()
			res += val * val
			wg.Done()
			mutex.Unlock()
		}(val)
	}
	wg.Wait()
	fmt.Println("ExampleOne -", res)
}

/* 2. Вывод суммы квадратов чисел используя небуферизированный канал */
func ExampleTwo() {
	var ch = make(chan int)
	defer close(ch)
	res := 0

	/* Отправляем данные в канал но данные отправятся не у всех горутин сразу, а как только освободится канал*/
	for _, val := range []int{2, 4, 6, 8, 10} {
		go func(val int) {
			ch <- val * val
		}(val)
	}

	for range []int{2, 4, 6, 8, 10} {
		res += <-ch
	}
	fmt.Println("ExampleTwo -", res)
}

/* 3. Вывод суммы квадратов чисел используя буферизированный канал */
func ExampleThree() {
	var ch = make(chan int, 4)
	defer close(ch)
	res := 0

	/* Отправляем данные в канал и данные отправятся сразу т.к. канал буферезированный */
	for _, val := range []int{2, 4, 6, 8, 10} {
		go func(val int) {
			ch <- val * val
		}(val)
	}

	for range [5]int{} {
		res += <-ch
	}

	fmt.Println("ExampleThree -", res)
}
