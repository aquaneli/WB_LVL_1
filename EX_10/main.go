package main

import "fmt"

func main() {
	/* Создали слайс с температурами */
	temp := []float64{
		-25.4, -27.0, 13.0, 19.0,
		15.5, 24.5, -21.0, 32.5,
	}

	/* Создали множество т.е. map со значением слайсов float64*/
	set := make(map[int][]float64)

	/* Добавляем подмножества */
	for _, val := range temp {
		i := int(val) / 10
		set[i*10] = append(set[i*10], val)
	}

	/* Вывод множества */
	fmt.Println(set)
}
