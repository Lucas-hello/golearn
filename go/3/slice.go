package main

import (
	"fmt"
)
func reverse(v []int){
	for i,z:=0,len(v)-1;z>=i;i,z=i+1,z-1{
		v[i],v[z]=v[z],v[i]
	}
}
func equal(x,y []string) bool{
	if len(x)!=len(y){
		return false
	}
	for i :=range x{
		if x[i]!=y[i]{
			return false
		}
	}
	return true
	
}

func main()  {
	month:=[...]string{1:"January",2:"February",3:"March",4:"April",5:"May",6:"June",7:"July",8:"August",9:"September",10:"October",11:"November",12:"December"}
	Q2:=month[4:7]
	Summer:=month[6:9]
	fmt.Println(Q2)
	fmt.Println(Summer)
	for _,i :=range Summer{
		for _,q:=range Q2{
			if i==q{fmt.Printf("%s is in both \n",i,)

		}
		}
	}
//fmt.Println(Summer[:20])//超出了cup
endlessSummmer:=Summer[:5]
fmt.Println(endlessSummmer)
k:=[]int{1,2,3,4,5,6,7,8,9}
reverse(k)
fmt.Println(k)
s:=[]int{0,1,2,3,4,5}
reverse(s[:2])
reverse(s[2:])
reverse(s)
fmt.Println(s)
r:=make([]int,3)[3:]
fmt.Println(r)
//make([]int,3,4)
}
