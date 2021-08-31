//指针式另一个变量的地址
//在函数体内的变量每回调用函数地址都是不同的
//改变指针对应的变量可以改变指针所对应的值
//任何类型的的指针的零值是nil
//如果指针相等要么是指针指向同一块地址，要么是全部为零值
//指针不等于nil,说名指针指向一个有效的变量
package main

import "fmt"

func main() {
	var star []string
	star = []string{"a", "b", "c"}
	b := &star[0]
	d := &star

	fmt.Println(b, "___________", d)
	fmt.Println(f() == f())
	fmt.Println(p)
	v := 1
	incr(&v)
	fmt.Println(incr(&v))

}

var p = f()

func f() *int {
	v := 1
	return &v
}
func incr(p *int) int {
	*p++
	return *p

}
