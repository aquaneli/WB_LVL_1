package main

import (
	"fmt"
	"sync"
	// "time"
)

/* Вычисление квадратов чисел из массива используя конкурентность */
func main() {
	var arr = [...]int{2, 4, 6, 8, 10}
	ExampleOne(arr)
	ExampleTwo(arr)
	ExampleThree(arr)

}

/* 1. Простой вариант вывода квадрата чисел */
func ExampleOne(arr [5]int) {
	wgOne := sync.WaitGroup{}
	wgOne.Add(5)
	for i := range arr {
		go func(val int) {
			fmt.Println(val * val)
			wgOne.Done()
		}(arr[i])
	}
	wgOne.Wait()
}

/* 2. Вывода квадрата чисел с использованием небуфиризированного канала */
func ExampleTwo(arr [5]int) {
	wgTwo := sync.WaitGroup{}
	wgTwo.Add(5)
	ch := make(chan int)
	defer close(ch)

	for _, val := range arr {
		go func(ch *chan int) {
			fmt.Println(<-*ch)
			wgTwo.Done()
		}(&ch)
		ch <- val * val
	}
	wgTwo.Wait()
}

/* 3. Вывода квадрата чисел с использованием небуфиризированного канала */
func ExampleThree(arr [5]int) {
	ch := make(chan int)
	defer close(ch)

	/* Кладем значение в канал первая горутина ставится на стоп пока значение
	не считается из канала и все отсльные горутины на стопе пока канал не освободится */
	for _, val := range arr {
		go func(ch *chan int, val int) {
			*ch <- val * val
		}(&ch, val)
	}

	for range [5]int{} {
		fmt.Println(<-ch)
	}

}
