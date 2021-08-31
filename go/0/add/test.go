package main

import (
	"bufio"
	"fmt"
	_ "go/0/add/function"
	"os"
)

func main() {
	in := bufio.NewScanner(os.Stdout)
	fmt.Println(in)

}
func init() {
	b := "中国ad"
	bs := []rune(b)
	fmt.Println(bs, len(b))
	for _, v := range b {
		fmt.Println(v)

	}

	//fmt.Println(os.Args)
}
