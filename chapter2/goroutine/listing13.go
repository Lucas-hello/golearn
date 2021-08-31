package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var (
	counter1 int64
	wg1 sync.WaitGroup
)


func main()  {
	wg1.Add(2)
	go incCounter1(1)
	go incCounter1(2)


wg1.Wait()
	 fmt.Println("Final Counter:",counter1)
}
func incCounter1(id int){
	defer wg1.Done()
	for count:=0;count<2;count++{
		atomic.AddInt64(&counter1,1)
		runtime.Gosched()
	}

}
