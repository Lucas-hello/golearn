package main

//import "fmt"
//
//type Human struct {
//	name string
//	age int
//	phones string
//}
//type Student struct {
//	Human
//	school string
//	loan float64
//}
//type Employee struct {
//	Human
//	company string
//	money float64
//}
//
//func (h Human)SayHi()  {
//	fmt.Printf("Hi I am %s ,you can call me on %s\n",h.name,h.phones)
//}
//func (h Human)Sing(lyrics string)  {
//	fmt.Println("la,la,la,la...",lyrics)
//}
//func (e Employee)SayHi()  {
//	fmt.Printf("Hi I am %s,I am work at%s,you can call me on%s\n",e.name,e.company,e.phones)
//
//}
//
//type Men interface {
//	SayHi()
//	Sing(lyrics string)
//}
//func main() {
//	mark := Student{Human{"Mark", 20, "1232312312"}, "MIT", 0.00}
//	bill := Student{Human{"Bill", 25, "1232433454"}, "WILL", 59.3}
//	shark := Employee{Human{"Shark", 87, "987654343"}, "IEO", 99999}
//	sum := Employee{Human{"Sum", 45, "98765678"}, "OIU", 88987}
//	var i Men
//	i=mark
//	i.SayHi()
//	i.Sing("November Ring")
//	i=sum
//	i.SayHi()
//	i.Sing("hefsd")
//	fmt.Println("Let's use a slice of Men and see what happens")
//	x:=make([]Men,3)
//	x[0],x[1],x[2]=bill,shark,sum
//	for _,value:=range x{
//		value.SayHi()
//	}
//
//}