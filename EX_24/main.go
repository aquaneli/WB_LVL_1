package main

import (
	"fmt"
	"math"
)

/* Структура Point с инкапсулированными параметрами x,y */
type Point struct {
	x float64
	y float64
}

/* Функция которая возвращает адрес созданного и проинициализированного объекта Point */
func New(x, y float64) *Point {
	p := Point{x: x, y: y}
	return &p
}

/* Нахождение расстояния между двумя точками по формуле AB = √(x2 - x1)^2 + √(y2 - y1)^2 */
func Distance(p1, p2 *Point) float64 {
	return math.Sqrt(math.Pow(p2.x-p1.x, 2) + math.Pow(p2.y-p1.y, 2))
}

func main() {
	p1 := New(1, 1.2)
	p2 := New(12.123, 1)
	fmt.Println(Distance(p1, p2))
}
