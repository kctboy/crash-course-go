package main

import (
	"fmt"
	"math"
)

type Shape interface {
	erea() float64
}

type Circle struct {
	x, y, radius float64
}

type Rectangle struct {
	width, height float64
}

func (c Circle) erea() float64 {
	return math.Pi * c.radius * c.radius
}

func (r Rectangle) erea() float64 {
	return r.width * r.height
}

func getArea(s Shape) float64 {
	return s.erea()
}

func main() {
	circle := Circle{x: 0, y: 0, radius: 5}
	rectangle := Rectangle{height: 4, width: 3}

	fmt.Printf("circle area : %f\n", getArea(circle))
	fmt.Printf("rectangle area : %f\n", getArea(rectangle))

}
