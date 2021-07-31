package main

import (
	"fmt"
	"sync"
)

type Chopstick struct {
	mu sync.Mutex
}

type Philosopher struct {
	leftCS, rightCS *Chopstick
	number int
	eatCount int
}

func (p Philosopher) eat() {
	for {
		p.leftCS.mu.Lock()
		p.rightCS.mu.Lock()

		fmt.Println("Starting to eat ", p.number)

		p.leftCS.mu.Unlock()
		p.rightCS.mu.Unlock()
		p.eatCount += 1;
		fmt.Println("Finishing eating ", p.number)

	}
}

func main() {
	CSticks := make([] *Chopstick, 5)

	for i := 0; i < 5; i++ {
		CSticks[i] = new(Chopstick)
	}

	philos := make([] *Philosopher, 5)

	for i := 0; i < 5; i++ {
		philos[i] = &Philosopher{CSticks[i], CSticks[(i+1)%5], i + 1, 0}
	}

	for i := 0; i < 5; i++ {
		go philos[i].eat()
	}
}

func host() {
	// iterate over philosophers
	/* 
	let them eat if:

	a) they have a free chopstick on each side
	b) they have eaten less than three times
	c) there are less than 2 currently eating
	*/
}


/*

The philosophers pick up the chopsticks in any order, not
lowest-numbered first (which we did in lecture).

*/
