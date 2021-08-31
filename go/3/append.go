package main

import "fmt"
func appendInt(x []int,y int)[]int{
	var z []int
	zlen:=len(x)+1
	if zlen<=cap(x){
		z=x[:zlen]
	}else {
		zcap:=zlen
		if zcap<2*len(x){
			zcap=2*len(x)
		}
		z = make([]int,zlen,zcap)
		copy(z,x)
	}
	z[len(x)]=y
	return z
}
func nonempty(strings []string)[]string  {
	i:=0
	for _,s :=range strings{
		if s!=""{
			strings[i]=s
			i++
		}
	}
	return strings[:i]
}
func nonempty2(strings []string) []string  {
	out:=strings[:0]
	for _,s :=range strings{
		if s!=""{
			out=append(out,s)
		}
		}
	return out
}
func remove(slice []int,i int) []int {
	slice[i]=slice[len(slice)-1]
	return slice[:len(slice)-1]
}
func removes(slice []int, i int) []int {
	copy(slice[i:],slice[i+1:])
	return slice[:len(slice)-1]

}

func main() {

	var runes []rune
	for _,i :=range "Hello, 世界"{
		runes=append(runes,i)
	}
	fmt.Printf("%q\n",runes)

}