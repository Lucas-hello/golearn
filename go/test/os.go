//Echo1 print its command-line arguments
package main

import (
	"fmt"
	"os"
	"strings"
)

func main()  {
	var s, sep string
	for i:=1;i<len(os.Args);i++{
		s+=sep+os.Args[i]
		sep=""//初始化
	}
	fmt.Println(s)
	fmt.Println(strings.Join(os.Args[1:],""))

	for index,value:=range os.Args{
		fmt.Println(index,value,"\n")
	}
}
