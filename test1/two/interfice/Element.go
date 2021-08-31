package main

import (

	"strconv"
)

type Element interface {}
type List Element
type Person struct {
	name string
	age int
}

func (p Person)String()string  {return "(name: " + p.name + " - age: "+strconv.Itoa(p.age)+ " years)"


}
//func main() {
//	list:=make(List,3)
//	list[0]=1
//	list[1]="hello"
//	list[2]=Person


}



