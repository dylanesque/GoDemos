package main

import (
	"fmt"
	"sync"
)

// Is this right?
type Chopstick struct {
	mu sync.Mutex
}

type Philosopher struct {
	leftCS, rightCS *Chopstick
}

func (p Philosopher) eat() {
	for {
		p.leftCS.mu.Lock()
		p.rightCS.mu.Lock()

		fmt.Println("I am now eating")

		p.leftCS.mu.Unlock()
		p.rightCS.mu.Unlock()

	}
}

func main() {
	CSticks := make([] *Chopstick, 5)

	for i := 0; i < 5; i++ {
		CSticks[i] = new(Chopstick)
	}

	philos := make([] *Philosopher, 5)

	for i := 0; i < 5; i++ {
		philos[i] = &Philosopher{CSticks[i], CSticks[(i+1)%5]}
	}

	for i := 0; i < 5; i++ {
		go philos[i].eat()
	}
}

/*
There should be 5 philosophers sharing chopsticks,
with one chopstick between each adjacent pair of philosophers.

Each philosopher should eat only 3 times
(not in an infinite loop as we did in lecture)

The philosophers pick up the chopsticks in any order, not
lowest-numbered first (which we did in lecture).

In order to eat, a philosopher must get permission from a host which executes in its own goroutine.

The host allows no more than 2 philosophers to eat concurrently.

Each philosopher is numbered, 1 through 5.

When a philosopher starts eating (after it has obtained necessary locks) it prints
“starting to eat <number>” on a line by itself, where <number> is the number of the philosopher.

When a philosopher finishes eating (before it has released its locks) it prints
“finishing eating <number>” on a line by itself, where <number> is the number of the philosopher.
*/
