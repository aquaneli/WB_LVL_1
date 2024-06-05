package main

import (
	"fmt"
	"sync"
)

func main() {
	// ExampleOne()
	ExampleTwo()
}

func ExampleOne() {
	wg := sync.WaitGroup{}
	wg.Add(10)
	mutex := sync.Mutex{}
	data := make(map[int]int)
	for i := 0; i < 10; i++ {
		go func(i int) {
			mutex.Lock()
			data[i] = i
			wg.Done()
			mutex.Unlock()
		}(i)
	}
	fmt.Println(data)
}

func ExampleTwo() {
	w := sync.Map{}
	wg := sync.WaitGroup{}
	data := make(map[int]int)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			w.Store(i, i*i)
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println(data)
}
