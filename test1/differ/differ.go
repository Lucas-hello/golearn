package main

import "os"

	var user = os.Getenv("USER")

func init()  {
	if user == ""{
		panic("on value for $USER")
	}

}
	func throwsPanic (f func())(b bool){
		defer func() {}()

	}
	


func main() {

}
