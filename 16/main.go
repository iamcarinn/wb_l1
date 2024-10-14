/*
Реализовать быструю сортировку массива (quicksort) встроенными методами языка.
*/

package main

import (
	"fmt"
)

func quicksort(data []int) []int {

	if len(data) < 2 {
		return data // массив из 0 или 1 элемента уже отсортирован
	}

	pivot := data[0]	// опорный эл-т
	less := []int{}		// эл-ты меньше или равные опорному
	more := []int{}		// эл-ты больше опорного

	// Делим эл-ты на две группы
	for _, num := range data[1:] {	// начинаем со второго эл-та среза
		if num <= pivot {
			less = append(less, num)
		} else {
			more = append(more, num)
		}
	}

	return append(append(quicksort(less), pivot), quicksort(more)...)
}

func main() {
	data := []int{4 , 6, 23, 5, 22, 1, 56, 4}
	fmt.Println("До: ", data)
	result := quicksort(data)
	fmt.Println("После: ", result)
}