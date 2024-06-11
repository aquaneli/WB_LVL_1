package main

import (
	"time"
)

func main() {
	sleep1(time.Second * 2)
	sleep2(time.Second * 2)
}

/*
В этом примере использовал метод After который отправляет в канал значение через заданное время
получаем сообщение и блокировка потока заканчивается
*/
func sleep1(d time.Duration) {
	<-time.After(d)
}

/* В данном примере я использовал метод Befor() который сообщает, находится ли текущий момент времени больше заданного */
func sleep2(d time.Duration) {
	/* Здесь мы добавляем время к текущему времени */
	tn := time.Now().Add(d)
	for {
		if !time.Now().Before(tn) {
			return
		}
	}
}
