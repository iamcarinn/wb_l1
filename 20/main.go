/*
Разработать программу, которая переворачивает слова в строке. 
Пример: «snow dog sun — sun dog snow».
*/

package main

import (
	"fmt"
	"strings"
)

func reverseWords(input string) string {
	words := strings.Fields(input)            // разбиваем строку на срез слов
	reversedWords := make([]string, len(words)) // создаем новый срез для перевернутых слов

	for i := len(words) - 1; i >= 0; i-- {
		reversedWords[len(words)-1-i] = words[i] // переворачиваем порядок слов
	}

	return strings.Join(reversedWords, " ") // соединяем слова обратно в строку
}

func main() {
	myString := "snow dog sun"
	fmt.Println("Строка до переворота: ", myString)
	newString := reverseWords(myString)
	fmt.Println("Строка после переворота: ", newString)
}
