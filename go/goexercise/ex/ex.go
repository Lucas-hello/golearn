package main

import (
	"fmt"
	"strings"
)

func main() {
	//f := square
	//fmt.Println(f(3))
	//f = negative
	//fmt.Println(f(3))
	//fmt.Printf("%T\n", f)
	//var f func(int) int
	//f(3)
	fmt.Println(strings.Map(add1, "HAL-9000"))
	fmt.Println(strings.Map(add1, ""))

}
func square(n int) int     { return n * n }
func negative(n int) int   { return -n }
func product(m, n int) int { return m * n }
func add1(r rune) rune     { return r + 1 }
