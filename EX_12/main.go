package main

import "fmt"

type set struct {
	data map[any]any
	size int
}

/* добавить элементы в множество */
func (set *set) insert(element ...any) {
	for _, val := range element {
		if _, ok := set.data[val]; !ok {
			set.data[val] = val
			set.size++
		}
	}
}

/* полностью очистить множество */
func (set *set) clear() {
	for _, val := range set.data {
		delete(set.data, val)
		set.size--
	}
}

/* объеденить 2 множества */
func (set *set) union(other set) {
	for _, val := range other.data {
		if _, ok := set.data[val]; !ok {
			set.size++
		}
		set.data[val] = val
	}
}

/* удалить элементы в множестве */
func (set *set) remove(value ...any) {
	for _, val := range value {
		_, ok := set.data[val]
		if ok {
			delete(set.data, val)
			set.size--
		}
	}
}

/* поиск элемента в множестве */
func (set *set) find(value any) bool {
	_, ok := set.data[value]
	return ok
}

/* количество элементов в множестве */
func (set set) len() int {
	return set.size
}

/* инициализация множества */
func (set *set) create() {
	set.data = make(map[any]any)
	set.size = 0
}

/* все элементы множества */
func (set set) values() any {
	return set.data
}

func main() {
	set_example_one := set{}
	set_example_one.create()

	set_example_one.insert("cat", "cat", "dog", "cat", "tree")
	set_example_one.clear()

	fmt.Println(set_example_one.values())

	set_example_one.insert("cat", "cat", "dog", "cat", "tree")

	set_example_two := set{}
	set_example_two.create()
	set_example_two.insert("fish", "leon", "bird", "bird", "dog")

	set_example_one.union(set_example_two)
	fmt.Println(set_example_one.values())

	set_example_one.remove("fish", "other")
	fmt.Println(set_example_one.values())

	fmt.Println(set_example_one.find("leon"))

	fmt.Println(set_example_one.len())

}
