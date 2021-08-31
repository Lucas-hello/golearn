package main

import (
	"fmt"
	"math"
)

func main() {
	var s Path = []Point{{1, 2}, {4, 6}, {5, 7}}
	fmt.Println(s.Distance())

}

type Point struct {
	X float64
	Y float64
}

func Distance(p, q Point) float64 {
	return math.Hypot(p.X-q.X, p.Y-q.Y)
}

func (p Point) Distance(q Point) float64 {
	return math.Hypot(p.X-q.X, p.Y-q.Y)
}

type Path []Point

func (path Path) Distance() float64 {
	sum := 0.0
	for i := range path {
		if i > 0 {
			sum += path[i-1].Distance(path[i])
		}
	}
	return sum

}
func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}
