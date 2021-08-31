package main

import (
	"fmt"
	"image/color"
)

type Point struct {
	X, Y float64
}

func main() {
	var cp ColoredPoint
	fmt.Println(cp.Point.X)
}

type ColoredPoint struct {
	Point
	Color color.RGBA
}
