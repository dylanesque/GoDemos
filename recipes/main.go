package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	arrs := [][]int{
		{13, 27, 18, 26},
		{4, 5, 1, 3},
		{32, 35, 37, 39},
		{1000, 1001, 857, 1},
	}
	fmt.Println(largestOfFour(arrs))
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
	largestNumbers := make([] int, 4)
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
	// confirmEnding("Bastian", "n") should return true
}
