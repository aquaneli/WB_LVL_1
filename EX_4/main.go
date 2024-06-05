package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"
)

func main() {
	/* Здесь я создал контекст чтобы все горутины смогли получить сигнал о том что я хочу завершить программу комбинацией Ctrl+C */
	ctx, cancel := context.WithCancel(context.Background())

	fmt.Println("Select the number of workers")
	workers := 0

	/* Ввожу количество воркеров*/
	/* Пул воркеров это паттерн для того чтобы я мог запустить ограниченное количество горутин для определенных задач и не расходовал лишние ресурсы и для
	того чтобы мы могли проще контролировать процесс */
	fmt.Scan(&workers)

	/* Создал буферезированный канал для того чтобы другие горутины не ждали в очереди поступающее сообщение */
	job := make(chan int, workers)
	defer close(job)

	/* Здесь я создал канал чтобы в него поступал какой-либо сигнал */
	sigChan := make(chan os.Signal, 1)

	/* Метод который настраивает канал на то что после нажатия комбинации клавиш Ctrl+C в канал попадет сообщение */
	signal.Notify(sigChan, os.Interrupt)
	defer close(sigChan)

	/* Анонимная функция в которой будет вызвана отмена контекста и затем во все горутины поступит сообщение об этом и воркеры завершатся.
	Затем также завершится горутина которая посылает сообщения в каналы */
	go func() {
		if _, ok := <-sigChan; ok {
			cancel()
		}
	}()

	/* Создание воркеров */
	for i := 0; i < workers; i++ {
		go worker(ctx, job)
	}

	jobs(ctx, job)

	fmt.Println("The program has stopped")

}

/* Воркер */
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

/* метод который посылает в воркеры сообщения на вывод */
func jobs(ctx context.Context, job chan int) {
loop:
	for i := 0; ; i++ {
		select {
		case <-ctx.Done():
			break loop
		default:
			job <- i
			time.Sleep(time.Second * 1)
		}
	}
}
