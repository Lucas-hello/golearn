package main

import (
	"encoding/json"
	"fmt"
)

type Movie struct {
	Title string
	Year  int
	Color bool
	Actors []string
}



func main() {
	var movies = []Movie{
		{Title: "Casablanca", Year: 1942, Color: false, Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
		{Title: "Cool Hand Luke", Year: 1967, Color: true, Actors: []string{"Paul Newman"}},
		{Title: "Bullitt", Year: 1968, Color: true, Actors: []string{"Steve McQueen", "Jacqueline Bisset"}}}
	//	data,err :=json.MarshalIndent(movies,"","   ")
	//	if err !=nil{
	//		log.Fatalf("JSON marshaling failed:%s",err)
	//		}
	//		fmt.Sprintf("%s\n",string(data))
	//
	//}
jsonByte,_:= json.Marshal(movies)
jsonStr:= string(jsonByte)
fmt.Println(jsonStr)

fmt.Println([]string{"name","age"})
var str =`{"Title":"Casablanca","Year":1942,"Color":false,"Actors":["Humphrey Bogart","Ingrid Bergman"]},{"Title":"Cool Hand Luke","Year":1967,"Color":true,"Actors":["Paul Newman"]},{"Title":"Bullitt","Year":1968,"Color":true,"Actors":["Steve McQueen","Jacqueline Bisset"]}`
	var s1 Movie
err:=json.Unmarshal([]byte(str),&s1)
if err!=nil{
	fmt.Print(err)
}
	fmt.Printf("%v\n",s1)
	fmt.Println(s1.Year)



}


