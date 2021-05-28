package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	fmt.Println(temperatureConverter(30))
}

func temperatureConverter(i float64) float64 {
	return i*9/5 + 32
}

// As noted in this SO thread(https://stackoverflow.com/questions/1752414/how-to-reverse-a-string-in-go), reversing a string
// in Go is far from straightforward.
func reverseString(s string) string {
	size := len(s)
	buf := make([]byte, size)
	for start := 0; start < size; {
		r, n := utf8.DecodeRuneInString(s[start:])
		start += n
		utf8.EncodeRune(buf[size-start:], r)
	}
	return string(buf)
}
