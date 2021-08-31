//利润(I)低于或等于10万元时，奖金可提10%；
//
//利润高于10万元，低于20万元时，低于10万元的部分按10%提成，高于10万元的部分，可可提成7.5%；
//
//20万到40万之间时，高于20万元的部分，可提成5%；
//40万到60万之间时高于40万元的部分，可提成3%；
//
//60万到100万之间时，高于60万元的部分，可提成1.5%；
//
//高于100万元时，超过100万元的部分按1%提成从键盘输入当月利润，求应发放奖金总数？
package main

import "fmt"

func main() {
	var a float64
	var b bool = true
	fmt.Println("You can write")
	fmt.Scan(&a)
	var c, d, e, f, g, h float64
	c = a * 0.1
	d = 10*0.1 + (a-10)*0.075
	e = 10*0.1 + (a-10)*0.075 + (a-20)*0.05
	f = 10*0.1 + (a-10)*0.075 + (a-20)*0.05 + (a-40)*0.3
	g = 10*0.1 + (a-10)*0.075 + (a-20)*0.05 + (a-40)*0.003 + (a-60)*0.015
	h = 10*0.1 + (a-10)*0.075 + (a-20)*0.05 + (a-40)*0.003 + (a-60)*0.0015 + (a-100)*0.01


	switch b {

	case a < 5:
		fmt.Printf("You can get %f thounds \n", c)
	case a < 20: // false in
		fmt.Printf("You can get %f thounds\n", d)
	case a < 40: //false
		fmt.Printf("You can get %f thounds\n", e)
	case a < 60:
		fmt.Printf("You can get %f thounds\n", f)
	case a < 100:
		fmt.Printf("You can get %f thounds\n", g)
	default:
		fmt.Printf("You can get %f thounds\n", h)
		if a < 5 {

		}
	}
}