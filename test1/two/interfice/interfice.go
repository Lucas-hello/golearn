package main
//
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
//func (h Human)SayHi(){
//	fmt.Printf("Hi I am %s ,you can call me on %s\n",h.name,h.phones)
//}
//func (h*Human)Sing(lyrics string ){
//	fmt.Println("La,la,la,la,la,la,la...",lyrics)
//}
//func (h*Human)Guzzle(beerStein string)  {
//	fmt.Println("Guzzle,Guzzle,Guzzle,Guzzle,Guzzle...",beerStein)
//}
//func (e*Employee)SayHi()  {
//	fmt.Printf("Hi I am %s,I work at%s,you can call me at%s\n",e.name,e.company,e.phones)
//}
//func (s*Student)BorrowMoney(amount float64) {
//	s.loan += amount
//}
//func (e*Employee)SpendSalary(amount float64)  {
//	e.money-=amount
//}
//
//type Men interface {
//	SayHi()
//	Sing(lyrics string)
//	Guzzle(beerStein string)
//}
//type YongChap interface {
//	SayHi()
//	Sing(lyrics string)
//	BorrowMoney()(amount float64)
//
//}
//type ElderlyChap interface {
//	SayHi()
//	Sing(lyrics string)
//	SpendSalary(amount float64)
//}
//func main() {
//
//
//}
