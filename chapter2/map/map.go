package main

import "fmt"

func main(){

	map1:=make (map[string]int)
	fmt.Println(map1)
	map1["blue"]=1
	map1["two"]=1
	for key,value:=range map1{
		fmt.Println(key,":",value)

	}
	//dict := map[int][]string{}
	var color map[string]int
	fmt.Println(color)
	color=map1
	value,ok:=color["blue"]
	if ok{
		fmt.Println(value)
	}
	delete(color,"blue")
	//var dur Duration
	//dur=int64(1000)

}

type Duration int64
