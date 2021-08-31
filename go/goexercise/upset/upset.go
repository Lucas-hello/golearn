package main

import (
	"fmt"
	"reflect"
)

func main() {
	s1 := "aaabbbvvv"
	s2 := "bavbavbav"
	fmt.Println(judSame(s2, s1))

}

func judge(m1, m2 map[string]int) bool {
	for k, m1v := range m1 {
		if m2v, ok := m2[k]; !ok || m2v != m1v {
			return false

		}

	}
	return true

}

func judSame(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	m1 := countChar(s1)
	m2 := countChar(s2)
	return judge(m1, m2)
}

func judge1(m1, m2 map[string]int) bool {
	return reflect.DeepEqual(m1, m2)
}

func countChar(s string) map[string]int {
	m := make(map[string]int)
	for _, v := range s {
		total := m[string(v)]
		total++
		m[string(v)] = total
	}
	return m
}
