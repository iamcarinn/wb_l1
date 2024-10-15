/*
Разработать программу, которая переворачивает подаваемую
на ход строку (например: «главрыба — абырвалг»). Символы могут быть unicode.
*/

package main

import (
	"fmt"
)

func reverse(myString string) string{
	runes := []rune(myString)	// приведем к рунам, чтобы узнать кол-во символов
	newRunes := make([]rune, len(runes)) // cоздаем новый срез для перевернутой строки

	for i := len(runes) - 1; i >= 0; i-- {
		newRunes[len(runes)-1-i] = runes[i]
	}
	return string(newRunes)
}

func main() {
	var myString string = "главрыба"
	fmt.Println("Строка до переворота: ", myString)
	newString := reverse(myString)
	fmt.Println("Строка после переворота: ", newString )
}