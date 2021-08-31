package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main(){
	fmt.Println(rand.Intn(100))
	wg4.Add(2)
	count:=make(chan int)//创建了一个无缓冲通道
	go player("Nadal",count)
	go player("Djokovic",count)
	count<- 1
	wg4.Wait()

}
var wg4 sync.WaitGroup
func init(){
	rand.Seed(time.Now().UnixNano())//以当前时间为rand初始化
}
func player(name string ,count chan int){
	defer wg4.Done()
	for{
		ball,ok :=<-count
		if!ok{
			fmt.Printf("Player %s Won\n",name)
			return
		}
		n:=rand.Intn(100)//随机取100以内的整数
		if n%13==0{
			fmt.Printf("Player %d Missed\n",ball)
			close(count)
			return

		}
		fmt.Printf("Player %s Hit %d\n",name,ball)
		ball++
		count<-ball
	}
}
