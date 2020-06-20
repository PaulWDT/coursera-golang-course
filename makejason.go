/*coursera Golang Chapter 1 Module 4 Activity: makejson.go
Write a program which prompts the user to first enter a name, and then enter an address.
Your program should create a map and add the name and address to the map using
the keys “name” and “address”, respectively. Your program should use Marshal() to create a
JSON object from the map, and then your program should print the JSON object. */
// PWagner June20th2020
package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func input() string {
	myReader := bufio.NewReader(os.Stdin)    // returns a Reader-object
	myString, _ := myReader.ReadString('\n') // read until end-of-line, ignore errors.
	myString = strings.TrimSpace(myString)   // get rid of nagging spaces before and after.
	return myString
}

func main() {
	//	var people map[string]string
	people := make(map[string]string)

	fmt.Print("\nEnter a Name (enter) > ")
	myName := input()
	fmt.Print("\nEnter an Address (enter) > ")
	myAddress := input()

	people["name"] = myName
	people["address"] = myAddress

	//	var myJsonOut []byte
	//	var err error
	myJsonOut, err := json.Marshal(people) // umplicit declaration with ":=" is much shorter ;-)
	if err != nil {
		fmt.Println("JSON Marshalling of ", people, "failed with :", err)
	}

	fmt.Println("Map 'people' was marshalled into : ", myJsonOut)
	fmt.Println("Same as a string : ", string(myJsonOut))
}
