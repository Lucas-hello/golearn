//字符串可以用运算符比较，比较的是编码值
//字符串可以用[:]来取里面的值
//字符串可以用t[]来取某个值返回的是一个编码的值
//处理字符串的函数主要有4个，1，strings 2，tybes 3，strconv 4,unicode
//数字转换成字符串 用Fprintf来格式化数字
//用strconv.itao()转换
//将字符串解析成数字strconv.Atoi() 或者strconv.parseInt(  ,  ,)
package main

import (
	"fmt"
)

func main() {
	var c []byte = []byte{'a', 'b', 'x'}
	fmt.Println(c)
	basename("a/b/c.go")

}

func basename(s string) string {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			fmt.Println(s)
			break
		}
	}
	for i := len(s) - 1; i >= 0; i-- {
		fmt.Println(s)
		if s[i] == '.' {
			s = s[:i]
		}
	}
	return s
}
