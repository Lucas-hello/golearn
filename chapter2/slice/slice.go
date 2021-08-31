package main

import "fmt"

func main()  {
	array:=[5]*int{0:new(int),1:new(int)}
	*array[0]=10
	var array1[5][2]int
	array1[0][1]=4
	var array3 [1e6]int
	foo(array3)
var arr =[5]int{1,2,3,4,5}
arr[3]=7
 ar:= arr[3:4]
 fmt.Println(ar)
 fmt.Println(cap(ar),len(ar))
 fmt.Printf("%T",ar)
 source:=[]string{"apple","orange","banana","greap","plum"}
 slice :=source[2:3:3]
 fmt.Println(slice)
 slice=append(slice,"kiwi")
 fmt.Println(slice)
 slice =append(slice,source...)
 fmt.Println(slice)
 slice1 :=[]int{10,20,30,40}
 for index,value:=range slice1{
 	fmt.Printf("Value:%d  Value-Addr: %X  ElemAddr: %X\n",value,&value,&slice1[index])
 	slice2:=[][]int{{10},{100,200}}
	 slice2[0]=append(slice2[0],20)
	 fmt.Println(slice2)
 }







}
func foo (array [1e6]int){

}