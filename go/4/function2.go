package main

import "fmt"

func main()  {
	fmt.Printf("%T\n",add1)
	fmt.Printf("%T\n",oper)
	sun:=oper(10,20,add1)
	fmt.Print(sun)
}
func add1 (a int, b int) int {
	return a+b
}
func oper(a,b int ,fun func(int,int)int)int  {
	fmt.Println(a,b,fun)
	res:=fun(a,b)
	return res
	
}
