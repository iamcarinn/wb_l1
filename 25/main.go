/*
Реализовать собственную функцию sleep.
*/

package main

import (
	"fmt"
	"time"
)

// Задержка на milliseconds
func sleep(milliseconds int) {
	start := time.Now() // текущее время
	// Цикл, пока не пройдет заданное время
	for time.Since(start).Milliseconds() < int64(milliseconds) {
		// ждем :)
	}
}

func main() {
	fmt.Println("Начинаем ожидание...")
	sleep(3000) // задержка на 3000 мс (3сек)
	fmt.Println("Ожидание завершено!")
}
