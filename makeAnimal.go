package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal struct {
	name       string
	food       string
	locomotion string
	noise      string
}

type animal interface {
	CreateAnimal()
	SeeAnimal()
}

// created animal must be of type cow, bird, or snake

func (a *Animal) Eat() string {
	return fmt.Sprintf("This animal eats %s", a.food)
}

func (a *Animal) Move() string {
	return fmt.Sprintf("This animal gets around via %s", a.locomotion)
}

func (a *Animal) Speak() string {
	return fmt.Sprintf(a.noise)
}

func main() {
	animals := []Animal{}

	var currentAnimal Animal

	for {
		fmt.Println("Enter a type of animal (cow, bird, or snake), and info about it (eat, move, or speak)")
		reader := bufio.NewReader(os.Stdin)
		str, _ := reader.ReadString('\n')

		animalQuery := strings.Fields(str)
		for _, a := range animals {
			if a.name == animalQuery[0] {
				currentAnimal = a

				if animalQuery[1] == "move" {
					fmt.Println(currentAnimal.Move())
				} else if animalQuery[1] == "eat" {
					fmt.Println(currentAnimal.Eat())
				} else {
					fmt.Println(currentAnimal.Speak())
				}
			}
		}

	}
}