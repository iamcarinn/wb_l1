/*
Дана последовательность чисел: 2,4,6,8,10.
Найти сумму их квадратов(22+32+42….)
с использованием конкурентных вычислений.
*/

package main

import (
	"fmt"
	"sync"
)

func main() {
	nums := []int{2, 4, 6, 8, 10}
	for _, num := range nums {
		fmt.Printf("%d ", num)
	}
	fmt.Printf("\n")
	result := sumSquares(nums)
	fmt.Printf("Sum of squares: %d\n", result)
}

func sumSquares(nums []int) int {
	var wg sync.WaitGroup
	ch := make(chan int, len(nums))

	// Горутины для вычисления квадратов
	for _, num := range nums {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			ch <- n * n
		}(num)
	}
	wg.Wait()
	close(ch)

	// Суммируем результаты из закрытого канала
	sum := 0
	for v := range ch {
		sum += v
	}

	return sum
}
