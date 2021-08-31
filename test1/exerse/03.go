//题目：输入某年某月某日，判断这一天是这一年的第几天？
//
//2.程序分析：以3月5日为例，应该先把前两个月的加起来，然后再加上5天即本年的第几天，
//
//特殊情况，闰年且输入月份大于3时需考虑多加一天。
package main

import (
	"fmt"
)

func main() {
	var (
		month, year, day, feb, sum int
	)

	fmt.Println("Please write date")
	fmt.Scan(&month,&year ,&day)
	if year%4==0 &&year%100!=0||year%400==0{
		feb=29
	} else {
	feb=28
	}
	switch month {
	case 1:sum=day
	case 2:sum= day + 31
	case 3:sum= feb + 31 + day
	case 4:sum= feb + 31 + 31+ day
	case 5:sum= feb + 31 +31+30+day
	case 6:sum= feb + 31 +31+30+31+day
	case 7:sum= feb + 31 +31+30+31+30+day
	case 8:sum= feb + 31 +31+30+31+30+31+day
	case 9:sum= feb + 31 +31+30+31+30+31+31+day
	case 10:sum= 31 + feb + 31 + 30 + 31 + 30 + 31 + 31 + day + 30
	case 11:sum= 31 + feb + 31 + 30 + 31 + 30 + 31 + 31 + 30 + day + 31
	case 12:sum= 31 + feb + 31 + 30 + 31 + 30 + 31 + 31 + 30 + 31 + 30 +day


	}
	fmt.Printf("This day is the %d day of the year",sum)
	}
