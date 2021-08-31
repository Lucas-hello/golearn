//求1+2!+3!+...+20!的和
package main

import "fmt"

func main() {
	fmt.Println(sumInt(20))
}
func sumInt(n int) int {
	sums := 0
	for i := 0; i <= n; i++ {
		sums = sums + i

	}
	return sums
}
