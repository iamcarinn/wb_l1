/*
Разработать программу, которая перемножает, делит, складывает,
вычитает две числовых переменных a,b, значение которых > 2^20.
*/

package main

import (
	"fmt"
)

// Вычисления
func calculate(a, b int) (int, int, int, float64) {
	sum := a + b
	difference := a - b
	product := a * b
	quotient := float64(a) / float64(b) // приводим к float64 для деления

	return sum, difference, product, quotient
}

// Валидация
func validateInput(a, b int) bool {
	return a > (1 << 20) && b > (1 << 20)
}

func main() {
	a := 1048577
	b := 1500000
	
	//  Проверка, что a и b больше 2^20
	if !validateInput(a, b) {
		fmt.Println("Оба числа должны быть больше 2^20.")
		return
	}

	sum, difference, product, quotient := calculate(a, b)

	fmt.Println("Сумма:", sum)
	fmt.Println("Разность:", difference)
	fmt.Println("Произведение:", product)
	fmt.Println("Частное:", quotient)
}
