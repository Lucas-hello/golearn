package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

type Expr interface{}

func main() {
	var w io.Writer
	w = os.Stdout
	f := w.(*os.File)
	c := w.(*bytes.Buffer)
	fmt.Println(f, c)

}
