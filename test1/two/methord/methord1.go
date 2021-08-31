package main

import "fmt"

//
////type Rectangle struct {
////	width, height float64
////}
////type Circle struct {
////	radiues float64
////
////}
////func(c Rectangle)area()float64{
////	return c.width*c.height
////}
////func (r Circle)area()float64 {
////	return r.radiues*r.radiues
////}
////
////
////
////func main() {
////	c1:=Rectangle{23,5}
////	c2:=Rectangle{44,23}
////	r1:=Circle{22}
////	r2:=Circle{70}
////	fmt.Print("Area of c1 is:",c1.area(),"\n" )
////	fmt.Print("Area of c2 is:",c2.area(),"\n")
////	fmt.Print("Area of r1 is:",r1.area(),"\n")
////	fmt.Print("Area of r2 is:",r2.area(),"\n")
////
////}
//const(
//	WHITE=iota
//	BLACK
//	BULE
//	RED
//	YELLOW)//定义一些常量
//
//
//type Colour byte //定义Colour类型底层格式是byte
//type Box struct {//定义Box 类型底层格式是struct
//	width,height,depth float64
//	colour Colour
//}
//type BoxList []Box //定义BoxList类底层结构是[]Box 的切片
//
//func (b Box)Volume() float64 {
//	return b.width*b.height*b.depth//定义一个Volume 函数计算box的体积，参数b是Box的格式返回值的格式是float64
//
//}
//func (b*Box)SetColour(c Colour)  { //定义一个SetColour 函数 返回值是c 返回c是Colour格式
//	 b.colour=c
//}
//
//
//func(bi BoxList) BiggestColour() Colour{//定义一个BiggestColour 的函数接受的是bi 格式是BoxList(一个box格式的切片 )
//	v:=0.0                         //初始化v
//	k:=Colour(WHITE)//初始化k,把WHITE 转换成Colour 格式
//	for _,b:=range bi{   //循环BoxList列表，忽略索引
//		if b.Volume()>v{
//			v=b.Volume()//求出面积最大的盒子
//			k=b.colour
//		}
//
//	}
//	return k
//}
//func (bi BoxList)PaintltBlack()  {
//	for i,_:=range bi{
//	bi[i].SetColour(BLACK)
//
//	}
//
//}
//func (c Colour)String()string {
//	strings:=[]string{"WHITE","BLACK","RED","BLACK","YELLOW"}
//	return strings[c]
//}
//func main()  {
//	boxes := BoxList{
//		Box{4, 4, 4, RED},
//		Box{10, 10, 1, YELLOW},
//		Box{1,1,20,BLACK},
//		Box{10,10,1,BULE},
//		Box{10,30,1,WHITE},
//		Box{20,20,20,WHITE},
//	}
//	fmt.Printf("We have %d boxes in our set\n",len(boxes))
//	fmt.Println("The Valume of first is",boxes[0].Volume(),"cm")
//	fmt.Println("The colour of the first one",boxes[len(boxes)-1].colour.String())
//	fmt.Println("The biggest one is",boxes.BiggestColour().String())
//	fmt.Println("Let us paint them all black")
//	boxes.PaintltBlack()
//	fmt.Println("The colour of second one is",boxes[1].colour.String())
//	fmt.Println("Obviously, now, the biggest one is",boxes.BiggestColour().String())
//}
type Human struct {
	name     string
	age      int
	phonemes string

}

type Students struct {
	Human
	school string
}

type Employee struct {
	Human
	company string

}

func (h Human)SayHi(){
	fmt.Printf("Hi I am %s , you can call me on %s\n",h.name,h.phonemes)
	}
func main()  {
	mark:=Students{Human{"Mark",20,"3123241223"},"MIT"}
	bill:=Employee{Human{"Bill",80,"3334545454"},"Golang lnc"}
	
	mark.SayHi()
	bill.SayHi()
}