package main

import (
	"fmt"
	"reflect"
)

func main() {
	var val1 int
	var val2 string
	var val3 bool
	var val4 chan any
	ExampleOne(val1, val2, val3, val4)
	ExampleTwo(val1, val2, val3, val4)
	ExampleThree(val1, val2, val3, val4)
}

/* можно узнать тип используя функцию Sprintf() где флаг %T означает чтобы взяли тип*/
func ExampleOne(value ...interface{}) {
	for _, val := range value {
		str := fmt.Sprintf("%T", val)
		fmt.Println(str)
	}
}

/* с помощью switch case можно узнать тип данных */
func ExampleTwo(value ...interface{}) {
	for _, val := range value {
		switch val.(type) {
		case int:
			fmt.Println("int")
		case string:
			fmt.Println("string")
		case bool:
			fmt.Println("bool")
		case chan any:
			fmt.Println("chan interface {}")
		}

	}
}

/* используя пакет reflect можем вызвать метод TypeOf и узнать какой тип у переменной */
func ExampleThree(value ...interface{}) {
	for _, val := range value {
		xType := reflect.TypeOf(val)
		fmt.Println(xType)
	}
}
