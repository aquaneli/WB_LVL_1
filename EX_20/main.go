package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "snow dog sun"
	ExampleOne(str)
	ExampleTwo(str)
}

/* Разделили строку на один слайс строк где разделитель символ ' ' */
func ExampleOne(str string) {
	strSplit := strings.Split(str, " ")
	l := len(strSplit)
	tmpStr := make([]string, l)

	/* Помещам значения первого слайса строк в новый слайс строк но с конца */
	for i, val := range strSplit {
		tmpStr[l-1-i] = val
	}

	/* Объеденяем слайс строк в строку где разделитель это пробел */
	fmt.Println(strings.Join(tmpStr, " "))
}

/* В этом варианте сначала перевернул строку а потом перевернул каждое слово */
func ExampleTwo(str string) {
	l := len(str)
	tmpStr := make([]byte, l)

	for i := 0; i < l; i++ {
		tmpStr[i] = str[l-i-1]
	}

	k := 0
	for i := 0; i < l; i++ {
		if tmpStr[i] == ' ' || i == l-1 {
			lnew := len(tmpStr[k : i+1])
			calc := 1
			if tmpStr[i] != ' ' {
				calc = 0
			}
			for j := 0; j < lnew/2; j++ {
				tmpStr[k+j], tmpStr[i-j-calc] = tmpStr[i-j-calc], tmpStr[k+j]
			}
			k = i + 1
		}
	}

	fmt.Println(string(tmpStr))
}
