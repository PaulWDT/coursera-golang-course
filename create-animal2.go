/* Write a program which allows the user to create a set of animals and to get information about those animals.
Each animal has a name and can be either a cow, bird, or snake.
With each command, the user can either create a new animal of one of the three types,
or the user can request information about an animal that he/she has already created.
Each animal has a unique name, defined by the user.
Note that the user can define animals of a chosen type, but the types of animals are restricted to either cow, bird, or snake.
The following table contains the three types of animals and their associated data.

Animal	Food eaten	Locomotion method	Spoken sound
cow		grass 		walk 				moo
bird	worms		fly					peep
snake	mice		slither				hsss

Your program should present the user with a prompt, “>”, to indicate that the user can type a request.
Your program should accept one command at a time from the user, print out a response, and print out a new prompt on a new line.
Your program should continue in this loop forever.
Every command from the user must be either a “newanimal” command or a “query” command.

Each “newanimal” command must be a single line containing three strings.
The first string is “newanimal”.
The second string is an arbitrary string which will be the name of the new animal.
The third string is the type of the new animal, either “cow”, “bird”, or “snake”.
Your program should process each newanimal command by creating the new animal and printing “Created it!” on the screen.

Each “query” command must be a single line containing 3 strings.
The first string is “query”.
The second string is the name of the animal.
The third string is the name of the information requested about the animal, either “eat”, “move”, or “speak”.
Your program should process each query command by printing out the requested data.

Define an interface type called Animal which describes the methods of an animal.
Specifically, the Animal interface should contain the methods Eat(), Move(), and Speak(), which take no arguments and return no values.
The Eat() method should print the animal’s food, the Move() method should print the animal’s locomotion,
and the Speak() method should print the animal’s spoken sound.

Define three types Cow, Bird, and Snake.
For each of these three types, define methods Eat(), Move(), and Speak() so that the types Cow, Bird, and Snake all satisfy the Animal interface.
When the user creates an animal, create an object of the appropriate type.
Your program should call the appropriate method when the user issues a query command. */

package main

import (
	"fmt"
	"os"
	"strings"
)

// Animal is an interface requiring the methods
type Animal interface {
	Eat()
	Move()
	Speak()
}

// Cow satisfies the interface Animal
type Cow struct{ species, food, locomotion, noise string }

// Bird satisfies the interface Animal
type Bird struct{ species, food, locomotion, noise string }

// Snake satisfies the interface Animal
type Snake struct{ species, food, locomotion, noise string }

// Eat is a method for Cow
func (a Cow) Eat() { fmt.Println("A Cow eats", a.food, ".") }

// Move is a method for Cow
func (a Cow) Move() { fmt.Println("A Cow will", a.locomotion, ".") }

// Speak is a method for Cow
func (a Cow) Speak() { fmt.Println("A Cow will talk with a", a.noise, ".") }

// Eat is a method for Bird
func (a Bird) Eat() { fmt.Println("A Bird eats", a.food, ".") }

// Move is a method for Bird
func (a Bird) Move() { fmt.Println("A Bird will", a.locomotion, ".") }

// Speak is a method for Bird
func (a Bird) Speak() { fmt.Println("A Bird will talk with a", a.noise, ".") }

// Eat is a method for Snake
func (a Snake) Eat() { fmt.Println("A Snake eats", a.food, ".") }

// Move is a method for Snake
func (a Snake) Move() { fmt.Println("A Snake will", a.locomotion, ".") }

// Speak is a method for Snake
func (a Snake) Speak() { fmt.Println("A Snake will talk with a", a.noise, ".") }

// getInput asks the user for some values
func getInput() (string, string, string) { // don't forget the brackets around multiple return values !
	var command, who, what string

	fmt.Print("Please give me a command (3 words) :\n > ")

	_, err := fmt.Scan(&command) // scan input and returns the first string found or error
	check(err)
	_, err = fmt.Scan(&who) // scan input and returns the next string found or error
	check(err)
	_, err = fmt.Scan(&what) // scan input and returns next string found or error
	check(err)

	// fmt.Println("Input values :", command, who, what) // debug only
	// fmt.Printf("Input value types : %T, %T\n", who, what) // debug only
	return command, who, what
}

// check exits with a message if an error occured
func check(err error) {
	if err != nil {
		panic("Input error. Exiting.") // exit on error
	}
}

// main - look in the print-statement below for what it does ;-)
func main() {
	// aMap is the "database" for all created Animals key is the nick-name and value is of struct-type Animal
	var aMap = make(map[string]Animal)

	// myBeast art the three different animals you can create with their properties
	var aCow = Cow{"cow", "grass", "walk", "moo"}
	var aBird = Bird{"bird", "worms", "fly", "peep"}
	var aSnake = Snake{"snake", "mice", "slither", "hsss"}

	fmt.Println("This program will let you create a map of animals by their nick-name.")
	fmt.Println("then it will allow you to query information about animals.")
	fmt.Println("It will ask for three words after the prompt '>'")
	fmt.Println("1st word : a command : 'newanimal'|'n' or 'query'|'q').") // the name of an animal (either 'cow', 'bird' or 'snake')")
	fmt.Println("2nd word : the nick-name of an animal.")
	fmt.Println("3rd word : for 'newanimal' you need to give the species of the animal to create.")
	fmt.Println("or       : for 'query' you should give the property you wish to know (either “eat”, “move”, or “speak”)")
	fmt.Println("example: > newanimal Milka cow   -> will create a purple bovine named 'Milka'.")
	fmt.Println("example: > query Milka speak   -> will print 'Milka', the cow says 'moo'.")
	fmt.Println("Use Ctrl+c to exit.\n")
	var create, query, isInMap bool
	var command, who, what string
	var beast Animal

	for { // loop forever until break or error.
		command, who, what = getInput()                                           // get a command, the animal and the parameter to apply on it from the user (3 strings).
		command, who, what = strings.ToLower(command), who, strings.ToLower(what) // beware : the name of the animal will be case sensitive!

		switch command { // first recognize what action to perform (either create new animal or query properties) ...
		case "newanimal", "n":
			create, query = true, false
		case "query", "q":
			create, query = false, true
		default:
			fmt.Println("Sorry, command not recognized ('newanimal'/'n' or 'query'/'q'). Exiting.")
			os.Exit(0) // exit on unknown command
		}

		if create {
			switch what { // recognize what kind of animal to create ...
			case "cow":
				aMap[who] = aCow
				fmt.Println("Created it!")
			case "bird":
				aMap[who] = aBird
				fmt.Println("Created it!")
			case "snake":
				aMap[who] = aSnake
				fmt.Println("Created it!")
			default:
				fmt.Println("Sorry, I can't create the beast " + what + ".") // be tolerant and continue on error or ...
				// os.Exit(0) // exit on unknown animal
			}
		}
		// uncommenting the following line will show the content of the map in each loop
		// fmt.Println("DEBUG-INFO: the Animal-Map: ", aMap) // debug only

		if query { // find the animal as the value in the map with the name (who) as the key.
			beast, isInMap = aMap[who] // isInMap will only be true if the beast is found in the map
			if isInMap == true {

				switch what { // and then call the appropriate method to answer the question.
				case "eat":
					beast.Eat()
				case "move":
					beast.Move()
				case "speak":
					beast.Speak()
				default:
					fmt.Println("Sorry, I don't know the property " + what + " of the beast " + who + ".")
					// os.Exit(0) // exit on unknown question
				}
			} else {
				fmt.Println("Sorry, the beast " + who + " has never been created!.") // who was not found in the "aMap"
				// os.Exit(0) // exit on unknown question}
			}
		}
	}
}
