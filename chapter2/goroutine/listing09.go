package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main(){
	wg.Add(2)
	go incCounter(1)
	go incCounter(2)
	wg.Wait()
	fmt.Println("Final Counter:",counter)

}
var(
	counter int
	wg sync.WaitGroup

)
func incCounter(id int){
	defer wg.Done()
	for count:=0;count<2;count++{
		value:=counter
		runtime.Gosched()//当前goroutine从线程退出并返回对列
		value++
		counter=value

	}
}
