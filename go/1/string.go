//字符串可以用==和<来运算，比较的是编码的字节
//字符串可以用len取长度
//
package main

import "fmt"

func main() {
	s := "hello world"
	fmt.Println(len(s))
	fmt.Println(s[0])
	t := s[:5]
	fmt.Println(s > t)
}
