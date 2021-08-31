package main

import "fmt"

func main() {
	a := [...]string{"a", "b", "c", "d", "e"}
	//c := &a
	d := a[:]
	rotate(d, 4)
	fmt.Println(a)
	//b := [...]*string{&a[0], &a[1], &a[2], &a[3], &a[4]}
	//reverse2(b)
	k := []int{0, 1, 2, 3, 4, 5}
	//reverse1(k[:2])
	//reverse1(k[2:])
	//reverse1(k)
	//fmt.Println(k)
	//reverse3(c)
	//
	//fmt.Println(a)
	reverse1(k)
	fmt.Println("-------------------", k)

}

func reverse1(s []int) {
	for f, l := 0, len(s)-1; f < l; f, l = f+1, l-1 {
		s[f], s[l] = s[l], s[f]
	}

}

//指针数组
func reverse2(s [5]*string) {
	for f, l := 0, len(s)-1; f < l; f, l = f+1, l-1 {
		*s[f], *s[l] = *s[l], *s[f]
	}
}

//数组指针
func reverse3(s *[5]string) {
	for f, l := 0, len(*s)-1; f < l; f, l = f+1, l-1 {
		(*s)[f], (*s)[l] = (*s)[l], (*s)[f]

	}

}

//编写一个rotate函数，通过一次循环完成旋转
func rotate(s []string, n int) []string {
	for i := 0; i < n; i++ {
		s = append(s, s[0])
		s = s[1:len(s)]
	}
	return s
}

//*s[f]
