//每隔三个数加一个点
package main

import "fmt"

func main() {
	a := "12345643345643567432543553245.35465464335435"
	b := judge(a)

	fmt.Println(b)
	//c := comma1("1254676754335")
	//fmt.Println(c)
	////d:=comma2("12.43535353")
	//b1 := choise1("-2324353342432423435332423")
	//fmt.Println(b1)
}

//处理整数
func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}

	for i := len(s) - 3; i > 0; i = i - 3 {
		s = s[:i] + "," + s[i:]
		//fmt.Println(s)
	}
	return s
}

//有问题
func comma1(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

type a interface {
}

//处理负数
func commaNeg(s string) string {
	n := len(s) - 1
	p := s[1:]
	if n <= 3 {
		return s
	}
	for i := 3; i < len(p); i = i + 4 {
		p = p[:i] + "," + p[i:]

	}
	return "-" + p
}

//处理浮点数
func commaFlo(s string) string {
	var k1, k2 *string
	for i := 3; i < len(s)-1; i++ {
		if s[i] == '.' {
			s1 := s[:i]
			s2 := s[i+1:]
			if len(s1) < 3 {
				k1 = &s1
			}
			for i1 := 3; i1 < len(s1); i1 = i1 + 4 {
				s1 = s1[:i1] + "," + s1[i1:]
				k1 = &s1
			}
			if len(s2) < 3 {
				k2 = &s2
			}
			for i2 := 3; i2 < len(s2); i2 = i2 + 4 {
				s2 = s2[:i2] + "," + s2[i2:]
				k2 = &s2
			}

		}
	}
	return *k1 + "." + *k2

}
func judge(s string) string {
	for i := 0; i < len(s)-1; i++ {
		i1 := int32(s[i])

		if i1 == '-' {
			return commaNeg(s)
		} else if i1 == '.' {
			return commaFlo(s)

		}

	}

	return ""
}
