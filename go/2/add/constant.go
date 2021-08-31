//常量的值不可以被修改
//常量可以和变量一样被批量声明
//常量所有的运算也是常量
package main

import (
	"fmt"
	"math"
	"time"
)

const noDelay time.Duration = 0
const timeout = 5 * time.Minute

const (
	s = 3
	d = 5
	k = 8
)

func main() {

	//s = 5
	fmt.Println(s)
	fmt.Printf("%T %[1]v\n", noDelay)
	fmt.Printf("%T %[1]v\n", timeout)
	fmt.Printf("%T %[1]v\n", time.Minute)
	fmt.Printf("%T\n", math.Pi)
	//a := d - s
	//b = a - 3
	//fmt.Println(b)
}
