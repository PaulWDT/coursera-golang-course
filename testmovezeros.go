package main

import "fmt"

func main() {
	res := MoveZeros([]int{2, 0, 5, 7, 0, 3, 11, 13})
	fmt.Println(res)
}

func MoveZeros(arr []int) []int {
	// TODO: Program me
	al := len(arr)
	res := make([]int, al)
	ap := 0
	bp := al - 1

	for i := 0; i < al; i++ {
		if arr[i] == 0 {
			res[bp] = 0
			bp -= 1
		} else {
			res[ap] = arr[i]
			ap += 1
		}
	}
	return res
}
