package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"time"
)

type Writer interface {
	Write(p []byte) (n int, err error)
}
type Reader interface {
	Read(p []byte) (n int, err error)
}
type Closer interface{ close() error }
type ReaderWriter interface {
	Reader
	Writer
}
type ReaderWriterCloser interface {
	Reader
	Writer
	Closer
}

func main() {

	var w io.Writer
	w = os.Stdout
	//Writer(os.Stdout)
	w = new(bytes.Buffer)
	//w = time.Second
	fmt.Println(w)
	var x interface{} = time.Now()
}

//func main()  {
//	w=
//
//}
