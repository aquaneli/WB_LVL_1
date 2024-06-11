package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {
	sl := []any{1, 2, 3, 4, 5}

	sl, err := ExampleOne(sl, 1)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(sl)

	sl, err = ExampleTwo(sl, 3)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(sl)
}

/* Удаление элемента копирование среза в другой и возвратом среза меньшего размера */
func ExampleOne(sl []any, i int) ([]any, error) {
	if i > len(sl)-1 || len(sl) == 0 || i < 0 {
		return nil, errors.New("error")
	} else {
		copy(sl[i:], sl[i+1:])
		sl = sl[:len(sl)-1]
	}
	return sl, nil
}

/* Создание среза меньшего размера и копирование в него всех элементов индекс которых не равен i */
func ExampleTwo(sl []any, i int) ([]any, error) {
	slNew := []any{}
	if i > len(sl)-1 || len(sl) == 0 || i < 0 {
		return slNew, errors.New("error")
	} else {
		slNew = make([]any, len(sl)-1)
		k := 0
		for j, val := range sl {
			if j != i {
				slNew[k] = val
				k++
			}
		}
	}
	return slNew, nil
}
