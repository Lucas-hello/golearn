package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main()  {

	runtime.GOMAXPROCS(1)//分配一个逻辑处理器给调度器使用
	var wg sync.WaitGroup//用来等待程序完成
	wg.Add(2)//说明要有两个goroutine
	fmt.Println("Start Goroutine")

	go func() {
		defer wg.Done()//用来通知main，函数已经完成
		for count:=0;count<3;count++{
			for char:='a';char<'a'+26;char++{
				fmt.Printf("%c",char)
			}
		}
	}()
	go func() {
		defer wg.Done()
		for count:=0;count<3;count++{
			for char:='A';char<'A'+26;char++{
				fmt.Printf("%c",char)
			}
		}
	}()
fmt.Println("Waiting To Finish")
wg.Wait()//等待所有的goroutine完成
fmt.Println("\nTerminating Program")

}
