package main

import (
	"fmt"
	"os"
)

func pri() {
	var s, sr string
	for i := 0; i < len(os.Args); i++ {
		s = os.Args[i] + sr
		sr = ""
		fmt.Println(s)

	}

}
func pria() {
	for index, value := range os.Args {
		fmt.Println(index, value)
	}

}

func main() {
	pri()
	pria()

}
