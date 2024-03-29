package main

import (
	"fmt"
	html "go/4/net"
	"net/http"
)

func main() {

}
func FindLinks(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("getting %s", url, resp.Status)

	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML:%V", url, err)
	}
	return visit(nil, doc), nil
}
