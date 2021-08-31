package main

import "fmt"

func main() {
	s := "abcdddd"
	ret := countChar(s)
	print(ret)

	s = "a空空abbdd中国中"
	//ret = countCharChinese(s)
	print(ret)

}

func print(m map[string]int) {
	for k, v := range m {
		fmt.Printf("char: %s, count: %d\n", k, v)
	}

}

func countChar(s string) map[string]int {
	m := map[string]int{}
	for _, c := range s {
		total := m[string(c)]
		total++
		m[string(c)] = total
	}

	return m
}

func countCharChinese(s string) map[string]int {
	m := map[string]int{}
	runes := []rune(s)

	//fmt.Printf("len(s) = %d, len(runes) = %d\n", len(s), len(runes))

	for i := 0; i < len(runes); i++ {

		k := string(runes[i])
		total := m[k]
		total++

		m[k] = total
	}
	return m
}
