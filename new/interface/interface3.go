package main

import "sort"

type Interface interface {
	sort.Interface
	Push(x interface{})
	Pob() interface{}
}

func main() {

}
