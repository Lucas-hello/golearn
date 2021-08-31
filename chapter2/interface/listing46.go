package main

import "fmt"

func main(){
	//duration(42).pretty()

}
type duration int
func (d *duration)pretty()string{
	return fmt.Sprintf("Duaation:%d",*d)

}
