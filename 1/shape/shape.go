package shape

import (
	"math"
)

type Shape interface {
	ShapeWithArea() float32
	ShapeWithPerimeter() float32
}

type Square struct {
	sideLength float32
}

func NewSquare(sideLength float32) Square {
	return Square{sideLength}
}

func (s Square) ShapeWithArea() float32 {
	return s.sideLength * s.sideLength
}

func (s Square) ShapeWithPerimeter() float32 {
	return s.sideLength * 4
}

type Circle struct {
	radius float32
}

func (c Circle) ShapeWithArea() float32 {
	return c.radius * c.radius * math.Pi
}

func (c Circle) ShapeWithPerimeter() float32 {
	return 2 * c.radius * math.Pi
}
