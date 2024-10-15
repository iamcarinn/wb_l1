/*
Реализовать бинарный поиск встроенными методами языка.
*/

package main

import (
	"fmt"
	"sort"
)

// Бинарный поиск вручную
func binarySearchManual(arr []int, num int) (bool, int, int) {
	low := 0             // начало массива
	high := len(arr) - 1 // конец массива
	steps := 0           // счётчик шагов

	// Пока есть элементы для поиска
	for low <= high {
		steps++
		mid := (low + high) / 2 // находим середину массива
		if arr[mid] == num {
			return true, mid, steps // возвращаем true, индекс и коли-во шагов
		} else if arr[mid] < num {
			low = mid + 1 // теперь ищем в правой половине
		} else {
			high = mid - 1 // теперь ищем в левой половине
		}
	}

	return false, -1, steps // если число не найдено, возвращаем false, -1 и кол-во шагов
}

// Бинарный поиск с исп-ем sort.Search
func binarySearchBuiltIn(arr []int, num int) (bool, int, int) {
	steps := 0

	index := sort.Search(len(arr), func(i int) bool {
		steps++
		return arr[i] >= num
	})

	// Проверяем, что найденный индекс содержит искомое число
	if index < len(arr) && arr[index] == num {
		return true, index, steps
	}

	return false, -1, steps
}

func main() {
	// Массив, в котором ищем (он должен быть отсортирован)
	arr := []int{2, 3, 5, 7, 8, 10, 17, 25, 34}
	number := 17

	// Вручную реализованный бинарный поиск
	if found, index, steps := binarySearchManual(arr, number); found {
		fmt.Printf("Ручной метод: Число %d найдено на индексе %d за %d шагов\n", number, index, steps)
	} else {
		fmt.Printf("Ручной метод: Число %d не найдено. Количество шагов: %d\n", number, steps)
	}

	// Бинарный поиск с использованием встроенной функции sort.Search
	if found, index, steps := binarySearchBuiltIn(arr, number); found {
		fmt.Printf("Встроенный метод: Число %d найдено на индексе %d за %d шагов\n", number, index, steps)
	} else {
		fmt.Printf("Встроенный метод: Число %d не найдено. Количество шагов: %d\n", number, steps)
	}
}

