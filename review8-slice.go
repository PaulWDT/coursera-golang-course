package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	slice := make([]int, 3)
	ban := true
	var userInput string
	for ban {
		fmt.Println("Enter a number or x to exit")
		fmt.Scan(&userInput)
		if strings.ToLower(userInput) != "x" {
			number, err := strconv.Atoi(userInput)
			if err == nil {
				slice = append(slice, number)
				sort.Ints(slice)
				fmt.Println(slice)
			} else {
				fmt.Println("Enter a valid number")
			}

		} else {
			ban = false
		}
	}
	fmt.Println(slice)
}
