//数组是有固定长度特定类型元素组成的序列
//数组每个元素可以通过索引的方式访问
//数组的每个元素初始化都会被初始化对应类型的零值
//[...]用来表示该类型的长度是元素的个数
//可以用指定一个索引的方式来初始化，但是打印的时候会按索引的顺序输出
//数组中的元素可以用运算符来比较的，但是必须是相同类型
//两个数组相比较的时候必须是两个数组的顺序类型个数和值完全相同的时候才会相等
package main

import "fmt"

type Currency int

const (
	USD Currency = 1
	EUR Currency = 2
	GBP Currency = 3
	RMB Currency = 4
)

func main() {
	fmt.Println(EUR)
	symbol := [...]string{EUR: "$", USD: "€", GBP: "£", RMB: "¥"}
	fmt.Println(symbol)
	for i, v := range symbol {
		fmt.Println(i, v, "\n")
	}
	//r := [...]int{99: -1}
	//fmt.Println(r)
	//s := [...]int{1, 2, 3}
	//t := [...]int{3, 2, 1}
	//fmt.Println(s == t) //值，类型顺序完全相等
	//fmt.Println(len(r))
	//var a [3]int
	//for i, v := range a {
	//	fmt.Printf("%d %d \n", i, v)

	//}
}
