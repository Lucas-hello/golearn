package main

import (
	"fmt"
	"strconv"
)





type Human struct {
	name string
	age int
	phones string
}

func (h Human)String()string  {
	return"《"+h.name+"-"+strconv.Itoa(h.age)+"-"+h.phones+"》"
	
}
func main()  {
	Bob:=Human{"Bob",20,"234354546"}
	fmt.Print("This human is",Bob)
	
}