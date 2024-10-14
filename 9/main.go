/*
Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x)
из массива, во второй — результат операции x*2,
после чего данные из второго канала должны выводиться в stdout.
*/

package main

import (
	"fmt"
	"sync"
)

func PrintChan(ch <-chan int) {
	for result := range ch {
		fmt.Printf("%d ", result)
	}
	fmt.Printf("\n")
}

func WriteToCh1(wg *sync.WaitGroup, nums []int, ch1 chan int) {
	defer wg.Done()
	for _, num := range nums {
		ch1 <- num
	}
	close(ch1)
}

func WriteToCh2(wg *sync.WaitGroup, ch1 chan int, ch2 chan int) {
	defer wg.Done()
	for num := range ch1 {
		ch2 <- num * 2
	}
	close(ch2)
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6}

	ch1 := make(chan int, len(nums))
	ch2 := make(chan int, len(nums))

	var wg sync.WaitGroup

	wg.Add(2)
	go WriteToCh1(&wg, nums, ch1)
	go WriteToCh2(&wg, ch1, ch2)
	
	wg.Wait()

	fmt.Println("Chan 2:")
	PrintChan(ch2)
}