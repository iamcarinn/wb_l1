/*
Разработать программу, которая в рантайме способна
определить тип переменной: int, string, bool, channel
из переменной типа interface{}.
*/

package main

import "fmt"

func findType(data interface{}) {
	switch data.(type) {
	default:
		fmt.Printf("unexpected type %T", data)
	case bool:
		fmt.Printf("boolean %t\n", data)
	case int:
		fmt.Printf("integer %d\n", data)
	case string:
		fmt.Printf("string %s\n", data)
	case chan int:
		fmt.Printf("chan int %d\n", data)
	case chan string:
		fmt.Printf("chan string %d\n", data)
	case chan bool:
		fmt.Printf("chan bool %d\n", data)
	}

}

func main() {
	var interInt interface{} = 16
	findType(interInt)

	var interBool interface{} = true
	findType(interBool)

	var interString interface{} = "Study"
	findType(interString)
}