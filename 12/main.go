/*
Имеется последовательность строк - (cat, cat, dog, cat, tree)
создать для нее собственное множество.
*/

package main

import (
	"fmt"
)

func typesAnimal(an []string) []string {
	res := []string{}
	myMap := make(map[string]int)

	// Добавим животных в мапу, где ключ это тип животного, а значение это кол-во животных данного вида
	for _, val := range an {
		myMap[val] = myMap[val] + 1
	}
	fmt.Println("Пересчитаем животных: ", myMap, "\n")

	// Добавим все ключи из мапы в результирующее множество
	for animal, _ := range myMap {
		res = append(res, animal)
	}

	return res
}

func main() {
	animals := []string{"cat", "cat", "dog", "cat", "tree"}

	result := typesAnimal(animals)

	fmt.Println("Результат:", result)
}
