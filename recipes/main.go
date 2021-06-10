package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	repeatAString("abc", 0)
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

	if n > 0 {
		return fact * factorializeNumber(n-1)
	}

	return fact
}

func findLongestWord(s string) int {
	words := strings.Split(s, " ")
	count := 0
	for _, word := range words {
		if len(word) > count {
			count = len(word)
		}
	}

	return count
}

func largestOfFour(a [][]int) []int {
	largestNumbers := make([]int, 4)
	for i, child := range a {
		for _, childNum := range child {
			if childNum > largestNumbers[i] {
				largestNumbers[i] = childNum
			}
		}
	}
	return largestNumbers
}

func confirmEnding(s string, t string) bool {
	ending := s[len(s)-1:]
	if ending == t {
		return true
	} else {
		return false
	}

}

func repeatAString(s string, n int32) {
	if n < 1 {
		fmt.Println("")
	} else {
		for i := 0; i <= len(s); i++ {
			fmt.Printf("%s", s)
		}
	}

}
