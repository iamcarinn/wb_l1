/*
Реализовать паттерн «адаптер» на любом примере.
*/

package main

import "fmt"

// Целевой интерфейс
type AudioPlayer interface {
	PlaySound()
}

// Адаптируемый класс
type OldAudioPlayer struct{}

func (o *OldAudioPlayer) PlayOldSound() {
	fmt.Println("Играет старый звук...")
}

// Адаптер
type AudioPlayerAdapter struct {
	oldPlayer *OldAudioPlayer
}

func (a *AudioPlayerAdapter) PlaySound() {
	a.oldPlayer.PlayOldSound()
}

// Функция, использующая целевой интерфейс
func playAudio(player AudioPlayer) {
	player.PlaySound()
}

func main() {
	oldPlayer := &OldAudioPlayer{}	// cоздаем старый аудиоплеер
	
	adapter := &AudioPlayerAdapter{oldPlayer: oldPlayer}	// cоздаем адаптер для старого аудиоплеера

	playAudio(adapter)	// исп-ем адаптер как целевой интерфейс
}
