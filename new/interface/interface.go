package man

import "fmt"





type Human struct {
	name string
	age int
	phones string
}
type Student struct {
	Human
	school string
	loan float64

}
type Employee struct {
	Human
	company string
	money float64
}

func (h Human) SayHi() {
	fmt.Printf("Hi I am %s,you can call me on %s\n",h.name,h.phones)

}
func (h Human) Sing (lyrics string)  {
	fmt.Print("La,la,la,la,la...",lyrics)
}
func (e Employee)SayHi()  {
	fmt.Printf("Hi I am %s ,I work at %s ,you can call me on%s\n",e.name,e.company,e.phones)
}
func (s Student)SayHi()  {
	fmt.Printf("Hi I am %s,I study at %s,you can call me on%s\n",s.age,s.school, s.phones)
}

type Men interface{
	SayHi()
	Sing(lyrics string)
}

func main()  {
	mark:=Student{Human{"Mark",29,"1232434323"},"HBD",3.55}
	bob:=Student{Human{"Mark",23,"1232332435"},"897",99.2}
	giany:=Employee{Human{"Giany",54,"543543453"},"FBI",90000}
	giles:=Employee{Human{"Giles",90,"34354325435"},"BBC",80000}

	var i Men
	i = mark
	fmt.Print("This is Mark,a Students:")
	i.SayHi()
	i=giles
	fmt.Print("This is Giels, a Emoloyee")
	i.SayHi()
	i.Sing("Lemon tree")
	fmt.Println("Let use slice of Men and see what happens")
	x:=make([]Men,3)
	x[0],x[1],x[2]=bob,giles,giany
	for index,value:=range x{
		fmt.Print(index)
		value.SayHi()

	}

}
