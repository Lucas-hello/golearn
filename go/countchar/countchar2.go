package main

import "fmt"

func main() {
	a := "aaavdf中中国"
	car := countChar2(a)
	print2(car)

}

func print2(m map[string]int) {
	for k, v := range m {
		fmt.Printf("%s : %d\n", k, v)
	}
}

func countChar2(s string) map[string]int {
	m := make(map[string]int)
	for _, v := range s {
		total := m[string(v)]
		total++
		m[string(v)] = total
	}
	return m

}
