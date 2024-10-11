/*
Написать программу, которая конкурентно рассчитает значение квадратов чисел
взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.
*/

package main

import (
	"fmt"
	"sync"
)

func processBase(arr []int) {
	wg := &sync.WaitGroup{}
	for _, num := range arr {
		wg.Add(1)
		// fmt.Printf("%#v\n", wg)
		go func(num int) {
			defer wg.Done()
			fmt.Println(num * num)
		}(num)
	}

	wg.Wait()
}

// Небуферизированный канал
func processWithChUnbuf(arr []int) {
	ch := make(chan int)

	for _, num := range arr {
		go func(n int) {
			result := n * n
			ch <- result
		}(num)
	}

	for i := 0; i < len(arr); i++ {
		fmt.Println(<-ch)
	}

	close(ch)
}

// Буферизированный канал
func processWithChBuf(arr []int) {
	ch := make(chan int, len(arr))

	for _, num := range arr {
		go func(n int) {
			result := n * n
			ch <- result
		}(num)
	}
	for i := 0; i < len(arr); i++ {
		fmt.Println(<-ch)
	}
	close(ch)
}

// Использование мьютекса
func processWithMutex(arr []int) {
	var mu sync.Mutex
	result := make([]int, len(arr))

	for i, num := range arr {
		mu.Lock()
		result[i] = num * num
		mu.Unlock()
	}

	for _, r := range result {
		fmt.Println(r)
	}
}

func main() {
	arr := []int{2, 4, 6, 8, 10}

	fmt.Println("func processBase():")
	processBase(arr)

	fmt.Println("func processWithChUnbuf():")
	processWithChUnbuf(arr)

	fmt.Println("func processWithChBuf():")
	processWithChBuf(arr)

	fmt.Println("func processWithMutex():")
	processWithMutex(arr)

}
