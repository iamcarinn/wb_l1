// Дана структура Human (с произвольным набором полей и методов).
// Реализовать встраивание методов в структуре Action от родительской
// структуры Human (аналог наследования).

package main

import "fmt"

type Action struct {
	h Human
}

func (a Action) DoSleep() {
	fmt.Printf("A human named %s %s is sleeping\n", a.h.name, a.h.lastname)
}

type Human struct {
	name string
	lastname string
	age int
}

func (h Human) GetName() {
	fmt.Println(h.name, h.lastname)
}

func (h Human) GetAge() {
	fmt.Printf("%d years\n", h.age)
}

func (h Human) Info() {
	h.GetName()
	h.GetAge()
}


func main() {
	action := Action{
		Human{
			name: "Bob",
			lastname: "Ivanov",
			age: 46,
		},
	}
	action.h.Info()
	action.DoSleep()
}