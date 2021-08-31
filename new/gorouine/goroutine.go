package main

import (
	"fmt"
	"runtime"
	"time"
)

func say(s string)  {
	for i:=0;i<5;i++{
		runtime.Gosched() //把cpu时间留给别人

		fmt.Println(s)
	}

}

func main() {

	say("hello")//携程怎么展现出来
	runtime.GOMAXPROCS(2)//告诉cpu使用两个线程
	go say("world")
	time.Sleep(time.Second)

}
