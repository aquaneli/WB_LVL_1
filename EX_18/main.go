package main

import (
	"fmt"
	"sync"
)

/* Структура - счетчик с использованием мьютекса */
type Counter struct {
	count int
	mutex sync.Mutex
}

/* Инкрементация счетчика с использованием мьютекса для того чтобы не было состояния гонки */
func (count *Counter) add() {
	count.mutex.Lock()
	count.count++
	count.mutex.Unlock()
}

func main() {
	obj := Counter{}
	/* Группа ожидания для того чтобы программа не завершилась раньше времени */
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			obj.add()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(obj.count)
}
