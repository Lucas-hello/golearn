package main

import (
	"fmt"
)

func main() {
	b := Max(56, 65, 74, 53, 6, 66)
	fmt.Println(b)

}

func Max(vals ...int) int {
	max := 0
	for _, val := range vals {
		if val > max {
			max = val
		}
	}
	return max

}
func Min(vals ...int) int {
	min := 0
	for _, val := range vals {
		if val < min {
			min = val
		}
	}
	return min

}
