package main

import "fmt"

type persion struct{
	name string
	age int

}
//比较两个人的年龄，返回年龄大的那个人并返回年龄差
//struct也是传值的
func Older(p1 persion, p2 persion) (persion,int){
	if p1.age>p2.age{
		return p1,p1.age - p2.age
	}
	return p2,p2.age - p1.age
}
func main()  {
	var tom persion
	tom.name,tom.age="Tom",8
	bob:=persion{age: 25,name: "Bob"}
	paul:=persion{"Pull",43}
	tb_Older,tb_diff:= Older(tom, paul)
	bp_Older,bp_diff:=Older(bob,paul)
	tp_Older,tp_diff:=Older(tom,paul)
	fmt.Printf("Of %s and %s,%s is older by %d years\n",tom.name,bob.name,tb_Older.name,tb_diff)
	fmt.Printf("Of %s and %s,%s is older by %d years\n",bob.name,paul.name,bp_Older.name,bp_diff)
	fmt.Printf("Of %s and %s,%s is older by %d years\n",tom.name,paul.name,tp_Older.name,tp_diff)

}