//结构体是一种聚合类型，是由零零至多个任意类型的实体聚合而成的
//如果结构体里面的成员的名字是以大写字母开头的那么说名这个成员是导出的
//结构体里的成员不能包含结构体本身
//结构体的零值是结构体成员每个成员对应的零值
//没有任何成员就是一个空结构体
//结构体可以顺序为结构体成员指定面值，还可以用成员的名字和面值初始化，但是两种方式不能混合使用
//如果在函数内部想改变结构体成员的面值的化必须是指针传入
//go语言中所有函数的参数都是拷贝传入
//如果结构体里面的全部成员都是可以比较的那么结构体也是可以比较的
//一个结构体里面不能有两个相同的匿名成员
//

package main

import "fmt"

type address struct {
	hostname string
	port     int
}

type Point struct {
	X, Y int
}

func main() {
	var a address
	fmt.Print(a, "\n")
	p := Point{1, 2}
	q := Point{2, 1}
	fmt.Println(p.X == q.Y && p.Y == q.X)
	fmt.Println(p == q)
	hist := make(map[address]int)
	hist[address{"ssfsfsf", 332}]++
}
