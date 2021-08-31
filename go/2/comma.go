//无法运行,每隔3个数字加一个，
package main

import "fmt"

func comma(s string)string  {
	n:=len(s)
	fmt.Println(n,s)
	if n<=3{
		return s
	}
	return comma(s[:n-3])+","+s[n-3:]
}

func comma1(s string) string {
	n := len(s)
	fmt.Println(n, s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

func main()  {
	b:= comma1("1233")
	fmt.Println(b)

}
