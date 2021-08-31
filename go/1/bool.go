//布尔值如果运算符左边可以确定整个布尔表达式的值那么运算符右边的表达式将不再被求值
//布尔值会被转换成0 或1 ，反之也不可以
//布尔值的零值时false
package main

import "fmt"

const ()

func btio(b bool) int {
	if b {
		return 1
	}
	return 0

}

func main() {
	var s = "abc"
	b := s != "" && s[0] == 'x'
	fmt.Println(b)

}
