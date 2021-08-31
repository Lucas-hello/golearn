//题目：有1、2、3、4个数字，能组成多少个互不相同且无重复数字的三位数？都是多少？
package main

import "fmt"

func main() {
	for i:=1 ;i<5;i++{
		for j:=1;j<5;j++ {
			for k:=1;k<5;k++{

					if (i!=j)&&(i!=j)&&(i!=k)&&(j!=k){
						fmt.Print(i*100+j*10+k,"\n")
					}

				}

			}

		}
	}

var number =[]int{1,2,3,4}




