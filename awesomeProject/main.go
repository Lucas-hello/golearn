package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func sayHello(w http.ResponseWriter,r *http.Request)  {
	ioutil.ReadFile("./hello.txt")
	fmt.Fprintln(w,  )
}
func main()  {
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":9094", nil)
	if err != nil {
		fmt.Printf("http sever fild err:%v\n",err)
		return
	}




}