package main

import (
	"fmt"
	"strings"
)

//func basename(s string) string {
//	for i:=len(s)-1;i>=0;i--{
//		if s[i]=='/'{
//			s=s[i+1:]
//			break
//			}
//		}
//	for i:=len(s)-1;i>=0;i--{
//		if s[i]=='.'{
//			s=s[:i]
//			break
//		}
//	}
//	return s
//}
func basename(s string) string {
	slash:=strings.LastIndex(s,"/")
	s=s[slash+1:]
	if dot:=strings.LastIndex(s,".");dot>=0{
		s=s[:dot]
	}
	return s
	
}

func main()  {
	a:="hello, world"
	fmt.Print(a[0:8],"\n")
	fmt.Print(a[9],a[2],"\n")
	b:=strings.ToUpper(a)
	fmt.Print(b,"\n")
	c:=strings.ToLower(b)
	fmt.Print(c)
	 //d:=unicode.Upper
	 //fmt.Print(d)
	 fmt.Println(basename("a/b/c.go"))

	 
}
