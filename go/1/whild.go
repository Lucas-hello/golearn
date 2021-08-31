//整形int8 int16 是bit的取值范围不同
//%取膜和被取膜的符号是一致的
//&^运算符是是后边数只要的二进制位只要出现1前面数字那个位无论是什么都要替换成0
//+x是0+x 没有意义,-x是0-x
package main

import (
	"fmt"
	"math"
)

func main() {
	// x int = 999
	fmt.Println(math.MaxInt8)
	fmt.Println(math.MaxUint8)
	var x int8 = 1<<1 | 1<<5 //第一位是0
	var z int8 = 1<<1 | 1<<2
	fmt.Println(x)
	fmt.Println(z)
	fmt.Printf("%08b\n", x&^z) //只要z所在为为1的全部为0-
}
