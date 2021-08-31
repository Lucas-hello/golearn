// 编写一个函数，判断两个字符串是否是是相互打乱的，也就是说它们有着相同的 字符，但是对应不同的顺序。
package main

import "fmt"

func same(a, b string) bool {
	for i := 0; i < len(a); i++ {
		for _, value := range b {
			an := int32(a[i])
			if value == an {
				return true
			}
			//fmt.Println(i1)
		}
	}
	return false
}
func main() {
	fmt.Println(same("aebcd", "efgh"))
}
