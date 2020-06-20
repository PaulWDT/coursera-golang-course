/*coursera Golang Chapter 1 Module 3 Activity: slice.go
Write a program which prompts the user to first enter a name, and then enter an address. 
Your program should create a map and add the name and address to the map using 
the keys “name” and “address”, respectively. Your program should use Marshal() to create a 
JSON object from the map, and then your program should print the JSON object. */
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func input() string {
	fmt.Printf("Enter a Name and an address or 'x' to quit > ")
	myReader := bufio.NewReader(os.Stdin)    // returns a Reader-object
	myString, _ := myReader.ReadString('\n') // read until end-of-line, ignore errors
	myString = strings.TrimSpace(myString)   // get rid of nagging spaces.
	fmt.Println("\ndebug info :", myString)
	if myString == "x" {
		fmt.Println("Exiting.") // user decides to stop - regular exit (0)
		os.Exit(0)
	}
	return myString
}

func main() {
	var people := make( map[string]string )
	myints := make([]int, 0, 3) // makes new slice of length ZERO but Capacity three ! Ambiguity in problem statment !!!

	fmt.Println("Enter numbers, I will add them to a sorted slice ...")
	for {
		temp = input()
		myints = append(myints, temp) // we would never overwrite the zeros if we really did "make(myints,3)" size understood as length and not capacity !
		sort.Ints(myints)

		fmt.Println("The slice is now :", myints)
	}
}
