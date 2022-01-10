/*Разработать программу нахождения расстояния между двумя точками, которые
представлены в виде структуры Point с инкапсулированными параметрами x,y и
конструктором.*/
package main

import (
	"fmt"
	"math"
)

type Point struct{
	x, y float64
}

func (p Point)Distance(anotherPoint *Point) float64 {
	return math.Sqrt(math.Pow(p.x - anotherPoint.x, 2) + math.Pow(p.y - anotherPoint.y, 2)) // distance
}

func NewPoint(x, y float64) *Point{
	p := new(Point) //returns pointer to element (zero value)
	p.x = x
	p.y = y
	return p // return pointer to element with parameters
}

func main(){
	a := NewPoint(0, 0)
	b := NewPoint(4, 5)
	fmt.Println(b.Distance(a))
}
