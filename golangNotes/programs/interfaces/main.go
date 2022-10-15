package main

import (
	"fmt"
	"math"
)

// Define interface
type Shape interface {
	area() float64
}

type Circle struct {
	radius float64
}

type Rectangle struct {
	width, height float64
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (r Rectangle) area() float64 {
	return r.width * r.height
}

func getArea(s Shape) float64 {
	return s.area()
}
func main() {
	circle := Circle{radius: 5}
	rectangle := Rectangle{width: 10, height: 5}
	fmt.Printf("Area of the circle is : %f\n", getArea(circle))
	fmt.Printf("Area of the Rectangle is : %f\n", getArea(rectangle))
}
