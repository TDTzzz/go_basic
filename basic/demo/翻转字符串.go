package main

import "log"

func main() {
	reverseString("abcdefg")
}

func reverseString(s string) (string, bool) {
	str := []rune(s)
	l := len(str)

	for i := 0; i < l/2; i++ {
		str[i], str[l-1-i] = str[l-1-i], str[i]
	}
	log.Println(string(str))
	return string(str), true
}
