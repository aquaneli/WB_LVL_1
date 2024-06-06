package main

import (
	"fmt"
	"math/rand"
)

func main() {
	/* Создал 2 set в которых будут записаны рандомные числа */
	set1 := make(map[int]any)
	set2 := make(map[int]any)

	RandomiserSet(set1, set2)
	ExampleOne(set1, set2)
	ExampleTwo(set1, set2)

}

/*  Закидываем в set рандомные числа */
func RandomiserSet(set1 map[int]any, set2 map[int]any) {
	for i := 0; i < 10; i++ {
		set1[i] = rand.Intn(10)
		set2[i] = rand.Intn(10)
	}
	fmt.Println(set1)
	fmt.Println(set2)
}

/* В этом примере я методом перебора нашел совпадения и если оно есть, тогда я записываю в результирующий set */
func ExampleOne(set1 map[int]any, set2 map[int]any) {
	set_res := make(map[int]any)
	for _, val1 := range set1 {
		for _, val2 := range set2 {
			if val1 == val2 {
				set_res[val1.(int)] = val1
				break
			}
		}
	}
	fmt.Println(set_res)
}

func ExampleTwo(set1 map[int]any, set2 map[int]any) {
	set_res := make(map[int]any)
	tmp := make(map[int]any)

	/* Здесь я в промежудочный set tmp помещаю значение и избавляюсь от дубликатов и в качестве ключа будет само значение */
	for _, val := range set1 {
		tmp[val.(int)] = val
	}

	/* Я смотрю если в промежуточном set существует ключ который равен значению второго set, тогда я помещаю ключ и значение в результирующий set */
	for _, val := range set2 {
		_, ok := tmp[val.(int)]
		if ok {
			set_res[val.(int)] = val.(int)
		}
	}

	fmt.Println(set_res)
}
