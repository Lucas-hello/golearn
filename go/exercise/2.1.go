// 向tempconv包添加类型、a
//常量和函数用来处理Kelvin绝对温度的转换
//，Kelvin 绝 对零度是−273.15°C，
//Kelvin绝对温度1K和摄氏度1°C的单位间隔是一样的。
//package main
//
//func main()  {
//	a:="abd"
//	str
//
//}

package main

import "fmt"

type User struct {
	Age int
	Name string
	X map[string]int
	Slice []int
	Arr [4]int
}

func get() *User {
	return &User{}
}

func get1() User {
	return User{}
}

func main()  {

	// *User user的指针
	a := [4]*User{get(), get(), get()}
	// nil
	fmt.Println("a[3]=", a[3]) // 8ge0

	// User user人值
	aa := [4]User{get1(), get1(), get1()}
	// User{Age:0}
	fmt.Printf("aa[3]=%#v\n", aa[3]) // nil

	b := [4]int{1, 2, 3}
	fmt.Println("b[3]=",b[3]) // 0

	c := [4]string{"a", "b", "c"}
	// ""
	fmt.Println("c[3]=",c[3],"==") // nil

	d := [4]int{}
	// []int{0, 0, 0, 0}
	fmt.Println("d=", d) // 空白

	//var q int
	//var s []int{}
	//[]
	//fmt.Println("s=", s)

	s1 := make([]int, 0, 10)
	fmt.Println("s1=", s1)

	s1 = make([]int, 2, 10)
	fmt.Println("s1=", s1)
	//len(s1)==0





}
