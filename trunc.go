package main

import (
	"fmt"
)

func main() {
	var f float64
	fmt.Println("Enter a floating point number")
	_, err := fmt.Scanf("%f", &f)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%d", int32(f))

}
