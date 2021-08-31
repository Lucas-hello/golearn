package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(basename("a/b/c/d/e.d.go"))

}
func basename(s string) string {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}

	}
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}
	return s

}
func basename1(s string) string {
	slash := strings.LastIndex(s, "/")
	s = s[slash+1:]
	if don := strings.LastIndex(s, "."); don >= 0 { //为什么用if判断
		s = s[:don]
	}
	return s
}
