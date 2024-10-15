/*
Разработать программу нахождения расстояния между двумя точками,
которые представлены в виде структуры Point с инкапсулированными
параметрами x,y и конструктором.
*/

package main

import (
	"fmt"
	"math"
)

// Структура Point для представления точки в 2D пространстве
type Point struct {
	x, y float64 // инкапсулированные параметры
}

// Конструктор для создания новой точки
func NewPoint(x, y float64) Point {
	return Point{x: x, y: y}
}

// Метод для вычисления расстояния между двумя точками
func (p1 Point) Distance(p2 Point) float64 {
	return math.Sqrt(math.Pow(p2.x-p1.x, 2) + math.Pow(p2.y-p1.y, 2))
}

func main() {
	point1 := NewPoint(3, 4)
	point2 := NewPoint(7, 1)

	distance := point1.Distance(point2)	// расстояние между точками

	// Выводим результат
	fmt.Printf("Расстояние между точками: %.2f\n", distance)
}
