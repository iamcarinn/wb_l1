/*
Удалить i-ый элемент из слайса.
*/

package main

import (
	"fmt"
)

// Удаляем i-й элемент из слайса
func removeElement(slice []int, i int) []int {
	if i < 0 || i >= len(slice) {
		fmt.Println("ERROR: Некорректный индекс для данного слайса")
		return slice
	}
	// Удаляем элемент, создавая новый слайс
	return append(slice[:i], slice[i+1:]...)
}

func main() {
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println("Исходный слайс:", slice)

	i := 2
	slice = removeElement(slice, i)	// yдаляем i-й элемент из слайса
	fmt.Println("Слайс после удаления элемента:", slice)
}
