// percipio/Coursera "golang-functions-methods"
// Module 3 Activity: ask-the-animal.go
// done Paul.Wagner@telekom.de 2020/06/12

/* Write a program which allows the user to get information about a predefined set of animals.
Three animals are predefined, cow, bird, and snake. Each animal can eat, move, and speak.
The user can issue a request to find out one of three things about an animal:
1) the food that it eats, 2) its method of locomotion, and 3) the sound it makes when it speaks.
The following table contains the three animals and their associated data which should be hard-coded into your program.

Animal	Food eaten	Locomotion method	Spoken sound
cow		grass		walk				moo
bird	worms		fly					peep
snake	mice		slither				hsss
Your program should present the user with a prompt, “>”, to indicate that the user can type a request.
Your program accepts one request at a time from the user, prints out the answer to the request,
and prints out a new prompt. Your program should continue in this loop forever.
Every request from the user must be a single line containing 2 strings.
The first string is the name of an animal, either “cow”, “bird”, or “snake”.
The second string is the name of the information requested about the animal, either “eat”, “move”, or “speak”.
Your program should process each request by printing out the requested data.

You will need a data structure to hold the information about each animal.
Make a type called Animal which is a struct containing three fields:
food, locomotion, and noise, all of which are strings. Make three methods called Eat(), Move(), and Speak().
The receiver type of all of your methods should be your Animal type.
The Eat() method should print the animal’s food, the Move() method should print the animal’s locomotion,
and the Speak() method should print the animal’s spoken sound.
Your program should call the appropriate method when the user makes a request. */
package main

import (
	"fmt"
	"os"
)

// Animal is the base type for this exercise
type Animal struct {
	food, locomotion, noise string
}

// Eat is a method for Animal
func (a Animal) Eat() { fmt.Println("This animal eats", a.food, ".") }

// Move is a method for Animal
func (a Animal) Move() { fmt.Println("This animal will", a.locomotion, ".") }

// Speak is a method for Animal
func (a Animal) Speak() { fmt.Println("This aimal will talk with a", a.noise, ".") }

// getInput asks the user for some values
func getInput() (string, string) { // don't forget the brackets around multiple return values !
	var who, what string

	fmt.Print("Please give me the name of the animal and what you wish to know about it :\n > ")

	_, err := fmt.Scan(&who) // scan input and returns the first string found or error
	check(err)

	_, err = fmt.Scan(&what) // scan input and returns next string found or error
	check(err)

	//	fmt.Println("Input values :", who, " - ", what)       // debug only
	//	fmt.Printf("Input value types : %T, %T\n", who, what) // debug only
	return who, what
}

// check exits with a message if an error occured
func check(err error) {
	if err != nil {
		panic("Input error. Exiting.") // exit on error
	}
}

// main - look in the print-statement below for what it does ;-)
func main() {
	var myCow = Animal{"grass", "walk", "moo"}
	var myBird = Animal{"worms", "fly", "peep"}
	var mySnake = Animal{"mice", "slither", "hsss"}

	fmt.Println("This program will give you information about animals.")
	fmt.Println("It will ask for two words after the prompt '>'")
	fmt.Println("1st word : the name of an animal (either 'cow', 'bird' or 'snake')")
	fmt.Println("2nd word : the property you wish to know (either “eat”, “move”, or “speak”) \n")

	var who, what string
	var thisAnimal Animal

	for { // loop forever until break or error.
		who, what = getInput() // get an animal and the question for it from the user

		switch who { // first recognize what kind of animal we have ...
		case "cow":
			thisAnimal = myCow
		case "bird":
			thisAnimal = myBird
		case "snake":
			thisAnimal = mySnake
		default:
			fmt.Println("Sorry, I don't know the beast " + who + ". Exiting.")
			os.Exit(0) // exit on unknown animal

		}

		switch what { // and then call the appropriate method to answer the question.
		case "eat":
			thisAnimal.Eat()
		case "move":
			thisAnimal.Move()
		case "speak":
			thisAnimal.Speak()
		default:
			fmt.Println("Sorry, I don't know the property " + what + " of the beast " + who + ". Exiting.")
			os.Exit(0) // exit on unknown question
		}

	}

}
