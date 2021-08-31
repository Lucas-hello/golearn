package main

func HasPrefix(s, prefix string) bool {
	if len(s) < len(prefix) {
		if prefix[:len(s)] == s {
			return true
		}
	}
	return false
}
func HasPrefix1(s, prefix string) bool {
	return len(s) < len(prefix) && s == prefix[:len(s)]

}
func HasSuffix(s, prefix string) bool {
	if len(s) < len(prefix) {
		if s == prefix[len(prefix)-len(s):] {
			return true
		}
	}
	return false
}
func HasSuffix1(s, prefix string) bool {
	return len(s) < len(prefix) && s == prefix[len(prefix)-len(s):]

}
func Contains(s, prefix string) bool {
	for i := 0; i < len(prefix); i++ {
		return HasPrefix(s, prefix[i:])
	}
	return false
}

func main() {

}
