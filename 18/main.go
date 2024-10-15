/*
Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде.
По завершению программа должна выводить итоговое значение счетчика.
*/

package main

import (
	"fmt"
	"sync"
)

// Структура счетчика
type Counter struct {
	mu     sync.Mutex
	value  int
}

// Метод инкрементирования счетчика
func (c *Counter) Increment() {
	c.mu.Lock()
	c.value++
	c.mu.Unlock()
}


func main() {
	var wg sync.WaitGroup
	count := Counter{}

	// Запускаем 100 горутин, каждая инкрементирует счетчик
	for i:=0; i < 100; i++ {
		wg.Add(1)
		go func ()  {
			defer wg.Done()
			count.Increment()
		}()
	}
	wg.Wait()

	fmt.Println("Ожидаемое значение счетчика: 100\nИтоговое значение счетчика: ", count.value)
}