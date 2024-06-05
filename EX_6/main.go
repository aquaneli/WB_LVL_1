package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	ch := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(7)
	ExampleOne(&wg, ch)
	ExampleTwo(&wg, ch)
	ExampleThree(&wg)
	ExampleFour(&wg)
	ExampleFive(&wg)
	ExampleSix(&wg)
	ExampleSeven(&wg)
	wg.Wait()
}

/* Горутина завершается как только в канал попадает какое-нибудь значение*/
func ExampleOne(wg *sync.WaitGroup, ch chan int) {
	go func() {
	loop:
		for {
			select {
			case <-ch:
				fmt.Println("ExampleOne stopped")
				break loop
			default:
				fmt.Println("ExampleOne works")
				time.Sleep(time.Second * 1)
			}
		}
	}()
	time.Sleep(time.Second * 3)
	ch <- 1
	wg.Done()

}

/* Горутина завершается как только канал закрывается */
func ExampleTwo(wg *sync.WaitGroup, ch chan int) {
	go func() {
	loop:
		for {
			select {
			case _, ok := <-ch:
				if !ok {
					fmt.Println("ExampleTwo stopped")
					break loop
				}
			default:
				fmt.Println("ExampleTwo works")
				time.Sleep(time.Second * 1)
			}
		}
	}()
	time.Sleep(time.Second * 3)
	close(ch)
	wg.Done()
}

/* Горутина завершается как только происходит отмена контекста */
func ExampleThree(wg *sync.WaitGroup) {
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
	loop:
		for {
			select {
			case <-ctx.Done():
				fmt.Println("ExampleThree stopped")
				break loop
			default:
				fmt.Println("ExampleThree works")
				time.Sleep(time.Second * 1)
			}
		}
	}()
	time.Sleep(time.Second * 3)
	cancel()
	wg.Done()
}

/* Горутина завершается как только Заканиваются все итерации в цикле */
func ExampleFour(wg *sync.WaitGroup) {
	go func() {
		for i := 0; i < 3; i++ {
			fmt.Println("ExampleFour works")
			time.Sleep(time.Second * 1)
		}
		fmt.Println("ExampleFour stopped")
		wg.Done()
	}()
}

/* Создается канал time.After в который подает значение через какое-то время и он автоматически закрывается */
func ExampleFive(wg *sync.WaitGroup) {
	go func() {
		res := time.After(time.Second * 3)
	loop:
		for {
			select {
			case <-res:
				fmt.Println("ExampleFive stopped")
				break loop
			default:
				fmt.Println("ExampleFive works")
				time.Sleep(time.Second * 1)
			}
		}
	}()
	wg.Done()
}

/* Завершить горутину после прочтения данных из буферизированного канала */
func ExampleSix(wg *sync.WaitGroup) {
	ch := make(chan int, 5)

	for i := 0; i < 5; i++ {
		ch <- i
	}

	go func() {
	loop:
		for {
			select {
			case <-ch:
				fmt.Println("ExampleSix works")
				time.Sleep(time.Second * 1)
			default:
				break loop
			}
		}
		fmt.Println("ExampleSix stopped")
		close(ch)
	}()
	wg.Done()
}

/* Завершил горутину используя группу ожидания */
func ExampleSeven(wg *sync.WaitGroup) {
	wgEx := sync.WaitGroup{}
	wgEx.Add(5)
	go func() {
		for {
			fmt.Println("ExampleSeven works")
			wgEx.Wait()
			break
		}
	}()
	for range [5]int{} {
		wgEx.Done()
		time.Sleep(time.Second * 1)
	}
	fmt.Println("ExampleSeven stopped")
	wg.Done()
}
