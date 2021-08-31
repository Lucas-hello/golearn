//package main
//
//import "fmt"
//
////func main(){
////	x:="12345"
////
////	fmt.Println(x)
////	var y string
////	for i := len(x);i>=0;i--{
////		y=y+x[i-1:i]
////
////	}
////	fmt.Printf(x)}
//
//
//func mber(a string) string {
//	var b string
//	for i := len(a);i>0;i--{
//		b = b+a[i-1:i]
//	}
//	return b
//
//}
//
//func main(){
//	a1:="12345"
//	fmt.Println(a1)
//	fmt.Println("over is",(a1))
//
//
//
//
//}
//
package main

import "fmt"

func flip(a []int)[]int  {
	for c,d:=0,len(a)-1;d>c;c,d=c+1,d-1{
		a[c],a[d]=a[d],a[c]

	}
	return a
}

func main()  {
	k:=[]int{1,2,3,4,5,6,7}
	fmt.Print(flip(k))

}