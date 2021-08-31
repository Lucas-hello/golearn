package main

import "fmt"

func main() {
	m1 := Mouse{"redfish"}
	fmt.Println(m1.name)
	f1 := FlashDisk{"dell"}
	fmt.Println(f1.name)

	testInterface(f1)
	var b USB
	USB.start(f1)
	USB.start(m1)

	fmt.Printf("%T\n", b)

}

type USB interface { //定义一个USB接口
	start()
	end()
}
type Mouse struct { //定义一个鼠标类型
	name string
}
type FlashDisk struct { //定义一个u盘类型
	name string
}

func (m Mouse) start()     { fmt.Println(m.name, "鼠标") } //定义鼠标的的方法
func (m Mouse) end()       { fmt.Println(m.name, "停止鼠标") }
func (f FlashDisk) start() { fmt.Println(f.name, "u盘") } //定义u盘的方法
func (f FlashDisk) end()   { fmt.Println(f.name, "停止U盘") }
func testInterface(usb USB) { //定义一个函数调用方法
	usb.start()
	usb.end()

}
