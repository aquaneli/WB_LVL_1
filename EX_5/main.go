package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	/* Создал контекст который отключится через заданное время */
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer close(ch)
	defer cancel()

	/* Горутина которая посылает данные в канал */
	go func() {
	loop:
		for i := 0; ; i++ {
			select {
			case <-ctx.Done():
				break loop
			default:
				ch <- i
				time.Sleep(time.Second * 1)
			}
		}
	}()

	/* Анонимная функция которая выводит данные из канала */
	func() {
	loop:
		for i := 0; ; i++ {
			select {
			case <-ctx.Done():
				break loop
			case res := <-ch:
				fmt.Println(res)
			}

		}
	}()

}
