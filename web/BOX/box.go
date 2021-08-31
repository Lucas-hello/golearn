package main

import "fmt"

func main (){
	boxes:=BoxList{
		Box{4,4,4,RED},
		Box{10, 10, 1, YELLOW},
		Box{1, 1, 20, BLACK},
		Box{10, 10, 1, BLUE},
		Box{10, 30, 1, WHITE},
		Box{20, 20, 20, YELLOW},
	}
	fmt.Println(boxes.BiggestColor())

}

const (
	WHITE=iota//定义颜色常量
	BLACK
	BLUE
	RED
	YELLOW
)
type Color byte//定义 Color 类型底层是byte类型
type Box struct {//定义box类型
	width,height,depth float64
	color Color
}
type BoxList []Box//定义Boxlist,底层是box类型的切片

func (b Box)Volume()float64{//定义一个函数返回volume的体积
	return b.width*b.height*b.depth
}
func (b *Box)SetColor(c Color){//定义一个函数返回box的color
	b.color=c
}
func (bl BoxList)BiggestColor()Color{//定义一个函数返回体积最大盒子的颜色
	v:=0.00//初始化体积
	k:=Color(WHITE)//初始化衍射
	for _,b:=range bl{//用循环遍历boxlist列表
		if b.Volume()>v{//利用循环和判断找出体积最大的盒子和颜色
			v=b.Volume()
			k=b.color
		}
	}
	return k //返回颜色

}
func (bl BoxList)PaintItBlack(){//定义一个函数把所有box的颜色改成black
	for i,_:=range bl{
		bl[i].SetColor(BLACK)
	}
}
func (c Color)String ()string{//
	strings:=[]string{"WHITE","BLACK","BLUE","RED","YELLOW"}
	return strings[c]

}
