//5. 生成一个随机数，让一个用户去猜这个数是多少？
//当用户输入的数字小于随机数，打印：您输入的数字太小
//当用户输入的数字大于随机数，打印：您输入的数字太大
//当用户输入的数字等于随机数，打印：恭喜您，答对了， 程序结束
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println(time.Now())
	rand.Seed(time.Now().UnixNano())
	d := rand.Intn(100)

	//fmt.Println(d)
	var n int
	for i := 0; i >= 0; i++ {
		fmt.Print("请输入数字：")
		_, _ = fmt.Scan(&n)

		if n == d {
			fmt.Println("恭喜您，答对了")
			break
		} else if n < d {
			fmt.Println("您输入的数字太小")
		} else if n > d {
			fmt.Println("您输入的数字太大")

		}
	}

}
