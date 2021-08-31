//有一分数序列：2/1，3/2，5/3，8/5，13/8，21/13...求出这个数列的前20项之和。
package main

import "fmt"

func main() {
	fmt.Println(end(6))

}
func up(n int) []float64 {
	var ints []float64
	n1, n2 := 1.0, 1.0
	for i := 0; i <= n; i++ {
		n1, n2 = n2, n1+n2
		ints = append(ints, n2)

	}
	return ints
}
func down(n int) []float64 {
	var ints []float64
	n1, n2 := 1.0, 1.0
	for i := 0; i <= n; i++ {
		n1, n2 = n2, n1+n2
		ints = append(ints, n1)
	}
	return ints
}

func number(a, b []float64) []float64 {
	var num []float64
	for i := 0; i < len(a); i++ {
		num = append(num, a[i]/b[i])

	}
	return num

}
func sum(num []float64) float64 {
	var an float64
	for _, v := range num {
		an = an + v
	}
	return an
}
func end(n int) float64 {
	a := up(n)
	b := down(n)
	num := number(a, b)
	return sum(num)
}
