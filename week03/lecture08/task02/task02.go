package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

type Square struct {
	side float64
}

type Circle struct {
	radius float64
}

func (s Square) Area() float64 {
	return s.side * s.side
}

func (c Circle) Area() float64 {
	return c.radius * c.radius * math.Pi
}

func NewSquare(side float64) *Square {
	return &Square{side: side}
}

func NewCircle(radius float64) *Circle {
	return &Circle{radius: radius}
}

type Shapes []Shape

func (shapes Shapes) LargestArea() float64 {
	var maxArea float64 = 0
	for _, shape := range shapes {
		if shape.Area() > maxArea {
			maxArea = shape.Area()
		}
	}
	return maxArea
}

func main() {
	shapes := Shapes{}

	for i := 0; i < 3; i++ {
		circle := NewCircle(float64(i))
		square := NewSquare(float64(i))

		shapes = append(shapes, circle)
		shapes = append(shapes, square)
	}

	fmt.Println(shapes.LargestArea())
}
