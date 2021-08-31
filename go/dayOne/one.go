package main

import (
	"fmt"
)

func main() {
	//First(10)
	//Second(10)
	//CycleArray([6]string{"ewr", "sd", "s", "sf", "sdf", "sdr"})
	//FizzBuzz(100)
	PrintA(5)
	fmt.Println(CountStr("asSASA ddd dsjkdsjs dk你好"))
	fmt.Println(ChangerStr("asSASA ddd dsjkdsjs dk你好"))
	fmt.Println(ReverseStr("abcdefg"))
	fmt.Println(AverageSlice([]float64{1.22, 3.44, 4.55, 5.3423, 4543.324}))
}

//First  创建一个基于 for 的简单的循环。使其循环 10 次，并且使用 fmt 包打印出计数器的值。

func First(n int) {
	for i := 1; i < n+1; i++ {
		fmt.Println(i)
	}

}

//Second  用 goto 改写 1 的循环。关键字 for 不可使用。
func Second(n int) {
	i := 1
here:

	fmt.Println(i)
	i++

	if i == n+1 {
		return

	}
	goto here
}

//Third  再次改写这个循环，使其遍历一个 array，并将这个 array 打印到屏幕上。

func CycleArray(arr [6]string) {
	for _, val := range arr {
		nval := fmt.Sprintf("%q", val)
		fmt.Println(nval)

	}

}

//FizzBuzz  编写一个程序，打印从 1 到 100 的数字。当是三的倍数就打印 “Fizz”代替数字，
//当是5的倍数就打印 “Buzz” 。当数字同时是三和5的倍数时，打印 “FizzBuzz” 。
func FizzBuzz(n int) {
	for i := 1; i < n+1; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")

		} else if i%3 == 0 {
			fmt.Println("Fizz")

		} else if i%5 == 0 {
			fmt.Println("Buzz")

		} else {
			fmt.Println(i)
		}
	}
}

//PrintA 建立一个 Go 程序打印下面的内容（到 100 个字符）：  A
// AA
// AAA
// AAAA
// AAAAA
// AAAAAA
// AAAAAAA

func PrintA(n int) {
	for i := 0; i <= n-1; i++ {
		for le := 0; le <= i; le++ {
			fmt.Print("A")
		}
		fmt.Println("")
	}
}

//CountStr 建立一个程序统计字符串里的字符数量：
// asSASA ddd dsjkdsjs dk你好
// 同时输出这个字符串的字节数。

func CountStr(s string) (int, int) {
	str := []rune(s)
	return len(s), len(str)
}

//ChangeStr 扩展/修改上一个问题的程序，替换位置 4 开始的三个字符为 “abc”。
func ChangerStr(s string) string {
	var news string
	if len(s) <= 4 {
		fmt.Println(s)
	} else {
		news = s[:3] + "abc" + s[6:]
	}
	return news

}

//ReverseStr  编写一个 Go 程序可以逆转字符串，例如 “foobar” 被打印成 “raboof”

func ReverseStr(s string) string {
	str := []rune(s)
	var rs string

	for f, l := 0, len(str)-1; f < l; f, l = f+1, l-1 {
		str[f], str[l] = str[l], str[f]
		fmt.Println(str)
	}
	for i := 0; i < len(str); i++ {
		si := fmt.Sprintf("%q", str[i])

		rs += si
	}
	return rs
}

// AverageSlice 编写计算一个类型是 float64 的 slice 的平均值的代码。
func AverageSlice(sli []float64) float64 {
	var ave float64
	for _, i := range sli {
		ave += i

	}
	lens := float64(len(sli))

	return ave / lens

}
