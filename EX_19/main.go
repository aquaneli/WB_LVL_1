package main

import "fmt"

func main() {
	str := string("главрыба")
	/* Перевел строку в срез рун , чтобы каждый символ unicod можно было проитерировать , если бы этого не сделали то символы
	были бы представлены в виде []byte где символ unicode может быть занимать от 1 до 4 байт  */
	runeStr := []rune(str)
	l := len(runeStr)
	r := make([]rune, l)

	/* В цикле прохожусь по срезу rune с конца до начала и сохраняю их в новый срез rune */
	for i := 0; i < l; i++ {
		r[i] = runeStr[l-i-1]
	}
	fmt.Println(string(r))
}
