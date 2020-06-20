package main

import "fmt"

var myFloat float64

func main() {
	fmt.Print("Please enter a floating point value : ")

	fmt.Scan(&myFloat) // scan input and returns the first number found or error

	fmt.Println("\nThis is the truncated number :", int(myFloat))
}
