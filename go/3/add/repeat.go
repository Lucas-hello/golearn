//去掉重复
package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 4, 5, 6, 7, 8, 8, 9, 2, 3}
	// 4, 4, 5, 6, 7, 8, 8, 9, 2, 3
	// 4, 5, 6, 7, 8, 8, 9, 2, 3

	x := repeat(a)
	fmt.Println("before: ", x)
	fmt.Println("after: ", a)

}
func repeat(s []int) []int {
	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] {

			copy(s[i:], s[i+1:])

			fmt.Println("copy s:", s)
			s = s[:len(s)-1]
			fmt.Println("s:", len(s))
			fmt.Println("z:", s)
			fmt.Printf("i=%d len(s)-1=%d\n", i, len(s)-1)
			if i == len(s)-1 {
				fmt.Println("jj:", s)
				return s
			}
		}

	}
	return s

}
