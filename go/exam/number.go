package main

import "fmt"

func main() {
	count("aaaaacccvvv")

}

func print(m map[string]int) {
	for k, v := range m {
		fmt.Printf("%s   :   %d\n", k, v)
	}
}
func CountStrings(s string) map[string]int {
	var m = make(map[string]int)
	for _, v := range s {
		num := m[string(v)]
		num++
		m[string(v)] = num
	}
	return m
}
func count(s string) {
	m := CountStrings(s)
	print(m)
}
