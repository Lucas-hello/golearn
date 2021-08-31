package main

import "fmt"

func main() {

	b := []string{"a", "b", "c", "d", "e"}
	b = make([]string, 5, 10)
	b[0] = "a"
	b[1] = "b"
	b[2] = "c"
	b[3] = "d"
	b[4] = "e"
	//c := whirl(b, 3)
	//fmt.Println(c)
	whirl1(b, 2)
	fmt.Println(b)

}
func whirl(s []string, n int) []string {
	for i := 0; i < n; i++ {
		s = append(s, s[0])
		//fmt.Println(s)

		s = s[1:]
		fmt.Println(s)
	}
	fmt.Println(s)
	return s

}
func whirl1(s []string, n int) []string {
	s = append(s, "999")
	//fmt.Println(&s)
	//for i := 0; i < n; i++ {
	//
	//	s = append(s, s[0])
	//	fmt.Println("a:", &s)
	//
	//	s = s[1:]
	//	fmt.Println("b:", &s)
	//}
	//fmt.Println("c:", &s)

	fmt.Println("s:", s)
	return s

}
