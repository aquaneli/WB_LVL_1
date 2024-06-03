package main

import (
	"fmt"
	"sync"
)

/* Вычисление квадратов чисел из массива используя конкурентность */
func main() {
	var arr = [...]int{2, 4, 6, 8, 10}
	ExampleOne(arr)
	ExampleTwo(arr)
	ExampleThree(arr)
	ExampleFour(arr)
	ExampleFive(arr)
}

 /* 1. Простой вариант вывода квадрата чисел и группы ожидания */
func ExampleOne(arr [5]int) {
	/* Создали группу ожидания */
	wg := sync.WaitGroup{}
	/* Добавили количество горутин которые нужно обработать */
	wg.Add(5)
	for i := range arr {
		go func(val int) {
			fmt.Printf("%d ", val*val)
			wg.Done()
		}(arr[i])
	}
	wg.Wait()
	fmt.Println("- ExampleOne")
}

/* 2. Вывода квадрата чисел с использованием небуфиризированного канала и группы ожидания */
func ExampleTwo(arr [5]int) {
	/* Создали группу ожидания */
	wg := sync.WaitGroup{}
	/* Добавили количество горутин которые нужно обработать */
	wg.Add(5)
	/* Объявили и проинициализировали небуфиризированный канал */
	ch := make(chan int)
	defer close(ch)

	for _, val := range arr {
		/* Т.к. канал небуфиризированный , запущенные горутины будут заблокированы пока в канал не попадет значение. */
		go func(ch *chan int, wg *sync.WaitGroup) {
			fmt.Printf("%d ", <-*ch)
			(*wg).Done()
		}(&ch, &wg)

		/* Кладем значение в канал, но мы так же не сможем положить что либо пока горутина не возьмет значение */
		ch <- val * val
	}
	wg.Wait()
	fmt.Println("- ExampleTwo")
}

/* 3. Вывода квадрата чисел с использованием небуфиризированного канала  */
func ExampleThree(arr [5]int) {
	ch := make(chan int)
	defer close(ch)

	/* Кладем значение в канал первая горутина ставится на стоп пока значение
	не считается из канала и все отсльные горутины на стопе пока канал не освободится */
	for _, val := range arr {
		go func(ch *chan int, val int) {
			*ch <- val * val
		}(&ch, val)
	}

	/* Считываем данные из канала */
	for range [5]int{} {
		fmt.Printf("%d ", <-ch)
	}
	fmt.Println("- ExampleThree")
}

/* 4.  Вывода квадрата чисел с использованием буфиризированного канала и группы ожидания */
func ExampleFour(arr [5]int) {
	ch := make(chan int, 4)
	defer close(ch)
	wg := sync.WaitGroup{}
	wg.Add(1)

	/* Считываем данные из канала, но пока в буферезированный канал не попадет хотябы одно значение горутина будет стоять на стопе*/
	go func(ch *chan int) {
		for range [5]int{} {
			fmt.Printf("%d ", <-*ch)
		}
		wg.Done()
	}(&ch)

	/* Забрасываем данные в канал, горутины не будут заблокировани каналом т.к. мы можем закидывать данные в канал пока буфер не будет заполнен */
	for _, val := range arr {
		go func(val int) {
			ch <- val * val
		}(val)
	}
	wg.Wait()
	fmt.Println("- ExampleFour")
}

/* 5.  Вывода квадрата чисел с использованием Mutex */
func ExampleFive(arr [5]int) {
	wg := sync.WaitGroup{}
	/* Создали mutex для того чтобы она блокировала доступ к переменной i и мы могли менять значение только к одной
	переменной и не было одновременного доступа к переменной в массиве т.е. состояния гонки */
	var mutex sync.Mutex
	i := 0
	wg.Add(5)

	for range [5]int{} {
		go func(arr *[5]int, i *int, mutex *sync.Mutex) {
			/* Тут заблокировали доступ к этим инструкциям и данным пока они не будут исполнены и переменная не изменит свое значение */
			(*mutex).Lock()
			fmt.Printf("%d ", (*arr)[*i]*(*arr)[*i])
			wg.Done()
			*i++
			(*mutex).Unlock()
		}(&arr, &i, &mutex)
	}
	wg.Wait()
	fmt.Println("- ExampleFive")
}
