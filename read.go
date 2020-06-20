/*coursera Golang Chapter 1 Final Activity: read.go
Write a program which reads information from a file and represents it in a slice of structs.
Assume that there is a text file which contains a series of names.
Each line of the text file has a first name and a last name, in that order,
separated by a single space on the line.

Your program will define a name struct which has two fields,
fname for the first name, and lname for the last name.
Each field will be a string of size 20 (characters).

Your program should prompt the user for the name of the text file.
Your program will successively read each line of the text file and
create a struct which contains the first and last names found in the file.
Each struct created will be added to a slice, and after all lines
have been read from the file, your program will have a slice
containing one struct for each line in the file.
After reading all lines from the file, your program should iterate through
your slice of structs and print the first and last names found in each struct.

Submit your source code for the program, “read.go”.

*/
// PWagner June20th2020
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	type name struct {
		fname string
		lname string
	}

	myFile, err := os.Open("names.txt") // check for existence and size of test-file
	if err != nil {
		fmt.Println("Can't open file. Error = ", err)
		os.Exit(1)
	} else {
		defer myFile.Close()
	}
	if fileinfo, _ := myFile.Stat(); fileinfo.Size() > 1000 { // avoid issues with huge files
		fmt.Println("I don't want to work with files larger than 1000 Bytes !")
		os.Exit(1)
	}

	myScanner := bufio.NewScanner(myFile)
	//	var myTxt []string

	for myScanner.Scan() {

		fmt.Println("Zeile:", myScanner.Text())

	}
}
