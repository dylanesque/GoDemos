package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
		"reflect"
)

func main() {
	var nums []int
	fmt.Println("Enter a minumum of 8 space-separated integers")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	nums = numbers(scanner.Text())
	fmt.Println(split(nums))

}

func numbers(s string) []int {
    var n []int
    for _, f := range strings.Fields(s) {
        i, err := strconv.Atoi(f)
        if err == nil {
            n = append(n, i)
        }
    }
    return n
}

func split(s []int) (x, y []int) {
	first := s[:len(s) / 2]
	second := s[len(s) / 2:]
	BubbleSort(first)
	BubbleSort(second)
	return first, second
}

func BubbleSort(s []int) {
	isSliceSorted := false

	for !isSliceSorted {
		swapped := false
		for i := 0; i < len(s)-1; i++ {
			if s[i] > s[i+1] {
				Swap(s, i)
				swapped = true
			}
		}
		if !swapped {
			isSliceSorted = true
		}
	}
	fmt.Println("sorting is done")
}

func Swap(sl []int, i int) {
	swapI := reflect.Swapper(sl)
	swapI(i, i+1)
}

/*

Write a program to sort an array of integers.
The program should partition the array into 4 parts,
each of which is sorted by a different goroutine.
Each partition should be of approximately equal size.
Then the main goroutine should merge the 4 sorted subarrays
into one large sorted array.

The program should prompt the user to input a series of integers.
Each goroutine which sorts ¼ of the array should print the subarray
that it will sort. When sorting is complete, the main goroutine
should print the entire sorted list.

*/
