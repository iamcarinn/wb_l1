/*
Реализовать пересечение двух неупорядоченных множеств
*/

package main

import (
	"fmt"
)

func group(one, two []int) []int {
	res := []int{}

	myMap := make(map[int]bool)

	// Добавляем в мапу все значения из первого массива
	for _, val := range one {
		myMap[val] = true
	}

	// Если в мапе встречается ключ со значением из two, то добавляем значение в результат
	for _, val := range two {
		if myMap[val] == true {
			res = append(res, val)
		}
	}

	return res
}

func main() {
	one := []int{1, 5, 52, 3, 2, 44, 14, 4}
	two := []int{7, 5, 19, 44, 1, 90, 46, 3}

	res := group(one, two)
	fmt.Println(res)
	
}
