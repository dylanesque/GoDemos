package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal struct {
	species    string
	food       string
	locomotion string
	noise      string
	name 	   string
}

type Cow struct {
	species "cow"
	food "grass"
	noise "moo"
	locomotion "walk"
}

type Bird struct {
	species "bird"
	food "worms"
	noise "peep"
	locomotion "fly"
}

type Snake struct {
	food "mice"
	locomotion "slither"
	noise "hiss"
	species "snake"
}

type animal interface {
	CreateAnimal()
	SeeAnimal()
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

animalData := []Animal{
	{species: "cow", food: "grass", locomotion: "walk", noise: "moo"},
	{species: "bird", food: "worms", locomotion: "fly", noise: "peep"},
	{species: "snake", food: "mice", locomotion: "slither", noise: "hsss"},
	}

	animals := make([]Animal{})

func main() {


	var currentAnimal Animal
	var animalList []Animal

	for {
		fmt.Println(`Enter 'newanimal <animal-type> <animal-name>' 
		to create a new animal, or 'query <animal-name> <animal-fact>' to get information about an animal`)
		reader := bufio.NewReader(os.Stdin)
		str, _ := reader.ReadString('\n')

		if strings.Contains(str, "newanimal") {
			animalCreate := strings.Fields(str)
			var newAnimal animal
			for _, a := range animalData {
				if a.type == animalCreate[1] {
					newAnimal = a 
					newAnimal.name = animalCreate[2]
				}
			}
			animalList = append(animalList, newAnimal)
		}

		if strings.Contains(str, "query") {
			// fetch animal facts
			animalQuery := strings.Fields(str)
		}
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