//4. 用 Golang 实现，斐波纳契数列
package main

import (
	"fmt"
)

func main() {
	ne := FibonacciSequence(45)
	fmt.Println(ne)

}

func FibonacciSequence(in int) int {
	x, y := 0, 1
	for i := 0; i > 0; i++ {
		x, y = y, x+y
		if x == in {
			return x
		}

		//fmt.Println(x)
	}
	return x

}
