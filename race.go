package main

/* 

A race condition is when a program produces results that are inconsistent
due to unpredictable timing, which is usually traced back to poorly executed
asynchronous code.

This program creates a race condition by putting two goroutines together, which sends them 
to the background, which is where they'll stay and never execute, because the "synchronous" goroutine of 
main will complete. Try removing the "go" keyword from in front of the calls to tester() to see the 
effects that it has!

*/


import (
	"fmt"
	"time"
)

func tester(s string) {
	for i := 0; i < 25; i++ {
		fmt.Println(s, i);
		time.Sleep(time.Millisecond * 300)
	}
}

func main() {
	go tester("This should be first")
	go tester("This should be second")
	fmt.Println("And we're done")

}
