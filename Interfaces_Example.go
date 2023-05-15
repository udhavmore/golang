package main

import (
	"fmt"
	"math"
)

// Interface shape with area function
type shape interface {
	area() float64
}

// Cirle with radius property
type circle struct {
	radius float64
}

// Cirle with Length & Width property
type rectangle struct {
	length float64
	width  float64
}

// function area of type rectangle
func (r rectangle) area() float64 {
	return r.length * r.width
}

// function area of type circle
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func getArea(s shape) float64 {
	return s.area()
}

func main() {
	c1 := circle{2}       // Object circle
	r1 := rectangle{2, 4} // Object rectangle
	all_shapes := []shape{c1, r1}

	for _, shape := range all_shapes {
		// getting area of shapes which have a function "area()"
		fmt.Println(getArea(shape))
	}

}
