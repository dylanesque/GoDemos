package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter a string")
	str, _ := reader.ReadString('\n')

	hasA := strings.Contains(strings.ToLower(str), "a")
	startsWithI := strings.HasPrefix(strings.ToLower(str), "i")
	endsWithN := strings.HasSuffix(strings.ToLower(strings.TrimSuffix(str, "\n")), "n")

	if hasA && startsWithI && endsWithN {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}