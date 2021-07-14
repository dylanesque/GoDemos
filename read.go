package main

import (
	"bufio"
	"fmt"
	"encoding/csv"
	"io"
	"os"
	"strings"
)

type name struct {
	fname string
	lname string
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the filename")

	filename, _ := reader.ReadString('\n')
	csvfile, err := os.Open(strings.TrimSpace(filename))

	if err != nil {
		fmt.Println(err)
	}

	r := csv.NewReader(csvfile)

	var names []name

	for {
		record, err := r.Read()

		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
		}
		
		for _, line := range record {
			nameParts := strings.Fields(line)
			var newName = name{fname: nameParts[0], lname: nameParts[1]}
			names = append(names, newName)
		}
	}
	fmt.Println(names)

}
