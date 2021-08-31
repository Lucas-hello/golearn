package main

import (
	"fmt"
	"strconv"
)

func main() {
	a := 123
	//b := string(a)
	y := fmt.Sprintf("%d", a)
	//x := fmt.Sprintf("%d", b)
	fmt.Println(y)
	fmt.Printf("%T\n", y)
	fmt.Println("")
	//fmt.Println(b)
	//fmt.Println(x)
	x, _ := strconv.Atoi("123")
	z, _ := strconv.ParseInt("123", 10, 64)
	fmt.Println(x)
	fmt.Println(z)
}
