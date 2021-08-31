package main

import "fmt"

//func sum (a...int)int{
//	ret:=0
//	for _,arg := range a{
//		ret = ret+arg}
//	return ret
//}
//func main()  {
//	r1 := sum(12)
//r2:=sum(1,3)
//r3:= sum(12,34,56)
//fmt.Println(r1)
//fmt.Println(r2)
//fmt.Println(r3)
//
//}
//

func add(a int)int{
	a=a+1
	return a
}
func main()  {
	x:=3
	fmt.Println("x=",x)
	fmt.Println(add(x))

}