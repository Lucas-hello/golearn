//打印乘法表
package main

import "fmt"

func main() {
	for x := 1; x < 10; x++ {
		fmt.Println("")
		for y := 1; y <= x; y++ {
			fmt.Printf("%d*%d=%d  ", y, x, x*y)

		}

	}
	fmt.Println("")

}
