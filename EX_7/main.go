package main

import (
	"fmt"
	"sync"
)

func main() {
	ExampleOne()
	fmt.Println("-----")
	ExampleTwo()
	fmt.Println("-----")
	ExampleThree()
}

/*  В этом примере для конкурентных вычислений используется mutex и группа ожидания */
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
	wg.Wait()
	for _, val := range data {
		fmt.Println(val)
	}
}

/*  В этом примере для конкурентных вычислений используется RWMutex и группа ожидания */
func ExampleTwo() {
	wg := sync.WaitGroup{}
	rwmutex := sync.RWMutex{}
	data := make(map[int]int)
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func(i int) {
			rwmutex.Lock()
			data[i] = i
			wg.Done()
			rwmutex.Unlock()
		}(i)
	}
	wg.Wait()
	wg.Add(10)

	for i := 0; i < 10; i++ {
		rwmutex.RLock()
		fmt.Println(data[i])
		rwmutex.RUnlock()
		wg.Done()
	}

	wg.Wait()
}

/*
В этом примере для конкурентных вычислений используется sync.Map и группа ожидания
sync.Map используется в основном для того чтобы избежать cache contetion
*/
func ExampleThree() {
	m := sync.Map{}
	wg := sync.WaitGroup{}

	func() {
		for i := 0; i < 10; i++ {
			wg.Add(1)
			go func(i int) {
				m.Store(i, i)
				wg.Done()
			}(i)
		}
	}()

	wg.Wait()
	func() {
		for i := 0; i < 10; i++ {
			wg.Add(1)
			go func(i int) {
				v, _ := m.Load(i)
				fmt.Println(v)
				wg.Done()
			}(i)
		}
	}()

	wg.Wait()
}
