//new一个匿名变量返回值是匿名变量的地址
//匿名变量初始化的值是new(T),T类型的零值
package main

import "fmt"

func main() {
	p := new(int)
	fmt.Println(p)

}
