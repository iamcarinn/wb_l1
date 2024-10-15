/*
Разработать программу, которая проверяет, что все символы в строке уникальные (true — если уникальные, false etc). Функция проверки должна быть регистронезависимой.

Например: 
abcd — true
abCdefAaf — false
aabcd — false
*/

package main

import (
	"fmt"
	"strings"
)

// Проверка уникальности
func isUnique(input string) bool {
	input = strings.ToLower(input)	// приводим к нижнему регистру
	charMap := make(map[rune]bool)

	for _, char := range input {
		if charMap[char] {	// если уже встречался, возвращаем false
			return false
		}
		charMap[char] = true // новые добавляем в мапу
	}
	return true	// если дошли до сюда, то значит все уникальные
}

func main() {
	// Примеры строк для проверки
	tests := []string{
		"abcd",        // true
		"abCdefAaf",   // false
		"aabcd",       // false
		"unique!",     // true
		"Unique",      // true
	}

	for _, test := range tests {
		fmt.Printf("Строка: %s -  %t\n", test, isUnique(test))
	}
}
