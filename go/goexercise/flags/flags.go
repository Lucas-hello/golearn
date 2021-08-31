package main

import (
	"fmt"
)

func main() {
	//var f float64 = 3 + 0i
	//f = 2
	//f = 1e123
	//f = 'a'
	//fmt.Println(5.0 / 9)
	//fmt.Println(KiB)
	//fmt.Println(MiB)
	//fmt.Println(GiB)
	//fmt.Println(TiB)
	//fmt.Println(PiB)
	//fmt.Println(EiB)
	fmt.Println(YiB / ZiB)
	//fmt.Println(YiB)
	//fmt.Printf("%T\n", YiB)
}

const (
	_ = 1 << (10 * iota)
	KiB
	MiB
	GiB
	TiB
	PiB
	EiB
	ZiB
	YiB
)
