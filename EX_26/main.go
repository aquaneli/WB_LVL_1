package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "qЙцуфыasа"
	fmt.Println(ExampleOne(str))
	fmt.Println(ExampleTwo(str))
}

/* В этом примере я перевел исходную строку в нижний регистр и проверяю каждый символ на совпадение в строке */
func ExampleOne(str string) bool {
	strLower := []rune(strings.ToLower(str))
	for i := range strLower {
		for j := i + 1; j < len(strLower); j++ {
			if strLower[i] == strLower[j] {
				return false
			}
		}
	}
	return true
}

/*
В этом примере создал map где в качестве значения будет ключ и если у меня не получилось взять значение
то значит элемента не существует. Запишем его в map и если в следующий раз получится вытазить его, значит было дублирование
*/
func ExampleTwo(str string) bool {
	strLower := strings.ToLower(str)
	strMap := map[rune]any{}
	for _, val := range strLower {
		_, ok := strMap[val]
		if !ok {
			strMap[val] = val
		} else {
			return false
		}
	}
	return true
}
