package main

import "fmt"

func main() {
	data := []string{"1", "", "2", "3"}
	fmt.Println(nonempty2(data))
	fmt.Println(data)

}
func nonempty(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}

	}
	return strings[:i]

}
func nonempty2(strings []string) []string {
	out := strings[:0]
	for _, s := range strings {
		if s != "" {
			out = append(out, s)
		}
	}
	return out
}
