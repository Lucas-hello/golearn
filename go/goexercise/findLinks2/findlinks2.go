package main

import (
	"fmt"
	"go/exercise/html"
	"net/http"
)

func main() {

}

func findLinks(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("getting %s:%s", url, resp.Status)

	}
	doc, err := html.Parse(resp.Body)
}
