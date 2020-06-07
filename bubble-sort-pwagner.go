// percipio/Coursera "golang-functions-methods"
// Module 1 Activity: Bubble Sort Program
// done Paul.Wagner@telekom.de 2020/06/07
// could be optimized to take advantage of the implicit call-by-reference for slices
package main

import (
	"fmt"
)

// getInput() Allows the user to enter a slice of ints of an arbitrary length.
// reads from stdin (therefore needs no argument) and returns a slice of integers (that can be empty!).
func getInput() []int {
	mySlice := make([]int, 0) // start with an empty slice
	fmt.Println("Please enter a quantity of integer numbers line by line. Typing 'quit' or anything other than numbers will finish the process.")

	var myNum int
	for { // infinite loop - until user decides to finish
		fmt.Print("Enter a number :")
		_, err := fmt.Scan(&myNum) // scan through the input and returns the first integer found or returns error
		if err != nil {
			fmt.Println("Input was not a number. Ending input process.") // consider error as the users decision to finish inputting numbers
			return mySlice                                               // exiting the for loop and returning the result
		}
		mySlice = append(mySlice, myNum)
	}
	// return mySlice // no need for "return" here - the for loop never ends here - !

}

// mySwap returns its two arguments just in reverse order.
func mySwap(a, b int) (int, int) {
	return b, a
}

// mySort is a simple Bubblesort, no optimizations.
// takes a slice of int as argument and returns a sorted slice (ascending order)
func mySort(inSlice []int) { //} []int {
	if len(inSlice) < 2 { // if the slice is empty or only 1 number return immediately
		return inSlice
	}
	for i := 1; i < len(inSlice); i++ { // loop n times : n is one less than the number of elements
		for j := 0; j < len(inSlice)-i; j++ { // iterate through the first (len-n) elements : the upper elements being already sorted
			if inSlice[j] > inSlice[j+1] { // if the element is bigger than its right neighbour then swap them
				inSlice[j], inSlice[j+1] = mySwap(inSlice[j], inSlice[j+1])
			}
		}
	}
	return inSlice // return a sorted slice
}

// main() gets a slice of ints from the user, sorts the slice and prints it.
func main() {
	inSlice := getInput()
	fmt.Println("This is the Slice you just gave me :", inSlice)

	fmt.Println("Trying to Bubblesort this Slice.")
	outSlice := mySort(inSlice)

	fmt.Println("Resulting in :", outSlice)
	fmt.Println("Finished. Exiting.")
}
