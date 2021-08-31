package main

import "fmt"

type IntList struct {
	Value int
	Tial  *IntList
}

func (list *IntList)sum()int {
	if list == nil{
		return 0
	}
	return list.Value+list.Tial.sum()

}

type Values map[string][]string //定义了一个map类型

func (v Values)Get(key string)string{
	if vs:=v[key];len(vs)>0{
		return vs[0]
	}
	return ""
}
func (v Values) Add(key,value string){
	v[key]=append(v[key],value)
}
func main()  {
	m:=Values{"lang":{"en"}}
	m.Add("item","1")
	m.Add("item","2")
	map1:=map[string]string{
		"1":"a",}
	map1["4"]="g"
	fmt.Println(map1)


}