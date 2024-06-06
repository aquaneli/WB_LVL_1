package main

import (
	"fmt"
	"log"
)

func main() {
	/* Число в двоичном виде */
	val := int64(0b0000000000000000000000000000000000000000000000000000000000000011)
	i := 0
	fmt.Println("Enter the bit number to replace")
	fmt.Scan(&i)

	if i < 1 || i > 64 {
		log.Fatalln("You entered a value less than 0 or greater than 63")
	}
	i--
	/* Проверили что в i бите стоит 0 или 1*/
	if (val>>i)&1 == 1 {
		val ^= 1 << i //установили в 0 , тут использовал оператор битовое исключающее ^. False если оба значения 0 или оба значения 1
	} else {
		val |= 1 << i //установили в 1 , тут использовал оператор битовое ИЛИ |. True если хотя бы одно значение 1
	}

	/* Вывод числа побитово */
	for i := 63; i >= 0; i-- {
		fmt.Print(val >> i & 1)
	}
	fmt.Print("\n", val)
}
