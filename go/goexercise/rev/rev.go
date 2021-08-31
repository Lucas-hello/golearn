package main

import "fmt"

func main() {
	b := []string{"a", "b", "c", "d", "e"}
	b1 := [5]string{"a", "b", "c", "d", "e"}
	rotate(b, 5)
	reverse(b)
	fmt.Println(b)
	a := &b1

	reverse2(a)
	fmt.Println(b1)

}
func reverse(s []string) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}

}

func reverse2(s *[5]string) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}

}
func rotate(s []string, n int) {
	for i := 0; i < n; i++ {
		copy(s, s[1:])
		s[len(s)-1] = s[0]

	}
}
