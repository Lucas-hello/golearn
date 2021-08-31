package main

import "fmt"

func main() {
	name := &Name{}
	age := &Age{}
	id := &Id{}
	Open(name)
	Open(age)
	Open(id)
}

type Database struct {
	Name *Name
	Age  *Age
	Id   *Id
}
type Name struct {
}

func (name *Name) Open() {
	fmt.Println("你看到来名字")

}

type Age struct {
}

func (age *Age) Open() {
	fmt.Println("你看到了年龄")

}

type Id struct {
}

func (id *Id) Open() {
	fmt.Println("你看到了id")
}

type Opened interface {
	Open()
}

func Open(open Opened) {
	open.Open()

}
