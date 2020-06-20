/*coursera Golang Chapter 1 Final Activity: read.go
Write a program which reads information from a file and represents it in a slice of structs.
Assume that there is a text file which contains a series of names.
Each line of the text file has a first name and a last name, in that order,
separated by a single space on the line.

Your program will define a name struct which has two fields,
fname for the first name, and lname for the last name.
Each field will be a string of size 20 (characters). // a string in Runes can't really have a lenght limit so truncating after reading

Your program should prompt the user for the name of the text file.
Your program will successively read each line of the text file and
create a struct which contains the first and last names found in the file.
Each struct created will be added to a slice, and after all lines
have been read from the file, your program will have a slice
containing one struct for each line in the file.
After reading all lines from the file, your program should iterate through
your slice of structs and print the first and last names found in each struct.*/
// PWagner June20th2020
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func input() string {
	fmt.Printf("Enter a filename to read > ")
	myReader := bufio.NewReader(os.Stdin)    // returns a Reader-object
	myString, _ := myReader.ReadString('\n') // read until end-of-line, ignore errors
	myString = strings.TrimSpace(myString)   // get rid of nagging spaces.
	//	fmt.Println("\ndebug info :", myString)
	if myString == "quit" {
		fmt.Println("Exiting.") // user decides to stop - regular exit (0)
		os.Exit(0)
	} else if myString == "" { // typing enter without text takes default "names.txt"
		myString = "names.txt"
		fmt.Println("Using default filename 'names.txt' in this directory.")
	}
	return myString
}

func main() {
	type name struct { // a name is the aggregate of fname&lname (first and last - name)
		fname string
		lname string
	}

	var myNames []name // as requested all read data shall be stored in a slice of struct type "name"

	myFileName := input()

	myFile, err := os.Open(myFileName) // check for existence and size of test-file
	if err != nil {
		fmt.Println("Can't open file. Error = ", err)
		os.Exit(1)
	} else {
		defer myFile.Close() // close the file at the end
	}
	if fileinfo, _ := myFile.Stat(); fileinfo.Size() > 1000 { // avoid issues with huge files
		fmt.Println("I don't want to work with files larger than 1000 Bytes !")
		os.Exit(1)
	}

	myScanner := bufio.NewScanner(myFile)
	var myLine string    // is the raw data after file.Scan().Text()
	var myWords []string // temp variable for splitting each line into tokens (=names)
	var tmpName name     // a temp struct variable
	var i int            // just an index variable

	for myScanner.Scan() { // continue to read until file.Scan() is False -> EoF reached
		myLine = myScanner.Text()
		myWords = strings.Split(myLine, " ") // split a line into slice of words separated by a Space (should be 2)
		//		fmt.Println("Zeile:", myWords) // debug only
		if len(myWords) != 2 { // if not 2 words in this line skip it
			fmt.Println("Houston, we've had a problem : Line in file does not consist of FirstName+Space+LastName.")
			continue
		} else { // first token assigned to .fname second to .lname
			//	     	fmt.Println("Words[0]:", myWords[0], " Words[1]:", myWords[1]) // debug only
			tmpName.fname = fmt.Sprintf("%1.20s", myWords[0]) // limit names to a maximum of 20 runes (=chars but not necessarily equal bytes)
			tmpName.lname = fmt.Sprintf("%1.20s", myWords[1])
			myNames = append(myNames, tmpName) // myNames is a growing slice of struct(fname/lname string)
		}
	}
	fmt.Println("\nThe slice 'myNames' en bloc :", myNames, "\n") // print the entire slice

	for i, tmpName = range myNames { // print slice line by line
		fmt.Println("Iterating through slice : myNames[", i, "].fname =", tmpName.fname, "+ myNames[", i, "].lname =", tmpName.lname)
	}
}
