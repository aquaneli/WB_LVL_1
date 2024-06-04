package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	workers := 0
	fmt.Println("Select the number of workers")
	fmt.Scan(&workers)
	job := make(chan int, workers)
	defer close(job)

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	defer close(sigChan)

	go func() {
		if _, ok := <-sigChan; ok {
			cancel()
		}
	}()

	for i := 0; i < workers; i++ {
		go func(ctx context.Context, job chan int) {
		loop:
			for {
				select {
				case <-ctx.Done():
					break loop
				case res := <-job:
					fmt.Println(res)
				}
			}
		}(ctx, job)
	}

loop:
	for i := 0; ; i++ {
		select {
		case <-ctx.Done():
			break loop
		default:
			job <- i
		}
	}

	fmt.Println("The program has stopped")

}

func worker(ctx context.Context, job chan int) {
loop:
	for {
		select {
		case <-ctx.Done():
			break loop
		case res := <-job:
			fmt.Println(res)
		}
	}
}
