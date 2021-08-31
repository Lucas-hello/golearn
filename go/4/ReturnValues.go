package main

import "fmt"

type rectangle struct {
	Length float64
	Width  float64
}

func perimeter(p rectangle) float64{
	return p.Width*2+p.Length*2
}
func area (a rectangle) float64{
	return a.Length*a.Width
}

func getSum (num ...int){
	//fmt.Printf("%T\n",num)
	sum:=0
	for i:=0;i<len(num);i++{
		sum=sum+num[i]
	}
}
func main()  {
	getSum()
	arr1:=[4]int{1,2,3,4}
	fmt.Println("函数调用前数组的数据",arr1)
	fun1(arr1)
	fmt.Println("函数调用后数组的数据",arr1)
	fmt.Println("--------------------------------------------------------------------------------------")
	s1:=[]int{1,2,3,4}
	fmt.Println("函数调用前切片的数据",s1)
	fun2(s1)
	fmt.Println("函数调用后切片的数据",s1)
}
func fun1(arr2 [4]int){
	fmt.Println("函数中数组的数据",arr2)
	arr2[0]=100
	fmt.Println("函数中，数组修改后的数据",arr2)
}
func fun2(s2 []int){
	fmt.Println("函数中切片的数据",s2)
	s2[0]=100
	fmt.Println("函数中，切片修改后数组的数据",s2)

}