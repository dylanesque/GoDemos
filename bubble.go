package main

import (
	"fmt"
	"reflect"
)

func main() {
	nums := []int{2, 12, 1, 4, 3, 9, 7, 0, 5, 10}
	BubbleSort(nums)
	fmt.Println(nums)
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
