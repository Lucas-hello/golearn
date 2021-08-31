//package main
//
//import (
//	"fmt"
//	"runtime"
//)
//
//func say(s string)  {
//	for i:=0 ;i<5;i++{
//		runtime.Gosched()
//		fmt.Println(s)
//
//	}
//
//}
//
//func main() {
//	go
//	say("hello")
//	say("world")
//
//
//}
//
package main

import (
	"fmt"
	"time"
)

func sayHello(name string){
	for i:=1;i<5;i++{
		fmt.Print("  hello  ",name,"    ",i,"\n")
	}
	
}
func main()  {
	 sayHello("gill")
	go sayHello("bill")
	 go sayHello("tina")
	 go sayHello("gdfcg")
	time.Sleep(time.Second)
}
