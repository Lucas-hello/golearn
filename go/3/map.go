package main

import (
	"fmt"
	"sort"
)

func equals (x,y map[string]int)bool  {
	if len(x)!=len(y) {
		return false
	}
	for k,xv:=range x{
		if yv,ok:=y[k];!ok||yv!=xv {

		return false
		}
	}
	return true
}

func main()  {
	ages:=make(map[string]int)
	ages["alice"]=31
	ages["charlie"]=40
	delete(ages,"alice")
	ages["bob"]=ages["bob"]+1
	fmt.Print(ages["bob"],"\n")
	for name,age :=range ages{
		fmt.Printf("%s  %d",name,age)
	}
	//var names []string
	names:=make([]string,0,len(ages))
	for name :=range ages{
		names=append(names ,name)

	}
	sort.Strings(names)
	for _,name:=range names{
		fmt.Printf("%s\t%d\n",name,ages[name])
	}
	//age,ok:=ages["bob"]
	//if !ok{
	//seen :=make(map[string]bool)
	//input:=bufio.NewScanner(os.Stdin)
	//for input.Scan(){
	//	line:=input.Text()
		//if !=seen[line]{
	//		seen[line] = true
	//	}

	

	//}
}
