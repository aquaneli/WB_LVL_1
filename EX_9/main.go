package main

import (
	"fmt"
)

func main() {
	arr := [5]int{1, 2, 3, 4, 5}
	ch1 := make(chan int)
	ch2 := make(chan int)
	defer close(ch1)
	defer close(ch2)

	/* Посылаем данные в канал ch1 */
	go func() {
		for _, val := range arr {
			ch1 <- val
		}
	}()

	/* Считываем данные с канала ch1 и пересылаем данные в ch2 */
	go func() {
		for val := range ch1 {
			ch2 <- val * 2
		}
	}()

	/* Считываем данные из канала ch2 и выводим */
	for i := 0; i < 5; i++ {
		fmt.Println(<-ch2)
	}

}
