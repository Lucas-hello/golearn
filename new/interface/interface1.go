package main

import (
	"fmt"
	"strconv"
)

type Stringer interface{
	String() string
}

func (h Human)String()string  {
	return "《"+h.name+"-"+strconv.Itoa(h.age)+"year - ✆"+h.phones+"》"

}
func main()  {
	Bob:=Human{"Bob",39,"12323214213"}
	fmt.Println("This Human is:",Bob)
}


