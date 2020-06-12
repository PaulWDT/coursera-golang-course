package main

import (
	"fmt"
	"strconv"
)

func main() {

	var first string
	var i int
	numbers := make([]int, 0)

	fmt.Println("Enter X to stop")

	for true {
		fmt.Print("Enter number: ")
		_, err := fmt.Scan(&first)
		i, err = strconv.Atoi(first)

		if err == nil {
			numbers = append(numbers, i) //append user input
			//sort.Ints(numbers)
			//fmt.Println("Sorted Array: ", numbers)
		} else {
			if first == "X" || first == "x" {
				break
			}
			fmt.Println("Got: " + first)
			fmt.Println("Enter valid value")
			continue
		}
	}

	fmt.Println("Entered array: ", numbers)
	bubblesort(numbers)
}

func bubblesort(items []int) {
	var (
		n      = len(items)
		sorted = false
	)
	for !sorted {
		swapped := false
		for i := 0; i < n-1; i++ {
			if items[i] > items[i+1] {

				items = swap(items, i)
				swapped = true
			}
		}
		if !swapped {
			sorted = true
		}
		n = n - 1
	}

	fmt.Println("Sorted array: ", items)
}

func swap(items []int, i int) []int {

	items[i+1], items[i] = items[i], items[i+1]
	return items
}
