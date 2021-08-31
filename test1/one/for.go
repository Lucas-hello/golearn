package main

import "fmt"

func main(){
	i:=10
	switch i {
	case 1:
		fmt.Println("i is equal to 1")
	case 2,3,4:
		fmt.Println("a is equal 2,3 or 4")
	case 10:
		fmt.Println("a is equal to 10")
	default:
		fmt.Println("All I know is that i is an inteager")

	
	}


}
// 输出:sum is equal to 45