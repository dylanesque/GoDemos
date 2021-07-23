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

/*

Examine the below code

func createNewAnimal(animalType string) Animal {
	switch animalType {
	case "cow":
		return Cow{"grass", "walk", "moo"}
	case "bird":
		return Bird{"worms", "fly", "peep"}
	case "snake":
		return Snake{"mice", "slither", "hsss"}
	default:
		return nil
	}
}

func main() {

		command := tokens[0]
		animalName := tokens[1]
		if command == "newanimal" {
			instance := createNewAnimal(tokens[2])
			if instance == nil {
				fmt.Println("invalid type")
				continue
			}

			animals[animalName] = instance
			fmt.Println("Created it!")
			continue
		}

		if command == "query" {
			animal, ok := animals[animalName]
			if ok {
				switch tokens[2] {
				case "eat":
					animal.Eat()
				case "move":
					animal.Move()
				case "speak":
					animal.Speak()
				default:
					fmt.Println("unknown query")
				}
			}
		}
	}

}

*/
