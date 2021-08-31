package main

import "fmt"

func main() {
	var f float32 = 34353246 //整数大于23个bit位时bit32将出现误差
	fmt.Println(f == f+1)

}
