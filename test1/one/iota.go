package main

import "fmt"

func main(){
	const (
		x= iota
		y= iota
		z=iota

	)
	fmt.Println(x,y,z)
}