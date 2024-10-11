/*
Разработать программу, которая будет последовательно отправлять значения в канал,
а с другой стороны канала — читать. По истечению N секунд программа должна завершаться.
*/

package main

import (
	"fmt"
	"time"
)

func main() {
	// Создаем канал
	ch := make(chan int)

	N := 5 // время работы программы в секундах (по умолчанию)
	fmt.Println("Введите время работы программы в секундах: ")
	fmt.Scanf("%d\n", &N)

	// Канал, в кот. будет отправлено значение ч/з указанную задержку (5 сек)
	stop := time.After(time.Duration(N) * time.Second)

	// запускаем горутину для отправки значений в канал
	go func(ch chan int, stop <-chan time.Time) {
		send(ch, stop)
	}(ch, stop)

	// запускаем горутину для чтения значений из канала
	go func(ch chan int, stop <-chan time.Time) {
		read(ch, stop)
	}(ch, stop)

	// Дожидаемся завершения работы таймера
	time.Sleep(time.Duration(N) * time.Second)
	fmt.Println("Время вышло! Программа завершена, всем пока")
}

// Отправка в канал
func send(ch chan int, stop <-chan time.Time) {
	i := 0
	for {
		select {
		case <-stop:	// если прошло указанное время, в канал записывается значение ->> из него можно читать
			return
		case ch <- i: // отправляем данные в канал
			i++
		}
		time.Sleep(500 * time.Millisecond) // задержка м/у отправками
	}
}

// Чтение из канала
func read(ch chan int, stop <-chan time.Time) {
	for {
		select {
		case <-stop:	// если прошло указанное время, в канал записывается значение ->> из него можно читать
			return
		case val := <-ch: // читаем данные из канала
			fmt.Println("Получено:", val)
		}
	}
}
