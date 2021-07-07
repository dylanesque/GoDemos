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
	/*
			Write a program which reads information from a file and represents it in a
			slice of structs. Assume that there is a text file which contains a series
			of names. Each line of the text file has a first name and a last name, in
			that order, separated by a single space on the line.

		Each field will be a string of size 20 (characters).

		 Your program will successively read each line of the text file and create a struct which contains the
		first and last names found in the file. Each struct created will be added to a slice,
		and after all lines have been read from the file, your program will have a slice
		containing one struct for each line in the file. After reading all lines from the file,
		your program should iterate through your slice of structs and print the first and last
		names found in each struct.
	*/

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the filename")

	filename, _ := reader.ReadString('\n')
	csvfile, err := os.Open(strings.TrimSpace(filename))

	if err != nil {
		fmt.Println(err)
	}

	r := csv.NewReader(csvfile)

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(record)
	}
}
