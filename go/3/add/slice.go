//slice 切片是一个可变长度的序列，序列中的元素的类型都是形同的
//一个slice是由3个部分组成的，一个是slice的指针指向底层数组的的第一个元素的地址，然后是slice的长度和slice的容量
//slice的长度不能超过容量
//slice的切片操作可以用slice[i,j]来进行，slice的长度会改变，但是容量不会改变i>0&&j<len(slice)
//slice 可以通过内置函数cap 和 len 函数来求出slice的容量和长度
//slice在创建的时候没有指明长度，但是slice会隐式的创建一个值个数的数组
//slice之间不能比较所以，不能用运算符来判断两个slice是否相等，但是slice可以和nil来比较
//一个零值的slice==nli，一个nil值的slice的slice没有底层数组
//append函数可以在slice的最后添加零至多个元素
//可以用make函数来创建一个指定长度的数组make([]T,len)
package main

import "fmt"

func main() {
	var a = []int{1, 2, 3}
	b := &a
	fmt.Printf("%T=%d\n", b, b)
	months := [...]string{1: "january", 2: "february", 3: "march", 4: "april", 5: "may", 6: "june", 7: "july", 8: "august", 9: "september", 10: "october", 11: "november", 12: "december"}
	fmt.Println(cap(months), len(months))
	months1 := months[1:]
	fmt.Printf("%T\n", months1)
	fmt.Println(cap(months1), len(months1))
}

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
