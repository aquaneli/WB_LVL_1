package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

/*
	Ситуация. У меня есть срез чисел и мне требуется перевернуть этот срез, но проблема в том что я могу воспользоваться только функцией Reverse() ,
	а в качестве аргумента она принимает только строки.
 	Для решения этой проблемы я создал адаптер ,который мы проинициализируем срезом и внутренними методами конвертируем срез чисел в строку и затем
 	перевернем строку . Этим же адаптером конвертируем обратно в срез чисел.
*/

type Adapter struct {
	data []int
}

func main() {
	num := []int{1, 2, 3, 4, 5, 123142141411241}
	a := Adapter{num}
	fmt.Println(num, a.StringConvertToNum(Reverse(a.NumConvertToString())))
}

/* Срез конвертировали в срез стро и затем конвертировали в одну строку */
func (d Adapter) NumConvertToString() string {
	res := []string{}
	for _, val := range d.data {
		res = append(res, strconv.Itoa(val))
	}
	return strings.Join(res, " ")
}

/* Перевернули строку */
func Reverse(str string) string {
	runeStr := []rune(str)
	l := len(runeStr)
	r := make([]rune, l)

	for i := 0; i < l; i++ {
		r[i] = runeStr[l-i-1]
	}
	return string(r)
}

/* Конвертировали строку обратно в срез чисел */
func (Adapter) StringConvertToNum(str string) []int {
	strSplit := strings.Split(str, " ")
	res := make([]int, len(strSplit))
	var err error
	for i, val := range strSplit {
		res[i], err = strconv.Atoi(val)
		if err != nil {
			log.Fatalln(err)
		}
	}
	return res
}
