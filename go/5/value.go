package main

type Point1 struct {
	x, y float64
}

func (p Point1) Add(q Point1) Point1 { return Point1{p.x + q.x, p.y + q.y} }
func (p Point1) Sub(q Point1) Point1 { return Point1{p.x - q.x, p.y - q.y} }

type Path1 []Point1

func (path1 Path1) TranslateBy(offest Point1, add bool) {
	var op func(p, q Point1) Point1
	if add {
		op = Point1.Add
	} else {
		op = Point1.Sub
	}
	for i := range path1 {
		path1[i] = op(path1[i], offest)
	}

}
