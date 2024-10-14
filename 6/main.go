/*
Реализовать все возможные способы остановки выполнения горутины.
*/

// Мои варианты:
// По таймеру
// Закрытием канала

package main

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	// Остановка с использованием sync.WaitGroup
	stopByWaitGroup()

	// Остановка горутины по таймеру
	stopByTimer()

	// Остановка горутины закрытием канала
	stopByClosingChannel()

	// Остановка с использованием sync/atomic
	stopByAtomic()
	
	// Остановка с использованием context.Context
	stopByContext()
}

// Остановка с использованием sync.WaitGroup
func stopByWaitGroup() {
	fmt.Println("Остановка с использованием sync.WaitGroup:")
	var wg sync.WaitGroup

	wg.Add(1) // добавляем 1 горутину в ожидание
	go func() {
		defer wg.Done() // вызываем wg.Add(-1) -> завершаем ожидание завершения одной горутины
		for i := 0; i < 3; i++ {
			fmt.Println("Работаем...")
			time.Sleep(1 * time.Second)
		}
		fmt.Println("Горутина завершена WaitGroup")
	}()

	wg.Wait() // ожидаем завершения всех горутин
}

// Остановка горутины по таймеру
func stopByTimer() {
	fmt.Println("\nОстановка по таймеру:")
	stop := time.After(3 * time.Second)	// канал, в кот. ч/з 3 секунды будет записано значение
	
	go func() {
		for {
			select {
			case <-stop:	// если можно считать из канала, значит туда попало значение ->> время вышло
				fmt.Println("Горутина остановлена по таймеру")
				return
			default:
				fmt.Println("Работаем...")
				time.Sleep(1 * time.Second)	// задержка для красоты вывода
			}
		}
	}()
	time.Sleep(4 * time.Second)
}

// Остановка горутины закрытием канала
func stopByClosingChannel() {
	fmt.Println("\nОстановка закрытием канала:")
	ch := make(chan int)

	go func() {
		for {
			select {
			case <-ch:	// здесь мы прочитаем нулевое значение типа канала, если он закрыт и пуст
				fmt.Println("Горутина остановлена закрытием канала")
				return
			default:
				fmt.Println("Работаем...")
				time.Sleep(1 * time.Second)
			}
		}
	}()

	time.Sleep(3 * time.Second)
	close(ch) // закрываем канал для остановки горутины
	time.Sleep(1 * time.Second)
}

// Остановка с использованием sync/atomic
func stopByAtomic() {
	fmt.Println("\nОстановка с использованием sync/atomic:")
	var signal int32 = 0	// сигнал остановки

	go func() {
		for {
			if atomic.LoadInt32(&signal) == 1 {	// Установлен ли сигнал остановки
				fmt.Println("Горутина остановлена atomic")
				return
			}
			fmt.Println("Работаем...")
			time.Sleep(1 * time.Second)
		}
	}()

	time.Sleep(3 * time.Second)
	atomic.StoreInt32(&signal, 1) // Меняем флаг на 1 для остановки
	time.Sleep(1 * time.Second)
}

// Остановка с использованием context.Context
func stopByContext() {
	fmt.Println("\nОстановка с использованием context.Context:")
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)	// возвращает контекст и функ. отмены, при этом сам по таймеру отменится

	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():	//если был вызван cancel
				fmt.Println("Горутина остановлена context")
				return
			default:
				fmt.Println("Работаем...")
				time.Sleep(1 * time.Second)
			}
		}
	}(ctx)

	time.Sleep(4 * time.Second)
	cancel() // вызываем cancel для остановки горутины
}

