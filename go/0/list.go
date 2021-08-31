//赋值，左右两边必须个数相等
//简短变量声明至少要声明一个新的变量
//用"_"来丢弃不需要的值
//元组赋值，等号左右两边个数必须相等
//运算的赋值可以简写++
package main

import "fmt"




func list(n int) int {
	if n <= 2 {
		return n
	}
	x, z := 0, 1
	for i := 0; i <= n; i++ {
		x, z = z, x+z

	}
	return z
}
func main() {
	var s int
	var d string
	s, d = 3, "S"
	f, g := 3, "S"
	fmt.Println(s, d, f, g)
	fmt.Println(list(10))
	fmt.Println(a)

}

a:=3
a=5

