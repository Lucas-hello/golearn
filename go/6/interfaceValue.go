package main

import (
	"fmt"
	"io"
	"os"
)

//type Writer interface {
//	Write(p []byte) (n int, err error)
//}
//Stdout = NewFile(uintptr(syscall.Stdout), "/dev/stdout")
func main() {
	var w io.Writer
	fmt.Println(w)
	fmt.Printf("%T\n", w)
	//w.Write(p []byte("hello"))
	w = os.Stdout
	var b US
	fmt.Printf("%T\n", b)
	fmt.Println(b)
}

type e struct {
	E string
}
type mouse struct {
	name string
}

func (m mouse) input() {
	fmt.Printf(m.name)

}

type US interface {
	input()
	output()
}
