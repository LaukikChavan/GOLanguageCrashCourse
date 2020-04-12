package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
}

type Circle struct {
	redius float64
}

type Rectangle struct {
	width, height float64
}

func (c Circle) area() float64 {
	return math.Pi * c.redius * c.redius
}

func (r Rectangle) area() float64 {
	return r.width * r.height
}

func getArea(s Shape) float64 {
	return s.area()
}

func main() {
	circle := Circle{redius: 20}
	rect := Rectangle{width: 10, height: 20}

	fmt.Printf("Area of Circle : %f\n", getArea(circle))
	fmt.Printf("Area of Rectangle : %f\n", getArea(rect))
}
