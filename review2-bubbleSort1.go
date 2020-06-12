package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func swap(arr []int, indx int) {
	tmp := arr[indx]
	arr[indx] = arr[indx+1]
	arr[indx+1] = tmp
}
func BubbleSort(unsorted []int) {
	for i, _ := range unsorted {
		for j := 0; j < len(unsorted)-i-1; j++ {
			if unsorted[j] > unsorted[j+1] {
				swap(unsorted, j)
			}
		}
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var input string
	intArr := make([]int, 0, 10)
	fmt.Printf("Enter numbers space-seperated: ")
	if scanner.Scan() {
		input = scanner.Text()
	}
	numbers := strings.Split(input, " ")
	for _, num := range numbers {
		numInt, _ := strconv.Atoi(num)
		intArr = append(intArr, numInt)
	}
	fmt.Printf("Unsorted list: %v \n", intArr)
	BubbleSort(intArr)
	fmt.Printf("Sorted list: %v \n", intArr)
}
