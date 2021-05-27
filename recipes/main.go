package main

import "fmt"

func main() {
	fmt.Println(celsToFaren(30))
}

func celsToFaren(i float64) float64 {
	return i * 9 / 5 + 32;
}


