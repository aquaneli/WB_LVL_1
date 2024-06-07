package main

import "fmt"

func main() {
	a := 1
	b := 3
	ExampleOne(a, b)
	ExampleTwo(a, b)
	ExampleThree(a, b)
}

/* просто поменяли переменные местами используя синтаксис go */
func ExampleOne(a, b int) {
	a, b = b, a
	fmt.Println(a, b)
}

/* поменяли местами используя арифметические операции */
func ExampleTwo(a, b int) {
	a = a + b
	b = a - b
	a = a - b
	fmt.Println(a, b)
}

/* поменяли местами используя битовый исключающий оператор оператор ^, он берет значение бита в одинаковых местах и если они неравны то будет 1 */
func ExampleThree(a, b int) {
	a = a ^ b
	b = a ^ b
	a = a ^ b
	fmt.Println(a, b)
}
