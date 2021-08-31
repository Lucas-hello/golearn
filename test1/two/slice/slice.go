package main

import "fmt"

//func main(){
//	a:=[]int{1,2,3,4,5,6,7,8}
//	for index:=len(a)-1;index>=0;index--{
//		fmt.Printf( "%d   %d\n",a[index])
//
//	}
//}

//func number(b []int)() {
//	var index1 int
//	for index1 := len(b)-1; index1 >= 0; index1-- {
//
//	}
//	return index1
//}
//func main()  {
//	c:=[]int{1,2,3,4,5,6,7,8}
//
//	fmt.Print(number(c))
//
//}
func revire (str []int)[]int{
	for i,j:=0,len(str)-1; j>i;i,j=i+1,j-1{
		str[i],str[j]=str[j],str[i]
	}
	return str
}

func main()  {
	a:=[]int{1,2,3,4,5,6,7,}
	fmt.Println("a len=", len(a), cap(a))
	fmt.Print(revire(a))


}