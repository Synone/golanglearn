package main

import (
	"fmt"
	"math"
)

type triangle struct {
	base float64
	height float64
}
type square struct {
	sideLength float64
}
type shape interface {
	getArea() float64
}

func main() {
	triangle01 := triangle{
		base: 4.1,
		height: 10.5,
	}
	square01 := square{
		sideLength: 4.2,
	}
	// printArea(triangle01)
	fmt.Println(triangle01)
	printArea(triangle01)
	printArea(square01)
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}

func (t triangle) getArea() float64 {
	return 0.5*t.base*t.height
}
func (s square) getArea() float64 {
	return math.Sqrt(s.sideLength)
}