package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {

	for _, url := range os.Args {
		resp, err := http.Get(url)
		fmt.Println("resp:", resp)
		fmt.Println("err:", err)

	}

}
