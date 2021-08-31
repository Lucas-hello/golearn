package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main(){
	wg3.Add(2)
	go incCounter3(1)
	go incCounter3(2)
	wg3.Wait()
	fmt.Println("Final Counter:",counter3)

}
var(
	counter3 int
	wg3 sync.WaitGroup
	mutex sync.Mutex
)
func incCounter3(id int){
	defer wg3.Done()
	for count:=0;count<2;count++{
		mutex.Lock()
		{
			value := counter3
			runtime.Gosched() //当前goroutine从线程退出并返回对列
			value++
			counter3 = value
		}
		mutex.Unlock()

	}
}
