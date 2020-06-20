package main

import "fmt"

func main() {
	var number1 float32
	fmt.Println("Enter a float number")
	fmt.Scan(&number1)
	var truncate = int64(number1)
	fmt.Println(truncate)
}
