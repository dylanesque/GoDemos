/*

Write a program which allows the user to get information about a
 predefined set of animals. Three animals are predefined, cow,
 bird, and snake. Each animal can eat, move, and speak.
 The user can issue a request to find out one of three things
 about an animal: 1) the food that it eats, 2) its method of
 locomotion, and 3) the sound it makes when it speaks. T
 he following table contains the three animals and their associated data which should be hard-coded into your program.

Animal	Food eaten	Locomotion method	Spoken sound
cow	grass	walk	moo
bird	worms	fly	peep
snake	mice	slither	hsss

Your program should present the user with a prompt, “>”,
to indicate that the user can type a request. Your program
accepts one request at a time from the user, prints out the
answer to the request, and prints out a new prompt. Your
program should continue in this loop forever. Every request
from the user must be a single line containing 2 strings.
The first string is the name of an animal, either “cow”,
“bird”, or “snake”. The second string is the name of the
information requested about the animal, either “eat”, “move”,
or “speak”. Your program should process each request by
printing out the requested data.

*/

package main

import (
	"fmt"
	"strings"
)

type Animal struct {
	name       string
	food       string
	locomotion string
	noise      string
}

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
	animals := []Animal{
		{name: "cow", food: "grass", locomotion: "walk", noise: "moo"},
		{name: "bird", food: "worms", locomotion: "fly", noise: "peep"},
		{name: "snake", food: "mice", locomotion: "slither", noise: "hsss"},
	}

	var s string
	var currentAnimal Animal

	for {
		fmt.Println("Enter a type of animal (cow, bird, or snake), and info about it (eat, move, or speak)")
		_, err := fmt.Scanf("%s", &s)

		if err != nil {
			fmt.Println(err)
		}

		animalQuery := strings.Fields(s)
		for _, a := range animals {
			if a.name == animalQuery[0] {
				currentAnimal = a
				for _, word := range animalQuery {
					fmt.Println(word)
				}
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
