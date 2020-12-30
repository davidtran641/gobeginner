package main

import (
	"fmt"
)

// Point point
type Point struct {
	X int
	Y int
}

// Square square
type Square struct {
	Origin Point
	Size   int
}

// Painter to pain
type Painter interface {
	Paint()
}

// Paint Square to implement Paint
func (s Square) Paint() {
	fmt.Println("Paint square")
}

// Paint Point to implement paint
func (s Point) Paint() {
	fmt.Println("Paint Point")
}

func typeCheck(painter Painter) {
	switch value := painter.(type) {
	case Square:
		fmt.Println("Square", value)
	default:
		fmt.Println("Unknown type", value)
	}
	painter.Paint()
}

func check(value interface{}) {
	num, ok := value.(int)
	if ok {
		fmt.Println("Given value is a number", num)
	} else {
		fmt.Println("Given value is not a number")
	}
}

func main() {
	point := Point{X: 0, Y: 10}
	typeCheck(point)

	square := Square{Origin: Point{X: 0, Y: 10}, Size: 2}
	typeCheck(square)

	check(point)
	check(10)
}
