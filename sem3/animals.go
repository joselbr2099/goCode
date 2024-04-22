package main

import "fmt"

type Animal struct {
	food       string
	locomotion string
	noise      string
}

var animals map[string]Animal = map[string]Animal{
	"cow": Animal{
		food:       "grass",
		locomotion: "walk",
		noise:      "moo",
	},
	"bird": Animal{
		food:       "worms",
		locomotion: "fly",
		noise:      "peep",
	},
	"snake": Animal{
		food:       "mice",
		locomotion: "slither",
		noise:      "hsss",
	},
}

func main() {
	for {
		var animalName, command string
		fmt.Print("> ")
		fmt.Scan(&animalName)
		fmt.Scan(&command)

		runCommand(animalName, command)
		fmt.Println()
	}
}

func runCommand(animalName, command string) {

	animal, ok := animals[animalName]
	if !ok {
		fmt.Println("Animal name was wrong!")
		return
	}

	switch command {
	case "eat":
		animal.Eat()
	case "move":
		animal.Move()
	case "speak":
		animal.Speak()
	default:
		fmt.Println("There is no command like", command)
	}
}

func (animal Animal) Eat() {
	fmt.Println(animal.food)
}

func (animal Animal) Move() {
	fmt.Println(animal.locomotion)
}
func (animal Animal) Speak() {
	fmt.Println(animal.noise)
}
