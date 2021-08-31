package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func main(){
	wg2.Add(2)
	go doWork("A")
	go doWork("B")
	time.Sleep(100*time.Second)
	fmt.Println("Shutdown now")
	atomic.StoreInt64(&shutdown,1)


 wg2.Wait()
}


var (
	shutdown int64
	wg2 sync.WaitGroup
)

func doWork(name string){
	defer wg2.Done()
	for {
		fmt.Printf("Doing %s Work\n",name)
		time.Sleep(2500*time.Millisecond)
		if atomic.LoadInt64(&shutdown)==1{
			fmt.Printf("Shutting %s Down\n",name)
			break
		}

	}

}