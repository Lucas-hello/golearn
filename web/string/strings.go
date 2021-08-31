package main

import "fmt"

func main (){
	slice:=[]int{1,2,3,4,5,7}
	fmt.Println("slice=",slice)
	odd :=filter(slice,isOdd)
	fmt.Println("Odd elements of slice are:=",odd)



}
type testlnt func(int)bool//定义一个函数类型的

func isOdd(integer int)bool{
	if integer%2==0{
		return false
	}
	return true
}
func isEven(integer int)bool{
	if integer%2==0{
		return true
	}
	return false
}
func filter(slice []int,f testlnt)[]int{
	var result []int//声明一个int类型的slice
	for _,value:=range slice{//遍历这个slice
		if f(value){//如果f函数返回结果是true就执行if里面的
			result=append(result,value)
		}
	}
	return result
}