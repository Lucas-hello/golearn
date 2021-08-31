package main

import "fmt"
var a =f()
func f() *int {
	v:=1
	return &v

}
func main()  {
	var x ,y int
	print(&x,"\n")
	fmt.Print(&x==&x,&x==&y,&x==nil,"\n")
	var z int
	p:=&z
	fmt.Print(p!=nil,"\n")

	fmt.Print(f(),"\n")
	v:=1
	incr(&v)
	fmt.Print(incr(&v))



}
func incr(p *int) int {
	*p++
	return *p

}