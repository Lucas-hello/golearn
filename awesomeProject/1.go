package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func ayHello(w http.ResponseWriter, r *http.Request) {
	ioutil.ReadFile("./hello.txt")
	fmt.Fprintln(w, )
}

func main() {
http.HandleFunc("/", ayHello)
err := http.ListenAndServe(":8087", nil)
if err != nil {
fmt.Printf("http server failed, err:%v\n", err)
return
}
}
