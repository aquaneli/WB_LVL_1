package main

import (
	"fmt"
	"math/big"
)

/* Подключили пакет который работает с большими числами */
func main() {

	/* Создали указатель на новый объект в памяти и в этот объект поместили
	число в виде строки и задали значение в десятичной системе */

	a, _ := new(big.Int).SetString("123123123123123123123123123123123123123123123123123123123123123123123123", 10)
	b, _ := new(big.Int).SetString("789789789789789789789789789789789879789789789789789789789789789789789789", 10)

	/* перемножает */
	mul := new(big.Int).Mul(a, b)
	/* делит */
	del := new(big.Int).Div(mul, a)
	/* складывает */
	sum := new(big.Int).Add(del, b)
	/* вычитает */
	sub := new(big.Int).Sub(a, mul)

	fmt.Println(sum)
	fmt.Println(sub)
	fmt.Println(mul)
	fmt.Println(del)
}
