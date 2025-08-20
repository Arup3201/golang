package main

import (
	"fmt"
	"golang/interface/rectangle"
)

type Geometry interface {
	Area() float64
	Perimeter() float64
}

func measure(g Geometry) {
	fmt.Printf("Geometry area: %f\n", g.Area())
	fmt.Printf("Geomery perimeter: %f\n", g.Perimeter())
}

func main() {
	r := rectangle.NewRectangle(23, 40)
	measure(r)
}
