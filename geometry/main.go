package main

import (
	"fmt"
	"math"
)

type Point struct{ X, Y float64 }

func (p *Point) Distance(q Point) float64 {
	return math.Hypot(p.X-q.X, p.Y-q.Y)
}

type Path []Point

func (path Path) Distance() float64 {
	var sum float64 = 0

	for i := range path {
		if i > 0 {
			sum += path[i-1].Distance(path[i])
		}
	}
	return sum
}

func main() {
	var a, b, c, d Point
	a = Point{X: 1, Y: 2}
	b = Point{X: 2, Y: 4}
	c = Point{X: 1, Y: 6}
	d = Point{X: 10, Y: 2}

	p := Path{a, b, c, d}
	fmt.Println(p.Distance())
}
