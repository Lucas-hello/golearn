package main

import "fmt"

type bit float64

const (
	KB bit = 1000 * (iota + 1)
	MB     = 1000 * KB
	GB     = 1000 * MB
	TB
	PB
	EB
	ZB
	YB
)

func main() {
	//f := func(x float64) int {
	//	return int(x)
	//}
	//x0 := f(math.Pow10(0 * 3))
	//x1 := f(math.Pow10(1 * 3))
	//x2 := f(math.Pow10(2 * 3))
	//fmt.Println(x0, x1, x2)
	f := func(x bit) int {
		return int(x)
	}
	x0 := f(KB)
	x1 := f(MB)
	x2 := f(GB)
	x3 := f(TB)
	fmt.Println(x0, x1, x2, x3)
}
