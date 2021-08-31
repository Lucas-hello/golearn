package main

import "fmt"

type testlnt func(int)bool
func isOdd(intege int) bool {
	if intege%2==0{
		return false

	}
	return true
}

func isEven(integer int)bool{
	if integer%2==0{
		return false
	}
	return true

}

func filter(slice []int,f testlnt)[]int  {
	var result []int
	for _,value:=range slice{
		if f(value){
			result=append(result,value)
		}
	}
	return result

}
func main(){
	slice:=[]int{1,2,3,4,5,7}
	fmt.Println("slice=",slice)
	Odd:=filter(slice,isOdd)
	fmt.Println("Odd elements of slice are:",Odd)
	even:=filter(slice,isEven)
	fmt.Println("Even elements of slice are:",even)


}