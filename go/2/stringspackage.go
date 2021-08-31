package main

import (
	"fmt"
	"strconv"
)

func main()  {
	a:=123
	y:=fmt.Sprintf("%d",a)
	fmt.Print(y,strconv.Itoa(a),"\n")
	fmt.Println(strconv.FormatInt(int64(a),2))
	b:=fmt.Sprintf("a=%d",a)
	fmt.Print(b)


}
