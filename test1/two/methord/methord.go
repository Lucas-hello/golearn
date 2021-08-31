package main

import (
	"fmt"
)

func main() {
	a:=[]int{1,2,3,4,5,6,7,8,9,10}
	b:=a[0:5]
	fmt.Println(a)
	fmt.Print(b)
	b[4]=8
	fmt.Println(b)
	fmt.Println(a)
	b =  append(b,70)
	fmt.Println(b)
	supermarket:=[]string{"apple","banana","orange","greap","pulm"}
	news:=supermarket[2:3:4]
	fmt.Println(news)
	fmt.Printf("%s\n",append(
		b,a...))
	for index,value:=range a {
		fmt.Printf("index: %d   value: %d \n",index,value)
	}
	for inx:=len(a)-1 ; inx >0 ;inx--{
		fmt.Printf("index: %d   value: %d \n",inx,a[inx])
	}


}
