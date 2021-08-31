package main

import (
	"crypto/sha256"
	"fmt"
)

type Currency int

const (
	USD Currency=iota
	EUR
	GBP
	RMB
	)

//func zero(prt *[32]byte)  {
//	for i :=range prt{
//		prt[i]=0
//	}
//}
func zero( prt *[32]byte) {
	*prt=[32]byte{}
}

func main() {
	 var a [3]int
	 fmt.Println(a[0])
	 fmt.Println(a[len(a)-1])
	 for i,v:=range a{
	 	fmt.Printf("%d %d \n",i,v)
	 }
	 for _,v:=range a{
	 	fmt.Printf("%d \n",v)
	 }
	 q:=[...]int{1,2,3,4}
	 fmt.Printf("%T\n",q)
	 symbol:=[...]string{USD:"$",EUR:"@",GBP:"&",RMB:"¥"}
	 fmt.Print(RMB,symbol[RMB],"\n")
	 k:=[2]int{1,2}
	 l:=[...]int{1,2}
	 m:=[...]int{1,3}
	 fmt.Print(k==l,k==m,k!=m,"\n")
	 //o:=[3]int{1,2}
	 //fmt.Print(k==o)
	 c1:=sha256.Sum256([]byte("x"))
	 c2:=sha256.Sum256([]byte("X"))
	 fmt.Printf("%x\n%x\n%t\n%T\n",c1,c2,c1==c2,c1)


 }
