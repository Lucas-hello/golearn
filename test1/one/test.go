package main

import "fmt"

func main()  {
	numbers:=make(map[string]int)
	numbers["one"]=1
	numbers["two"]=2
	numbers["three"]=3
	fmt.Println("第三个数字是：",numbers["three"])

}