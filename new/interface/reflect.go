package main

import "reflect"

func main() {
	t:=reflect.TypeOf(i)
	v:=reflect.ValueOf(i)
	tag:=t.Elem().Field().Tag
	name:=v.Elem().Field().String()
}
