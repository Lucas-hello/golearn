package main

import (
	"fmt"
	"os"
)

func main() {
	a := os.Args //获取命令行参数（一个string的slice）
	fmt.Println(a)
	fmt.Printf("%T\n", a)
	for {
		fmt.Println("1")
	}

}
