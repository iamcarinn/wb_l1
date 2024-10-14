/*
Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0.
*/

package main

import (
	"fmt"
)

// Установка i-го бита в 1
func setBitTo1(n int64, i uint) int64 {
	return n | (1 << i)	// побитовый сдвиг (сдвигаем единицу влево на i позиций) и побитовое ИЛИ
}

// Установка i-го бита в 0
func setBitTo0(n int64, i uint) int64 {
	return n &^ (1 << i)	// побитовый сдвиг (сдвигаем единицу влево на i позиций) и побитовое И с инверсией
}

func main() {
	var num int64 = 16
	var i uint = 3

	fmt.Printf("Число до изменения:\n%064b\n", num)

	num = setBitTo1(num, i)
	fmt.Printf("Число после установки i-го бита в 1:\n%064b\n", num)

	num = setBitTo0(num, i)
	fmt.Printf("Число после установки i-го бита в 0:\n%064b\n", num)

}