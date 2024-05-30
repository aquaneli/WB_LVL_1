package main

import "fmt"

/* Создали структру Human*/
type Human struct {
	Age  int
	Name string
}

/* Определили метод для структуры Human */
func (Human) Speak() {
	fmt.Println("Hello")
}

/* Определили метод для структуры Human , но используя получатель (recipient) мы можем изменить значение поля структуры при вызове метода*/
func (man *Human) NewAge(yers int) {
	man.Age += yers
}

/* Здесь мы встроили поля из структуры Human, в том числе и методы */
type Action struct {
	Human
}

func main() {
	man := Action{
		Human: Human{
			Age:  1,
			Name: "Maksim",
		},
	}
	man.Speak()

	/* После вызова этого метода значение Age изменится*/
	man.NewAge(10)
	fmt.Println(man.Age, man.Name)
}
