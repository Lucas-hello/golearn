package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func sayhelloName(w http.ResponseWriter,r *http.Request){
	_=r.ParseForm()//解析参数，默认时不会解析
	fmt.Println(r.Form)//是输出到服务器端的打印信息
	fmt.Println("path",r.URL.Path)
	fmt.Println("scheme",r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k,v:=range r.Form{
		fmt.Println("key:",k)
		fmt.Println("val:",strings.Join(v,""))
	}
	_,_=fmt.Fprintf(w,"Hello astaxie!")

}
func main()  {
	http.HandleFunc("/",sayhelloName)
	 err:=http.ListenAndServe(":9090",nil)
	 if err!=nil{
	 	log.Fatal("ListenAndSever:",err)
	 	c,err:=

	}

}
