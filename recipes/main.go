package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	fmt.Println(factorializeNumber(5))
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

// A simple recursive solution
func factorializeNumber(n int) int {
	if n == 0 {
		return 1
	}

	fact := n

	if (n > 0) {
		return fact * factorializeNumber(n -1)
	}

	return fact
}
