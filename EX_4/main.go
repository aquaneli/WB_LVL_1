package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
)

func main() {
	ctx, cancel := context.WithCancel()
	workers := 0
	fmt.Scan(&workers)
	job := make(chan int, workers)

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)

	defer close(job)

	for i := 0; i < workers; i++ {
		go func(job chan int) {
		loop:
			for {
				select {
				case <-sigChan:
					close(sigChan)
					break loop
				default:
					fmt.Println(<-job)
				}

			}
		}(job)
	}

	for {
		job <- 1
	}

}
