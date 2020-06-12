// percipio/Coursera "golang-functions-methods"
// Module 2 Activity: calculate-linear-movement
// done Paul.Wagner@telekom.de 2020/06/11

package main

import (
	"fmt"
)

// check exits with a message if an error occured
func check(err error) {
	if err != nil {
		panic("Input error. Exiting.") // exit on error
	}

}

// getInput asks the user for some values
func getInput() (float64, float64, float64) { // don't forget the brackets around multiple return values !
	var position, speed, acceleration float64

	fmt.Println("Please enter an initial position [m], an initial speed [m/s] and a constant acceleration [m/s²] as float values.")

	_, err := fmt.Scan(&position) // scan input and returns the first number found or error
	check(err)

	_, err = fmt.Scan(&speed) // scan input and returns next number found or error
	check(err)

	_, err = fmt.Scan(&acceleration) // scan input and returns the next number found or error
	check(err)

	// fmt.Println("Input values :", position,  speed,  acceleration) // debug only
	// fmt.Printf("Input value types : %T, %T, %T \n", position, speed, acceleration) // debug only

	return position, speed, acceleration
}

// getTime asks the user for a time value
func getTime() float64 {
	var time float64

	fmt.Println("Please enter a duration [seconds] for the movement as a float value.")

	_, err := fmt.Scan(&time) // scan input and returns the number found or error
	check(err)

	// fmt.Println("Input value :", time, " seconds.") // debug only
	// fmt.Printf("Input value type : %T \n", time)  // debug only

	return time
}

// GenDisplaceFn generates a function that calculates linear movement as func(time)
func GenDisplaceFn(ival1, ival2, ival3 float64) func(float64) float64 {
	/*	var generatedFunction func(float64) float64  // verbose version of code
		generatedFunction = func(time float64) float64 {
			posAfterTime := ival1 + ival2*time + 0.5*ival3*time*time
			return posAfterTime
		}
		return generatedFunction */

	return func(time float64) float64 { // short code with same result // brevity is the soul of wit (W.S.)
		return ival1 + ival2*time + 0.5*ival3*time*time
	}
}

// main asks for user input, creates a function with these parameter
// asks for another user input and the evaluates and prints the result of the function.
func main() {
	s0, v0, a := getInput()
	posFunction := GenDisplaceFn(s0, v0, a)
	t := getTime()

	fmt.Println("The initial position", s0, "[m], speed", v0, "[m/s] and acceleration", a, "[m/s²]")
	fmt.Println("will yield after", t, "[seconds] of linear movement")
	fmt.Println("a position of", posFunction(t), "[meters].")
}
