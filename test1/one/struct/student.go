package main

import "fmt"

type Human struct {
	name string
	age int
	wight int
}
type Student struct {
	Human
	expical string
}
func main(){
	mark:=Student{Human{"Mark",40,125},"Very clever"}
	fmt.Println("Student'name is",mark.name)
	fmt.Println("Student's age is",mark.age)
	fmt.Println("Studen's weight is",mark.wight)
	fmt.Println("Studen's expeciai is",mark.expical)
	mark.expical="he become a superman"
	fmt.Println("Student get a chance he is",mark.expical)
	mark.age=90
	fmt.Println("The time is past an his name is",mark.age)
	mark.wight=999
	fmt.Println("It's very danger his weight is ",mark.wight)
}