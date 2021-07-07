package main

import (
	"fmt"
	"encoding/json"
	"bufio"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter a name")
	name, _ := reader.ReadString('\n')
	fmt.Println("Enter an address")
	address, _ := reader.ReadString('\n')
	addressBook := map[string]string {
		"name": name,
		"address": address,
	}
	jsonBook, err := json.Marshal(addressBook)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf(string(jsonBook))

}
