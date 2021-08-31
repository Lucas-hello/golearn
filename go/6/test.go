package main

import (
	"fmt"
	"math"
)

func main() {
	var t1 Triangle
	t1 = Triangle{3, 4, 5}
	fmt.Println(t1.area())
	fmt.Println(t1.peri())
	fmt.Println(t1.a, t1.b, t1.c)
	var c1 Circle
	c1 = Circle{4}
	fmt.Println(c1.area())
	fmt.Println(c1.peri())
	fmt.Println(c1.r)
	var s1 Shape
	s1 = t1
	fmt.Println(s1.peri())
	fmt.Println(s1.area())
	//fmt.Println(s1.a)//不能访问实例中的参数
	getType(t1)

}
func getType(s Shape) {
	if instance, ok := s.(Triangle); ok { //返回值是Triangle的实例，赋值给instance
		fmt.Println("Triangle", instance.a, instance.b, instance.c)
	} else if instance, ok := s.(Circle); ok {
		fmt.Println("Circle", instance.r)
	} else {
		fmt.Println("none")
	}
}
func getType2(s Shape) {
	switch instance := s.(type) {
	case Circle:
		fmt.Println("Circle", instance.r)
	case Triangle:
		fmt.Println("Triangle", instance.a, instance.b, instance.c)

	}

}

type Shape interface {
	peri() float64
	area() float64
}
type Triangle struct {
	a, b, c float64
}

func (t Triangle) peri() float64 {
	return t.a + t.b + t.c
}
func (t Triangle) area() float64 {
	var p = t.a + t.b + t.c
	return math.Sqrt(p * (p - t.a) * (p - t.b) * (p - t.c))
}

type Circle struct {
	r float64
}

func (c Circle) peri() float64 {
	return 2 * math.Pi * c.r
}
func (c Circle) area() float64 {
	return math.Pi * c.r * c.r
}
