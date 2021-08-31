package main

import "fmt"

func main()  {
	rating:=map[string]float32{"C":5,"Go":4.5,"python":4.5,"C++":2}
	pickup,ok:=rating["c#"]
	if ok{fmt.Println("C# is in the map and its rating is", pickup)
	}else{
		fmt.Println("C# is not in the map and its rating is",pickup)
	}

}