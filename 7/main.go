/*
Реализовать конкурентную запись данных в map.
*/

package main

import (
	"fmt"
	"sync"
)

// sync.Mutex
// sync.Map



// Мьютекс гарантирует, что одна горутина может записывать в карту одновременно


func main() {
	myMap := make(map[int]int)
	

	var mut sync.Mutex
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int)  {
			defer wg.Done()
			mut.Lock()
			myMap[i] = i*i	// благодаря мьютексу никто не изменит значение myMap параллельно
			mut.Unlock()
		} (i)
	}
	wg.Wait()

	mut.Lock()
	fmt.Println("Результат:", myMap)
	mut.Unlock()
}