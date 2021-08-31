package main

import "fmt"

func main() {
	fmt.Println(comma("fddsfefd"))

}
func comma(s string) string {
	lenS := len(s)
	if lenS <= 3 {
		return s
	}
	return comma(s[:lenS-3]) + "," + s[lenS-3:]

}
