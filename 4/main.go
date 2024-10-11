/*
Реализовать постоянную запись данных в канал (главный поток). Реализовать
набор из `N` воркеров, которые читают произвольные данные из канала и
выводят в `stdout`. Необходима возможность выбора количества воркеров при
старте.
Программа должна завершаться по нажатию `Ctrl+C`. Выбрать и обосновать
способ завершения работы всех воркеров.
*/

package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"time"
)

func main() {
	// Запускаем постоянную запись в канал
	c := make(chan int)
	go send(c)

	// Задаем кол-во воркеров
	var workers int
	fmt.Println("Введите количество воркеров: ")
	fmt.Scanf("%d\n", &workers)

	wg := &sync.WaitGroup{}
	// Запускаем N(workers) воркеров, которые читают данные из канала
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func(ch chan int) {
			read(ch)
			wg.Done()
		}(c)
	}
	wg.Wait()

	// Создаем канал для сигнала выхода
	exitChan := make(chan os.Signal, 1)
	signal.Notify(exitChan, os.Interrupt)

	// Запускаем горутину для обработки сигнала выхода
	go func() {
		<-exitChan
		fmt.Println("\nПолучен сигнал выхода. Очистка...")
		time.Sleep(time.Second) // Даем время для закрытия каналов
		os.Exit(0)
	}()
	fmt.Println("Программа работает... [Нажмите Ctrl + C для выхода]")
	// Основной цикл программы
	for {
		time.Sleep(time.Second)
	}
}

// Запись данных в канал
func send(c chan int) {
	for {
		c <- rand.Intn(10) // случайное число записываем в канал
	}
}

// Чтение из канала
func read(c chan int) {
	select {
	case data := <-c:
		fmt.Printf("Получено число: %d\n", data)
	}
}
