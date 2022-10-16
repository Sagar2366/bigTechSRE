package main

import "fmt"

type shape interface {
	getArea() float64
}

type traingle struct {
	base   float64
	height float64
}

type square struct {
	sideLength float64
}

func main() {
	t := traingle{
		base:   10,
		height: 1,
	}
	printArea(t)

	s := square{
		sideLength: 10,
	}
	printArea(s)
}

func (t traingle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func printArea(s shape) {
	fmt.Println("Area is : ", s.getArea())
}
