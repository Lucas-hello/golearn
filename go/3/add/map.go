//map 是一个无序的key对value的集合
//一个man中所有的key是相同的类型，所有的value也是相同的类型
//map不能进行取址因为map的增长，从新分配更大的内存空间
//可以通过make来创建一个map
//key 必须支持==比较运算，来判断key是否存在
//使用内置函数delete（name，key）可以删除元素
//通过key访问对应的value，如果key不存在则返回value对应类型的零值
//map的零值是nil,一个nil 值的map不能向其中倒入元素
//
package main

import (
	"fmt"
)

func main() {
	var b = []string{"w", "b", "e", "k"}
	k(b)
	fmt.Println(b)
	//seen := make(map[string]bool)
	//input := bufio.NewScanner()
	arge := make(map[string]int)
	//args1 := map[string]int{}
	arge["alice"] = 31
	arge["charlie"] = 45
	arge["alice"] = 32
	fmt.Println(arge)
	delete(arge, "alice")
	fmt.Println(arge)
	arge["bob"]++
	fmt.Println(arge)
	for name, age := range arge {
		fmt.Printf("%s %d \n", name, age)
	}
	arg, ok := arge["bob"]
	if ok {
		fmt.Println(arg)
	}
	if !ok { /*"bob" is not a key in this map; age ==0. */
	}

}
func equal(x, y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}
	for k, xv := range x {
		if yv, ok := y[k]; !ok || yv != xv {
			return false
		}
	}
	return true

}

var m = make(map[string]int)

func k(list []string) string {
	p := fmt.Sprintf("%q", list)
	fmt.Printf("%T\n", p)
	return p
}
func Add(list []string) { m[k(list)]++ }
func Count(list []string) int {
	return m[k(list)]

}
