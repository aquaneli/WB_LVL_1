package main

import (
	"fmt"
	"sync"
)

func main() {
	data := make(map[int]int)
	mutex := sync.Mutex{}
	wg := sync.WaitGroup{}
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func(i int) {
			mutex.Lock()
			data[i] = i
			wg.Done()
			mutex.Unlock()
		}(i)
	}

	wg.Wait()
	fmt.Println(data)
}
