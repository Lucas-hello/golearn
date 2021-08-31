//输入三个整数x,y,z，请把这三个数由小到大输出。

package main

import "fmt"

func main() {
	fmt.Println(sort(1, 2, 3))
}
func sort(x, y, z int) []int {

	var str []int
	var result []int
	str = append(str, x)
	str = append(str, y)
	str = append(str, z)
	for _, i := range str {
		if i >= x && i >= y && i >= z {
			result = append(result, i)

		}
	}

	for _, i := range str {
		if i <= x && i <= y && i <= z {
			result = append(result, i)

		}
	}
	return result
}
