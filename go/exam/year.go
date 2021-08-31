//输入日期打印年份
package main

import "fmt"

func main() {
	var b Date
	b.month = October
	b.day = 25
	fmt.Println(NumberDay(b))

}

type Month int
type Date struct {
	month Month
	day   int
}

const (
	January Month = iota
	February
	March
	April
	May
	June
	July
	August
	September
	October
	November
	December
)

func NumberDay(d Date) int {

	switch d.month {
	case January:
		if d.day < 31 {
			return d.day
		}
	case February:
		if d.day < 28 {
			return 31 + d.day
		}
	case March:
		if d.day < 31 {
			return 31 + 28 + d.day
		}
	case April:
		if d.day < 30 {
			return 31 + 28 + 31 + d.day
		}
	case May:
		if d.day < 31 {
			return 31 + 28 + 31 + 30 + d.day
		}
	case June:
		if d.day < 30 {
			return 31 + 28 + 31 + 30 + 31 + d.day
		}
	case July:
		if d.day < 31 {
			return 31 + 28 + 31 + 30 + 31 + 30 + d.day
		}
	case August:
		if d.day < 31 {
			return 31 + 28 + 31 + 30 + 31 + 30 + 31 + d.day
		}
	case September:
		if d.day < 30 {
			return 31 + 28 + 31 + 30 + 31 + 30 + 31 + 31 + d.day
		}
	case October:
		if d.day < 31 {
			return 31 + 28 + 31 + 30 + 31 + 30 + 31 + 31 + 30 + d.day
		}
	case November:
		if d.day < 30 {
			return 31 + 28 + 31 + 30 + 31 + 30 + 31 + 31 + 30 + 31 + d.day
		}
	case December:
		if d.day <= 31 {
			return 31 + 28 + 31 + 30 + 31 + 30 + 31 + 31 + 30 + 31 + 30 + d.day
		}
	default:
		fmt.Println("请输入正确月份")

	}
	fmt.Println("请输入正确日期")
	return 0

}
