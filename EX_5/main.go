package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)

	cancel()

	go func() {
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

	func() {
	loop:
		for i := 0; ; i++ {
			select {
			case <-ctx.Done():
				break loop
			default:
				ch <- i
			}
		}
	}()

}
