package main

import (
	"fmt"
	"sync"
)
var wg = sync.WaitGroup{}

type Host struct {
	seenEating int
	mu sync.Mutex
}

func (h *Host) permitEating() bool {
	if h.seenEating >= 2 {
		return false;
	}
	h.mu.Lock();
	h.seenEating += 1;
	h.mu.Unlock();
	return true;
}

func (h *Host) serveAnother() {
	h.mu.Lock()
	h.seenEating -= 1
	h.mu.Unlock();
}

type Chopstick struct {
	mu sync.Mutex
}

type Philosopher struct {
	leftCS, rightCS *Chopstick
	host Host
	number int
	helpingsEaten int
}

func (p Philosopher) eat(helpings int) {
	for i := 0; i < helpings; i++ {
		for !p.host.permitEating() {}
		p.leftCS.mu.Lock()
		p.rightCS.mu.Lock()

		fmt.Println("Starting to eat ", p.number)

		p.leftCS.mu.Unlock()
		p.rightCS.mu.Unlock()

		p.helpingsEaten += 1;
		fmt.Println("Finishing eating ", p.number)

		p.host.serveAnother()

	}
	wg.Done();
}

func main() {
	tonightsHost := Host{}

	// Creates chopsticks and philosophers
	CSticks := make([] *Chopstick, 5)

	for i := 0; i < 5; i++ {
		CSticks[i] = new(Chopstick)
	}

	philos := make([] *Philosopher, 5)

	for i := 0; i < 5; i++ {
		philos[i] = &Philosopher{CSticks[i], CSticks[(i+1)%5], tonightsHost, i + 1, 0}
	}

	wg.Add(5);


	for i := 0; i < 5; i++ {
		go philos[i].eat(3)
	}
	wg.Wait();
}

