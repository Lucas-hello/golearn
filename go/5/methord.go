package main

import (
	"fmt"
	//"fmt"
	"image/color"
	"math"
)

type Point struct{ x, y float64 }
type ColoredPoint struct {
	Point
	Color color.RGBA
}

func Distance(p, q Point) float64 {
	return math.Hypot(q.x-p.x, q.y-p.y)
}
func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.x-p.x, q.y-p.y)
}

//func xx () {
//	var a, b, c, d int
//	p := ColoredPoint{Point{1,1},color.RGBA{255,0,0,255}}

//	a, b, c = 1, 2, 3
//	fmt.Println(a, b, c)
//}
func main() {
	red := color.RGBA{255, 0, 0, 255}
	blue := color.RGBA{0, 0, 255, 255}
	var p = ColoredPoint{Point{1, 1}, red}
	var q = ColoredPoint{Point{5, 4}, blue} //Color这个参数
	fmt.Println(p.Distance(q.Point))
	//p:=Point{1,2}
	//q:=Point{4,6}
	////fmt.Println(Distance(p,q))
	////fmt.Println(p.Distance(q))
	//prrim:=Path{{1,1},{5,1},{2,8},{1,1}}
	//fmt.Println(Path.Distance(prrim))
	//r:=&Point{1,2}
	//r.ScaleBy(2)
	//fmt.Println(*r)
	//p.ScaleBy(2)//为什么回允许执行
	//distanceFromp:= p.Distance
	//fmt.Println(distanceFromp(q))
	//var org Point
	//fmt.Println(distanceFromp(org))
}

type Path []Point //是一个底层是Point的切片

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
	p.x *= factor
	p.y *= factor

}
