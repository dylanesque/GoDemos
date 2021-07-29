package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
)

func main() {
	var nums []int
	fmt.Println("Enter a minumum of 8 space-separated integers")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	nums = numbers(scanner.Text())
	firstHalf, secondHalf := split(nums)
	firstSl, secondSl := split(firstHalf)
	thirdSl, fourthSl := split(secondHalf)

	sorted := []int{}

	go BubbleSort(firstSl)
	go BubbleSort(secondSl)
	go BubbleSort(thirdSl)
	go BubbleSort(fourthSl)

	sorted = append(sorted, firstSl...)
	sorted = append(sorted, secondSl...)
	sorted = append(sorted, thirdSl...)
	sorted = append(sorted, fourthSl...)

	BubbleSort(sorted)
	fmt.Println(sorted)

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
	first := s[:len(s)/2]
	second := s[len(s)/2:]
	return first, second
}

func BubbleSort(s []int) {
	fmt.Println("Sorting ", s)
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

