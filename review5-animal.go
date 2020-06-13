package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type animalStruct struct {
	Animal     string
	Food       string
	Locomotion string
	Spoke      string
}

func (a animalStruct) Eat() {
	fmt.Println("Food :", a.Food)
}

func (a animalStruct) Move() {
	fmt.Println("Move :", a.Locomotion)
}

func (a animalStruct) Speak() {
	fmt.Println("Speak :", a.Spoke)
}

func initAnimal() map[string]animalStruct {
	mp := make(map[string]animalStruct)
	mp["cow"] = animalStruct{
		Animal:     "cow",
		Food:       "grass",
		Locomotion: "walk",
		Spoke:      "moo",
	}
	mp["bird"] = animalStruct{
		Animal:     "bird",
		Food:       "worms",
		Locomotion: "fly",
		Spoke:      "peep",
	}
	mp["snake"] = animalStruct{
		Animal:     "snake",
		Food:       "mice",
		Locomotion: "slither",
		Spoke:      "hsss",
	}
	return mp
}

func processInput(animal, command string, animalMap map[string]animalStruct) {
	getAnimal, found := animalMap[animal]
	if found == false {
		fmt.Println("animal request not found")
	}
	fmt.Println("Animal :", getAnimal.Animal)
	switch strings.ToLower(command) {
	case "eat":
		getAnimal.Eat()
		break
	case "move":
		getAnimal.Move()
		break
	case "speak":
		getAnimal.Speak()
		break
	default:
		fmt.Println("command not found")
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	animalMap := initAnimal()
	fmt.Println("Type the input, type x and enter for finish input")
	for {
		fmt.Print("> ")
		scanner.Scan()
		val := scanner.Text()
		if strings.ToLower(val) == "x" {
			fmt.Println("time to stop")
			break
		}
		words := strings.Fields(val)
		if len(words) == 2 {
			processInput(words[0], words[1], animalMap)
		} else {
			fmt.Println("request not valid!")
		}
	}
}
