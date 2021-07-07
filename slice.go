package main

import (
	"fmt"
	"sort"
)

func main() {
	s := []int{}
	var x rune
	for i := 0; x != 'X'; i++ {
		n := 0
		fmt.Println("Enter an integer")
		fmt.Scan(&n)
		if n == 0 {
			fmt.Println("Stopping program")
			break
		}
		s = append(s, n)
		sort.Ints(s)
		fmt.Println(s)
	}
}
