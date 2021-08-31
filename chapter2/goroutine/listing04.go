package main

import (
	"fmt"
	"sync"

       "runtime")
var Wg sync.WaitGroup
func main()  {
	runtime.GOMAXPROCS(20)
	Wg.Add(2)
	go printPrime("A")
	go printPrime("B")
	Wg.Wait()
	fmt.Println(runtime.GOMAXPROCS(runtime.NumCPU()))


}
func printPrime(prefix string){
	defer Wg.Done()
next:
	for outer:=2;outer<5000;outer++{
		//fmt.Println("outer:",outer)

		for inner:=2;inner<outer;inner++{
			//fmt.Println("inner:",inner)
			if outer%inner==0{
				continue next
			}
		}
		fmt.Printf("%s:%d\n",prefix,outer)
		runtime.Gosched()
	}
	fmt.Println("Completed",prefix)
}

