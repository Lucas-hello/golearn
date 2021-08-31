//ci:=make(chan string)
//cn:=make(chan int)
//ch:=make(chan interface{})
//ch<-v//发送v到channel ch
//v:=<-ch//从ch接收数据并赋值给v
package main

import "fmt"

//func sum(a []int,c chan int)  {
//	sum:=0
//		for _,v:=range a{
//		sum +=v
//	}
//	c<-sum
//}

func test(x []int) {
	fmt.Println(x)
}

func main() {
	//a:=[]int{7,2,8,-9,4,0}
	//c:=make(chan int)
	//go sum(a[:len(a)/2],c)
	//go sum(a[len(a)/2:],c)//先执行下面
	//x,y:=<-c,<-c
	//fmt.Println(x,y,x+y)

	go test([]int{1, 2})
	go test([]int{3, 4})
	go test([]int{5, 6})

	for{}


}
