//不同的类型不能在一起比较运算
//类型可以通过T（）来转换类型但是原类型和新类型要有相同的底层类型，或有相同结构的指针类型
//T（）只会改变类型但是不会改变类型的值
//手写之母大写说明是对外的，在外部包也可以使用

package main

import "fmt"

type cat struct {
	name string
	age  int
}
type bird struct {
}

func main() {
	fmt.Println("")
	b := "xyz"
	c := b[0] == 'x'
	fmt.Println(b[0])
	fmt.Println(c)

}
