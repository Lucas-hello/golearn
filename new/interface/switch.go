package main

import (
	"fmt"
	"strconv"
)

type Element interface {}
type List []Element
type Persion struct {
	name string
	age int
}

func (p Persion)String() string {
	return "(name:"+p.name+"- age:"+strconv.Itoa(p.age)+"years)"
}

func main()  {
	list:=make(List,3)
	list[0]=1
	list[1]="Hello"
	list[2]=Persion{"bill",54}
	for index,element:= range list{
		switch value :=element.(type){
		case int:fmt.Printf("list[%d]is an int and its value is %d\n",index,value)
		case string:fmt.Printf("list[%d] is a string and its value is %s\n",index,value)
		case Persion:fmt.Printf("lisr[%d] is a Persion and its value is %s\n ",index,value)
		default:
			fmt.Println("list[%d] is of a different type",index)




		}

	}
}


