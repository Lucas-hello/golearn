package main

import (
	"fmt"
	"math"
)

type Rectangle struct {

	long,wield float64

}
type Circle struct {
	half float64
}
func (r Rectangle) area() float64 {
	return r.long*r.wield
}
func (c Circle) area()  float64 {
	return c.half * c.half * math.Pi
	
}
func main()  {
	r1 := Rectangle{12, 2}
	r2 := Rectangle{9, 4}
	c1 := Circle{10}
	c2 := Circle{25}
	fmt.Println("the first is",r1.area())
	fmt.Println("second is ",r2.area())
	fmt.Println("Area of c1 is: ", c1.area())
	fmt.Println("Area of c2 is: ", c2.area())
}